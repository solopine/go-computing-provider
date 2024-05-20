package wallet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/wallet/contract/collateral"
	"github.com/swanchain/go-computing-provider/wallet/contract/swan_token"
	"github.com/swanchain/go-computing-provider/wallet/tablewriter"
	"golang.org/x/xerrors"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	WalletRepo  = "keystore"
	KNamePrefix = "wallet-"
)

var (
	ErrKeyInfoNotFound = fmt.Errorf("key info not found")
	ErrKeyExists       = fmt.Errorf("key already exists")
)

var reAddress = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func SetupWallet(dir string) (*LocalWallet, error) {
	cpPath, exit := os.LookupEnv("CP_PATH")
	if !exit {
		return nil, fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
	}
	if err := conf.InitConfig(cpPath, true); err != nil {
		return nil, fmt.Errorf("load config file failed, error: %+v", err)
	}

	kstore, err := OpenOrInitKeystore(filepath.Join(cpPath, dir))
	if err != nil {
		return nil, err
	}
	return NewWallet(kstore), nil
}

type LocalWallet struct {
	keys     map[string]*KeyInfo
	keystore KeyStore

	lk sync.Mutex
}

func NewWallet(keystore KeyStore) *LocalWallet {
	w := &LocalWallet{
		keys:     make(map[string]*KeyInfo),
		keystore: keystore,
	}
	return w
}

func (w *LocalWallet) WalletSign(ctx context.Context, addr string, msg []byte) (string, error) {
	defer w.keystore.Close()
	ki, err := w.FindKey(addr)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("signing using private key '%s': %w", addr, ErrKeyInfoNotFound)
	}
	signByte, err := Sign(ki.PrivateKey, msg)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signByte), nil
}

func (w *LocalWallet) WalletVerify(ctx context.Context, addr string, sigByte []byte, data string) (bool, error) {
	hash := crypto.Keccak256Hash([]byte(data))
	return Verify(addr, sigByte, hash.Bytes())
}

func (w *LocalWallet) FindKey(addr string) (*KeyInfo, error) {
	defer w.keystore.Close()
	w.lk.Lock()
	defer w.lk.Unlock()

	k, ok := w.keys[addr]
	if ok {
		return k, nil
	}
	if w.keystore == nil {
		log.Warn("FindKey didn't find the key in in-memory wallet")
		return nil, nil
	}

	ki, err := w.tryFind(addr)
	if err != nil {
		if xerrors.Is(err, ErrKeyInfoNotFound) {
			return nil, nil
		}
		return nil, xerrors.Errorf("getting from keystore: %w", err)
	}

	w.keys[addr] = &ki
	return &ki, nil
}

func (w *LocalWallet) tryFind(key string) (KeyInfo, error) {
	ki, err := w.keystore.Get(KNamePrefix + key)
	if err == nil {
		return ki, err
	}

	if !xerrors.Is(err, ErrKeyInfoNotFound) {
		return KeyInfo{}, err
	}

	return ki, nil
}

func (w *LocalWallet) WalletExport(ctx context.Context, addr string) (*KeyInfo, error) {
	defer w.keystore.Close()
	k, err := w.FindKey(addr)
	if err != nil {
		return nil, xerrors.Errorf("failed to find key to export: %w", err)
	}
	if k == nil {
		return nil, xerrors.Errorf("private key not found for %s", addr)
	}

	return k, nil
}

func (w *LocalWallet) WalletImport(ctx context.Context, ki *KeyInfo) (string, error) {
	defer w.keystore.Close()
	if ki == nil || len(strings.TrimSpace(ki.PrivateKey)) == 0 {
		return "", fmt.Errorf("not found private key")
	}

	_, publicKeyECDSA, err := ToPublic(ki.PrivateKey)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	existAddress, err := w.tryFind(address)
	if err == nil && existAddress.PrivateKey != "" {
		return "", xerrors.Errorf("This wallet address already exists")
	}

	if err := w.keystore.Put(KNamePrefix+address, *ki); err != nil {
		return "", xerrors.Errorf("saving to keystore: %w", err)
	}
	return "", nil
}

