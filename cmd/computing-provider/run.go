package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"github.com/olekukonko/tablewriter"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/initializer"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/swanchain/go-computing-provider/wallet"
	"github.com/swanchain/go-computing-provider/wallet/contract/collateral"
	"github.com/urfave/cli/v2"
	"io"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Start a cp process",
	Action: func(cctx *cli.Context) error {
		logs.GetLogger().Info("Start in computing provider mode.")

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		initializer.ProjectInit(cpRepoPath)

		r := gin.Default()
		r.Use(cors.Middleware(cors.Config{
			Origins:         "*",
			Methods:         "GET, PUT, POST, DELETE",
			RequestHeaders:  "Origin, Authorization, Content-Type",
			ExposedHeaders:  "",
			MaxAge:          50 * time.Second,
			ValidateHeaders: false,
		}))
		pprof.Register(r)

		v1 := r.Group("/api/v1")
		cpManager(v1.Group("/computing"))

		shutdownChan := make(chan struct{})
		httpStopper, err := util.ServeHttp(r, "cp-api", ":"+strconv.Itoa(conf.GetConfig().API.Port), true)
		if err != nil {
			logs.GetLogger().Fatal("failed to start cp-api endpoint: %s", err)
		}

		finishCh := util.MonitorShutdown(shutdownChan,
			util.ShutdownHandler{Component: "cp-api", StopFunc: httpStopper},
		)
		<-finishCh

		return nil
	},
}

func cpManager(router *gin.RouterGroup) {
	router.GET("/host/info", computing.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", computing.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", computing.RedeployJob)
	router.DELETE("/lagrange/jobs", computing.CancelJob)
	router.GET("/lagrange/cp", computing.StatisticalSources)
	router.POST("/lagrange/jobs/renew", computing.ReNewJob)
	router.GET("/lagrange/spaces/log", computing.GetSpaceLog)
	router.POST("/lagrange/cp/proof", computing.DoProof)

	router.GET("/cp", computing.StatisticalSources)
	router.GET("/cp/info", computing.GetCpInfo)
	router.POST("/cp/ubi", computing.DoUbiTaskForK8s)
	router.POST("/cp/receive/ubi", computing.ReceiveUbiProofForK8s)

}

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print computing-provider info",
	Action: func(cctx *cli.Context) error {
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		localNodeId := computing.GetNodeId(cpRepoPath)

		k8sService := computing.NewK8sService()
		var count int
		if k8sService.Version == "" {
			count = 0
		} else {
			count, _ = k8sService.GetDeploymentActiveCount()
		}

		chainRpc, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		var balance, collateralBalance, ownerBalance string
		var contractAddress, ownerAddress, beneficiaryAddress, ubiFlag, chainNodeId string

		cpStub, err := account.NewAccountStub(client)
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount failed, error: %v", err)
			}
			if cpAccount.UbiFlag == 1 {
				ubiFlag = "Accept"
			} else {
				ubiFlag = "Reject"
			}
			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			beneficiaryAddress = cpAccount.Beneficiary.BeneficiaryAddress
			chainNodeId = cpAccount.NodeId
		}

		balance, err = wallet.Balance(context.TODO(), client, conf.GetConfig().HUB.WalletAddress)
		collateralStub, err := collateral.NewCollateralStub(client, collateral.WithPublicKey(conf.GetConfig().HUB.WalletAddress))
		if err == nil {
			collateralBalance, err = collateralStub.Balances()
		}

		if ownerAddress != "" {
			ownerBalance, err = wallet.Balance(context.TODO(), client, ownerAddress)
		}

		var domain = conf.GetConfig().API.Domain
		if strings.HasPrefix(domain, ".") {
			domain = domain[1:]
		}
		var taskData [][]string

		taskData = append(taskData, []string{"Multi-Address:", conf.GetConfig().API.MultiAddress})
		taskData = append(taskData, []string{"Node ID:", localNodeId})
		taskData = append(taskData, []string{"ECP:"})
		taskData = append(taskData, []string{"   Contract Address:", contractAddress})
		taskData = append(taskData, []string{"   UBI FLAG:", ubiFlag})
		taskData = append(taskData, []string{"   Owner:", ownerAddress})
		taskData = append(taskData, []string{"   Beneficiary Address:", beneficiaryAddress})
		taskData = append(taskData, []string{"   Available(SWAN-ETH):", ownerBalance})
		taskData = append(taskData, []string{"   Collateral(SWAN-ETH):", "0"})
		taskData = append(taskData, []string{"FCP:"})
		taskData = append(taskData, []string{"   Wallet:", conf.GetConfig().HUB.WalletAddress})
		taskData = append(taskData, []string{"   Domain:", domain})
		taskData = append(taskData, []string{"   Running deployments:", strconv.Itoa(count)})
		taskData = append(taskData, []string{"   Available(SWAN-ETH):", balance})
		taskData = append(taskData, []string{"   Collateral(SWAN-ETH):", collateralBalance})

		var rowColor []tablewriter.Colors
		if ubiFlag == "Accept" {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
		} else {
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgRedColor}}
		}

		var rowColorList []RowColor
		rowColorList = append(rowColorList, RowColor{
			row:    4,
			column: []int{1},
			color:  rowColor,
		})

		header := []string{"Name:", conf.GetConfig().API.NodeName}
		NewVisualTable(header, taskData, rowColorList).Generate(false)
		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
		return nil
	},
}