func (w *LocalWallet) WalletList(ctx context.Context, chainName string, contractFlag bool) error {
	defer w.keystore.Close()
	addressList, err := w.addressList(ctx)
	if err != nil {
		return err
	}

	addressKey := "Address"
	balanceKey := "Balance"
	nonceKey := "Nonce"
	errorKey := "Error"

	chainRpc, err := conf.GetRpcByName(chainName)
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return err
	}
	defer client.Close()

	var wallets []map[string]interface{}
	for _, addr := range addressList {
		var balance string
		if contractFlag {
			tokenStub, err := swan_token.NewTokenStub(client, swan_token.WithPublicKey(addr))
			if err == nil {
				balance, err = tokenStub.BalanceOf()
			}
		} else {
			balance, err = Balance(ctx, client, addr)
		}

		var errmsg string
		if err != nil {
			errmsg = err.Error()
		}

		nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
		if err != nil {
			errmsg = err.Error()
		}

		wallet := map[string]interface{}{
			addressKey: addr,
			balanceKey: balance,
			errorKey:   errmsg,
			nonceKey:   nonce,
		}
		wallets = append(wallets, wallet)
	}

	tw := tablewriter.New(
		tablewriter.Col(addressKey),
		tablewriter.Col(balanceKey),
		tablewriter.Col(nonceKey),
		tablewriter.NewLineCol(errorKey))

	for _, wallet := range wallets {
		tw.Write(wallet)
	}
	return tw.Flush(os.Stdout)
}

func (w *LocalWallet) WalletNew(ctx context.Context) (string, error) {
	defer w.keystore.Close()

	w.lk.Lock()
	defer w.lk.Unlock()

	privateK, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateK)
	privateKey := hexutil.Encode(privateKeyBytes)[2:]

	_, publicKeyECDSA, err := ToPublic(privateKey)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	keyInfo := KeyInfo{PrivateKey: privateKey}
	if err := w.keystore.Put(KNamePrefix+address, keyInfo); err != nil {
		return "", xerrors.Errorf("saving to keystore: %w", err)
	}
	w.keys[address] = &keyInfo

	return address, nil
}

func (w *LocalWallet) walletDelete(ctx context.Context, addr string) error {
	w.lk.Lock()
	defer w.lk.Unlock()

	if err := w.keystore.Delete(KNamePrefix + addr); err != nil {
		return xerrors.Errorf("failed to delete key %s: %w", addr, err)
	}

	delete(w.keys, addr)

	return nil
}

func (w *LocalWallet) WalletDelete(ctx context.Context, addr string) error {
	defer w.keystore.Close()
	if err := w.walletDelete(ctx, addr); err != nil {
		return xerrors.Errorf("wallet delete: %w", err)
	}
	fmt.Printf("%s has been deleted from the local success \n", addr)
	return nil
}

func (w *LocalWallet) WalletSend(ctx context.Context, chainName string, from, to string, amount string) (string, error) {
	defer w.keystore.Close()
	chainUrl, err := conf.GetRpcByName(chainName)
	if err != nil {
		return "", err
	}
	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private %w,", from, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	sendAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	txHash, err := sendTransaction(client, ki.PrivateKey, to, sendAmount)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func (w *LocalWallet) WalletCollateral(ctx context.Context, chainName string, from string, amount string, collateralType string) (string, error) {
	defer w.keystore.Close()
	sendAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByName(chainName)
	if err != nil {
		return "", err
	}
	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", from, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if collateralType == "fcp" {
		tokenStub, err := swan_token.NewTokenStub(client, swan_token.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}

		swanTokenTxHash, err := tokenStub.Approve(sendAmount)
		if err != nil {
			return "", err
		}

		timeout := time.After(3 * time.Minute)
		ticker := time.Tick(3 * time.Second)
		for {
			select {
			case <-timeout:
				return "", fmt.Errorf("timeout waiting for transaction confirmation, tx: %s", swanTokenTxHash)
			case <-ticker:
				receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(swanTokenTxHash))
				if err != nil {
					if errors.Is(err, ethereum.NotFound) {
						continue
					}
					return "", fmt.Errorf("mintor swan token Approve tx, error: %+v", err)
				}

				if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
					collateralStub, err := collateral.NewCollateralStub(client, collateral.WithPrivateKey(ki.PrivateKey))
					if err != nil {
						return "", err
					}
					collateralTxHash, err := collateralStub.Deposit(sendAmount)
					if err != nil {
						return "", err
					}
					return collateralTxHash, nil
				} else if receipt != nil && receipt.Status == 0 {
					return "", fmt.Errorf("swan token approve transaction execution failed, tx: %s", swanTokenTxHash)
				}
			}
		}
	} else {
		zkCollateral, err := account.NewCollateralStub(client, account.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}
		collateralTxHash, err := zkCollateral.Deposit(sendAmount)
		if err != nil {
			return "", err
		}
		return collateralTxHash, nil
	}
}

func (w *LocalWallet) CollateralInfo(ctx context.Context, chainName string, collateralType string) error {
	defer w.keystore.Close()
	addrs, err := w.addressList(ctx)
	if err != nil {
		return err
	}

	addressKey := "Address"
	balanceKey := "Balance"
	collateralKey := "Collateral"
	frozenKey := "Escrow"
	errorKey := "Error"

	chainRpc, err := conf.GetRpcByName(chainName)
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return err
	}
	defer client.Close()

	var wallets []map[string]interface{}
	for _, addr := range addrs {
		var balance, collateralBalance, frozenCollateral string
		balance, err = Balance(ctx, client, addr)

		if collateralType == "fcp" {
			collateralStub, err := collateral.NewCollateralStub(client, collateral.WithPublicKey(addr))
			if err == nil {
				collateralBalance, err = collateralStub.Balances()
			}
			frozenCollateral, err = getFrozenCollateral(addr)
		}
		//} else {
		//	zkCollateral, err := account.NewCollateralStub(client)
		//	if err == nil {
		//		cpInfo, err := zkCollateral.CpInfo(addr)
		//		if err == nil {
		//			collateralBalance = cpInfo.CollateralBalance
		//			frozenCollateral = cpInfo.FrozenBalance
		//		}
		//	}
		//}

		var errmsg string
		if err != nil {
			errmsg = err.Error()
		}

		wallet := map[string]interface{}{
			addressKey:    addr,
			balanceKey:    balance,
			collateralKey: collateralBalance,
			frozenKey:     frozenCollateral,
			errorKey:      errmsg,
		}
		wallets = append(wallets, wallet)
	}

	tw := tablewriter.New(
		tablewriter.Col(addressKey),
		tablewriter.Col(balanceKey),
		tablewriter.Col(collateralKey),
		tablewriter.Col(frozenKey),
		tablewriter.NewLineCol(errorKey))

	for _, wallet := range wallets {
		tw.Write(wallet)
	}
	return tw.Flush(os.Stdout)
}

func (w *LocalWallet) CollateralWithdraw(ctx context.Context, chainName string, to string, amount string, collateralType string) (string, error) {
	defer w.keystore.Close()
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByName(chainName)
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(to)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", to, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	if collateralType == "fcp" {
		collateralStub, err := collateral.NewCollateralStub(client, collateral.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}
		return collateralStub.Withdraw(withDrawAmount)
	} else {
		zkCollateral, err := account.NewCollateralStub(client, account.WithPrivateKey(ki.PrivateKey))
		if err != nil {
			return "", err
		}
		return zkCollateral.Withdraw(withDrawAmount)
	}
}

func (w *LocalWallet) CollateralSend(ctx context.Context, chainName, from, to string, amount string) (string, error) {
	defer w.keystore.Close()
	withDrawAmount, err := convertToWei(amount)
	if err != nil {
		return "", err
	}

	chainUrl, err := conf.GetRpcByName(chainName)
	if err != nil {
		return "", err
	}

	ki, err := w.FindKey(from)
	if err != nil {
		return "", err
	}
	if ki == nil {
		return "", xerrors.Errorf("the address: %s, private key %w,", to, ErrKeyInfoNotFound)
	}

	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		return "", err
	}
	defer client.Close()

	collateralStub, err := swan_token.NewTokenStub(client, swan_token.WithPrivateKey(ki.PrivateKey))
	if err != nil {
		return "", err
	}
	withdrawHash, err := collateralStub.Transfer(to, withDrawAmount)
	if err != nil {
		return "", err
	}

	return withdrawHash, nil
}

func (w *LocalWallet) addressList(ctx context.Context) ([]string, error) {
	defer w.keystore.Close()
	all, err := w.keystore.List()
	if err != nil {
		return nil, xerrors.Errorf("listing keystore: %w", err)
	}

	addressList := make([]string, 0, len(all))
	for _, a := range all {
		if strings.HasPrefix(a, KNamePrefix) {
			addr := strings.TrimPrefix(a, KNamePrefix)
			addressList = append(addressList, addr)
		}
	}
	return addressList, nil
}

func convertToWei(ethValue string) (*big.Int, error) {
	ethFloat, ok := new(big.Float).SetString(ethValue)
	if !ok {
		return nil, fmt.Errorf("conversion to float failed")
	}
	weiConversion := new(big.Float).SetFloat64(1e18)
	weiFloat := new(big.Float).Mul(ethFloat, weiConversion)
	weiInt, acc := new(big.Int).SetString(weiFloat.Text('f', 0), 10)
	if !acc {
		return nil, fmt.Errorf("conversion to Wei failed")
	}
	return weiInt, nil
}

func getFrozenCollateral(walletAddress string) (string, error) {
	url := fmt.Sprintf("%s/check_holding_collateral/%s", conf.GetConfig().HUB.ServerUrl, walletAddress)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var frozenResp struct {
		Data struct {
			FrozenCollateral big.Int `json:"Frozen_Collateral"`
		} `json:"data"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}

	if err = json.Unmarshal(body, &frozenResp); err != nil {
	}
	fbalance := new(big.Float)
	fbalance.SetString(frozenResp.Data.FrozenCollateral.String())
	etherQuotient := new(big.Float).Quo(fbalance, new(big.Float).SetInt(big.NewInt(1e18)))
	ethValue := etherQuotient.Text('f', 5)
	return ethValue, nil
}