var stateCmd = &cli.Command{
	Name:  "state",
	Usage: "Print computing-provider info on the chain",
	Subcommands: []*cli.Command{
		stateInfoCmd,
	},
}

var stateInfoCmd = &cli.Command{
	Name:      "cp-info",
	Usage:     "Print computing-provider chain info",
	ArgsUsage: "[cp_account_contract_address]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "chain",
			Usage: "Specify which rpc connection chain to use",
			Value: conf.DefaultRpc,
		},
	},
	Action: func(cctx *cli.Context) error {
		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		chainRpc, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			return err
		}
		client, err := ethclient.Dial(chainRpc)
		if err != nil {
			return err
		}
		defer client.Close()

		var ownerBalance string
		var contractAddress, ownerAddress, beneficiaryAddress, ubiFlag, chainNodeId, chainMultiAddress string

		cpStub, err := account.NewAccountStub(client, account.WithContractAddress(cctx.Args().Get(0)))
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount failed, error: %v", err)
			}
			if cpAccount.UbiFlag == 1 {
				ubiFlag = "Accept"
			} else {
				ubiFlag = "Reject"
			}
			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			beneficiaryAddress = cpAccount.Beneficiary.BeneficiaryAddress
			chainNodeId = cpAccount.NodeId
			chainMultiAddress = strings.Join(cpAccount.MultiAddresses, ",")
		}

		if ownerAddress != "" {
			ownerBalance, err = wallet.Balance(context.TODO(), client, ownerAddress)
		}

		var taskData [][]string

		taskData = append(taskData, []string{"Multi-Address:", chainMultiAddress})
		taskData = append(taskData, []string{"Node ID:", chainNodeId})
		taskData = append(taskData, []string{"ECP:"})
		taskData = append(taskData, []string{"   Contract Address:", contractAddress})
		taskData = append(taskData, []string{"   UBI FLAG:", ubiFlag})
		taskData = append(taskData, []string{"   Owner:", ownerAddress})
		taskData = append(taskData, []string{"   Beneficiary Address:", beneficiaryAddress})
		taskData = append(taskData, []string{"   Available(SWAN-ETH):", ownerBalance})
		taskData = append(taskData, []string{"   Collateral(SWAN-ETH):", "0"})

		header := []string{"Name:", conf.GetConfig().API.NodeName}
		NewVisualTable(header, taskData, []RowColor{}).Generate(false)
		return nil
	},
}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new cp",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "multi-address",
			Usage: "The multiAddress for libp2p(public ip)",
		},
		&cli.StringFlag{
			Name:  "node-name",
			Usage: "The name of cp",
		},
		&cli.IntFlag{
			Name:  "port",
			Usage: "The cp listens on port, default: 9085",
		},
	},
	Action: func(cctx *cli.Context) error {
		multiAddr := cctx.String("multi-address")
		port := cctx.Int("port")
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("the multi-address field required")
		}
		nodeName := cctx.String("node-name")

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}
		return conf.UpdateConfigFile(cpRepoPath, multiAddr, nodeName, port)
	},
}

var accountCmd = &cli.Command{
	Name:  "account",
	Usage: "Manage account info of CP",
	Subcommands: []*cli.Command{
		createAccountCmd,
		changeMultiAddressCmd,
		changeOwnerAddressCmd,
		changeBeneficiaryAddressCmd,
		changeUbiFlagCmd,
	},
}

var createAccountCmd = &cli.Command{
	Name:  "create",
	Usage: "Create a cp account to chain",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "ownerAddress",
			Usage: "Specify a OwnerAddress",
		},
		&cli.StringFlag{
			Name:  "beneficiaryAddress",
			Usage: "Specify a beneficiaryAddress to receive rewards. If not specified, use the same address as ownerAddress",
		},
		&cli.BoolFlag{
			Name:  "ubi-flag",
			Usage: "Whether to accept the UBI task",
			Value: false,
		},
	},
	Action: func(cctx *cli.Context) error {
		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		beneficiaryAddress := cctx.String("beneficiaryAddress")
		if strings.TrimSpace(beneficiaryAddress) == "" {
			beneficiaryAddress = ownerAddress
		}

		ubiFlag := cctx.Bool("ubi-flag")

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}
		return createAccount(cpRepoPath, ownerAddress, beneficiaryAddress, ubiFlag)
	},
}

var changeMultiAddressCmd = &cli.Command{
	Name:      "changeMultiAddress",
	Usage:     "Update MultiAddress of CP (/ip4/<public_ip>/tcp/<port>)",
	ArgsUsage: "[multiAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a multiAddress")
		}

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		multiAddr := cctx.Args().Get(0)
		if strings.TrimSpace(multiAddr) == "" {
			return fmt.Errorf("failed to parse : %s", multiAddr)
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			logs.GetLogger().Fatal(err)
		}
		return changeMultiAddress(ownerAddress, multiAddr)

	},
}

var changeOwnerAddressCmd = &cli.Command{
	Name:      "changeOwnerAddress",
	Usage:     "Update OwnerAddress of CP",
	ArgsUsage: "[newOwnerAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is required")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a new ownerAddress")
		}

		newOwnerAddr := cctx.Args().Get(0)
		if strings.TrimSpace(newOwnerAddr) == "" {
			return fmt.Errorf("failed to parse : %s", newOwnerAddr)
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the old owner address: %s, private key %v,", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		submitUBIProofTx, err := cpStub.ChangeOwnerAddress(common.HexToAddress(newOwnerAddr))
		if err != nil {
			logs.GetLogger().Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeOwnerAddress: %s \n", submitUBIProofTx)

		return nil
	},
}

var changeBeneficiaryAddressCmd = &cli.Command{
	Name:      "changeBeneficiaryAddress",
	Usage:     "Update beneficiaryAddress of CP",
	ArgsUsage: "[beneficiaryAddress]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		beneficiaryAddress := cctx.Args().Get(0)
		if strings.TrimSpace(beneficiaryAddress) == "" {
			return fmt.Errorf("failed to parse target address: %s", beneficiaryAddress)
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}
		newQuota := big.NewInt(int64(0))
		newExpiration := big.NewInt(int64(0))
		changeBeneficiaryAddressTx, err := cpStub.ChangeBeneficiary(common.HexToAddress(beneficiaryAddress), newQuota, newExpiration)
		if err != nil {
			logs.GetLogger().Errorf("change owner address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("changeBeneficiaryAddress Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeUbiFlagCmd = &cli.Command{
	Name:      "changeUbiFlag",
	Usage:     "Update ubiFlag of CP (0:Reject, 1:Accept)",
	ArgsUsage: "[ubiFlag]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "ownerAddress",
			Usage:    "Specify a OwnerAddress",
			Required: true,
		},
	},
	Action: func(cctx *cli.Context) error {

		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a beneficiaryAddress")
		}

		ubiFlag := cctx.Args().Get(0)
		if strings.TrimSpace(ubiFlag) == "" {
			return fmt.Errorf("ubiFlag is not empty")
		}

		if strings.TrimSpace(ubiFlag) != "0" && strings.TrimSpace(ubiFlag) != "1" {
			return fmt.Errorf("ubiFlag must be 0 or 1")
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			logs.GetLogger().Fatal(err)
		}

		chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
		if err != nil {
			logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
			return err
		}

		localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
		if err != nil {
			logs.GetLogger().Errorf("setup wallet ubi failed, error: %v,", err)
			return err
		}

		ki, err := localWallet.FindKey(ownerAddress)
		if err != nil || ki == nil {
			logs.GetLogger().Errorf("the address: %s, private key %v. Please import the address into the wallet", ownerAddress, wallet.ErrKeyInfoNotFound)
			return err
		}

		client, err := ethclient.Dial(chainUrl)
		if err != nil {
			logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
			return err
		}
		defer client.Close()

		cpStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
		if err != nil {
			logs.GetLogger().Errorf("create cp client failed, error: %v,", err)
			return err
		}

		cpAccount, err := cpStub.GetCpAccountInfo()
		if err != nil {
			return fmt.Errorf("get cpAccount faile, error: %v", err)
		}
		if !strings.EqualFold(cpAccount.OwnerAddress, ownerAddress) {
			return fmt.Errorf("the owner address is incorrect. The owner on the chain is %s, and the current address is %s", cpAccount.OwnerAddress, ownerAddress)
		}

		newUbiFlag, _ := strconv.ParseUint(strings.TrimSpace(ubiFlag), 10, 64)

		changeBeneficiaryAddressTx, err := cpStub.ChangeUbiFlag(uint8(newUbiFlag))
		if err != nil {
			logs.GetLogger().Errorf("change ubi flag tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeUbiFlag Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

func DoSend(contractAddr, height string) error {
	type ContractReq struct {
		Addr   string `json:"addr"`
		Height int    `json:"height"`
	}
	h, _ := strconv.ParseInt(height, 10, 64)
	var contractReq ContractReq
	contractReq.Addr = contractAddr
	contractReq.Height = int(h)

	jsonData, err := json.Marshal(contractReq)
	if err != nil {
		logs.GetLogger().Errorf("JSON encoding failed: %v", err)
		return err
	}

	resp, err := http.Post(conf.GetConfig().UBI.UbiUrl+"/contracts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logs.GetLogger().Errorf("POST request failed: %v", err)
		return err
	}
	defer resp.Body.Close()

	var resultResp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data any    `json:"data,omitempty"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.GetLogger().Errorf("read response failed: %v", err)
		return err
	}
	err = json.Unmarshal(body, &resultResp)
	if err != nil {
		logs.GetLogger().Errorf("response convert to json failed: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("register cp to ubi hub failed, error: %s", resultResp.Msg)
	}
	return nil
}
