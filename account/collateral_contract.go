// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package account

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ECPCollateralCPInfo is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralCPInfo struct {
	Cp            common.Address
	Balance       *big.Int
	FrozenBalance *big.Int
	Status        string
}

// ECPCollateralContractInfo is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralContractInfo struct {
	SlashedFunds    *big.Int
	BaseCollateral  *big.Int
	TaskBalance     *big.Int
	CollateralRatio *big.Int
	SlashRatio      *big.Int
}

// ECPCollateralTask is an auto generated low-level Go binding around an user-defined struct.
type ECPCollateralTask struct {
	CpAccountAddress common.Address
	Collateral       *big.Int
	Status           *big.Int
}

// CollateralMetaData contains all meta data concerning the Collateral contract.
var CollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frozenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operation\",\"type\":\"string\"}],\"name\":\"CollateralAdjusted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"CollateralLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"CollateralSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"CollateralUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fundingWallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"TaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStatus\",\"type\":\"uint256\"}],\"name\":\"TaskStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAddress\",\"type\":\"address\"}],\"name\":\"cpInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"balance\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"frozenBalance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structECPCollateral.CPInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cpStatus\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"frozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getECPCollateralInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slashedFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"taskBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashRatio\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"getTaskInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structECPCollateral.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"lockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCollateral\",\"type\":\"uint256\"}],\"name\":\"setBaseCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_collateralRatio\",\"type\":\"uint256\"}],\"name\":\"setCollateralRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashRatio\",\"type\":\"uint256\"}],\"name\":\"setSlashRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"slashCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"cpAccountAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"taskContractAddress\",\"type\":\"address\"}],\"name\":\"unlockCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cpAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5033600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000885760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016200007f919062000222565b60405180910390fd5b62000099816200011960201b60201c565b50620000ab336200011960201b60201c565b6001600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600560048190555060026005819055506200023f565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200020a82620001dd565b9050919050565b6200021c81620001fd565b82525050565b600060208201905062000239600083018462000211565b92915050565b612ca5806200024f6000396000f3fe6080604052600436106101a05760003560e01c80637f58a7e5116100ec578063ce3518aa1161008a578063f2fde38b11610064578063f2fde38b146105ee578063f31cc88314610617578063f340fa0114610642578063f3fef3a31461065e576101b0565b8063ce3518aa14610571578063d27ca89b1461059a578063e0ed539d146105c5576101b0565b80639b5ddf09116100c65780639b5ddf09146104b55780639d1fef81146104e0578063a664c21614610509578063b4eae1cb14610546576101b0565b80637f58a7e5146104365780638da5cb5b1461045f5780639939cd181461048a576101b0565b80633fe65177116101595780637048027511610133578063704802751461038c57806370b72944146103b5578063715018a6146103e057806377c237fd146103f7576101b0565b80633fe65177146102fb5780636060663e146103385780636f99f15c14610361576101b0565b80630b4b73d1146101b55780631785f53c146101f257806324d7806c1461021b578063266565a91461025857806327e235e3146102955780633f001fd9146102d2576101b0565b366101b0576101ae33610687565b005b600080fd5b3480156101c157600080fd5b506101dc60048036038101906101d79190611f05565b61074e565b6040516101e99190611f9c565b60405180910390f35b3480156101fe57600080fd5b5061021960048036038101906102149190611f05565b610811565b005b34801561022757600080fd5b50610242600480360381019061023d9190611f05565b610874565b60405161024f9190611fd2565b60405180910390f35b34801561026457600080fd5b5061027f600480360381019061027a9190611f05565b610894565b60405161028c9190611ffc565b60405180910390f35b3480156102a157600080fd5b506102bc60048036038101906102b79190611f05565b6108ac565b6040516102c99190612030565b60405180910390f35b3480156102de57600080fd5b506102f960048036038101906102f49190611f05565b6108c4565b005b34801561030757600080fd5b50610322600480360381019061031d9190611f05565b610c0a565b60405161032f91906120db565b60405180910390f35b34801561034457600080fd5b5061035f600480360381019061035a9190612129565b610caa565b005b34801561036d57600080fd5b50610376610cbc565b6040516103839190611ffc565b60405180910390f35b34801561039857600080fd5b506103b360048036038101906103ae9190611f05565b610cc2565b005b3480156103c157600080fd5b506103ca610d25565b6040516103d79190611ffc565b60405180910390f35b3480156103ec57600080fd5b506103f5610d2f565b005b34801561040357600080fd5b5061041e60048036038101906104199190611f05565b610d43565b60405161042d93929190612165565b60405180910390f35b34801561044257600080fd5b5061045d60048036038101906104589190612129565b610d8d565b005b34801561046b57600080fd5b50610474610e23565b604051610481919061219c565b60405180910390f35b34801561049657600080fd5b5061049f610e4c565b6040516104ac919061221f565b60405180910390f35b3480156104c157600080fd5b506104ca610e8a565b6040516104d79190611ffc565b60405180910390f35b3480156104ec57600080fd5b5061050760048036038101906105029190611f05565b610e90565b005b34801561051557600080fd5b50610530600480360381019061052b9190611f05565b611292565b60405161053d91906122f6565b60405180910390f35b34801561055257600080fd5b5061055b61141d565b6040516105689190611ffc565b60405180910390f35b34801561057d57600080fd5b5061059860048036038101906105939190612129565b611423565b005b3480156105a657600080fd5b506105af611435565b6040516105bc9190611ffc565b60405180910390f35b3480156105d157600080fd5b506105ec60048036038101906105e79190612318565b61143b565b005b3480156105fa57600080fd5b5061061560048036038101906106109190611f05565b611772565b005b34801561062357600080fd5b5061062c6117f8565b6040516106399190611ffc565b60405180910390f35b61065c60048036038101906106579190611f05565b610687565b005b34801561066a57600080fd5b506106856004803603810190610680919061236b565b6117fe565b005b61069081611b49565b34600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106df91906123da565b925050819055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62346040516107439190611ffc565b60405180910390a350565b610756611dfe565b600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815250509050919050565b610819611cab565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b60066020528060005260406000206000915054906101000a900460ff1681565b60086020528060005260406000206000915090505481565b60076020528060005260406000206000915090505481565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610950576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161094790612490565b60405180910390fd5b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081836001015411610a12578260010154610a14565b815b905080600860008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a8991906124b0565b9250508190555080600760008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b0391906123da565b9250508190555060028360020181905550610b418360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b49565b8260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f817d69f45cb0984f7141c288a3bea0d2dd7bc075758f311056d0aa31798614fa8286604051610bad9291906124e4565b60405180910390a28373ffffffffffffffffffffffffffffffffffffffff167f43a64df649a51c960307a3b5a4ec25a67c76629f0881669a214d162ae4ad2b5f6002604051610bfc9190611ffc565b60405180910390a250505050565b600a6020528060005260406000206000915090508054610c299061253c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c559061253c565b8015610ca25780601f10610c7757610100808354040283529160200191610ca2565b820191906000526020600020905b815481529060010190602001808311610c8557829003601f168201915b505050505081565b610cb2611cab565b8060048190555050565b60015481565b610cca611cab565b6001600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000600254905090565b610d37611cab565b610d416000611d32565b565b60096020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154905083565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610e19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1090612490565b60405180910390fd5b8060028190555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610e54611e35565b6040518060a001604052806001548152602001600254815260200160035481526020016004548152602001600554815250905090565b60025481565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610f1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1390612490565b60405180910390fd5b6000600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060006005548260010154610f73919061256d565b90506000600860008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000818311610fec5782610fee565b815b9050600081841161100057600061100d565b818461100c91906124b0565b5b905081600860008760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461108291906124b0565b9250508190555080600760008760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546110fc91906125af565b92505081905550836001600082825461111591906125f2565b92505081905550600385600201819055506111538560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16611b49565b8460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f0f51f32e6325a6b213a7e9459df2a4564f058d9dca8309ff9b2508f6a83cf59385886040516111bf9291906124e4565b60405180910390a28573ffffffffffffffffffffffffffffffffffffffff167f43a64df649a51c960307a3b5a4ec25a67c76629f0881669a214d162ae4ad2b5f600360405161120e9190611ffc565b60405180910390a28460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e48383604051611282929190612672565b60405180910390a2505050505050565b61129a611e64565b60405180608001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546113959061253c565b80601f01602080910402602001604051908101604052809291908181526020018280546113c19061253c565b801561140e5780601f106113e35761010080835404028352916020019161140e565b820191906000526020600020905b8154815290600101906020018083116113f157829003601f168201915b50505050508152509050919050565b60045481565b61142b611cab565b8060058190555050565b60055481565b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166114c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114be90612490565b60405180910390fd5b81600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541215611549576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154090612720565b60405180910390fd5b81600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461159891906125af565b9250508190555081600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546115ee91906125f2565b9250508190555060405180606001604052808473ffffffffffffffffffffffffffffffffffffffff1681526020018381526020016001815250600960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050506116cd83611b49565b8273ffffffffffffffffffffffffffffffffffffffff167f56b60085244fe1afa467f92231debadc8bcb6127e9dbf35dbb58ca8406a423fe83836040516117159291906124e4565b60405180910390a28073ffffffffffffffffffffffffffffffffffffffff167f3f9b921f6bb5d577cdf49122202d27f7c50e3cc981c47543dc4a59c00dee52298484604051611765929190612740565b60405180910390a2505050565b61177a611cab565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036117ec5760006040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016117e3919061219c565b60405180910390fd5b6117f581611d32565b50565b60035481565b61180782611b49565b6000808373ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f893d20e8000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516118b191906127b0565b6000604051808303816000865af19150503d80600081146118ee576040519150601f19603f3d011682016040523d82523d6000602084013e6118f3565b606091505b509150915081611938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161192f90612839565b60405180910390fd5b60008180602001905181019061194e9190612897565b905083600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412156119d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119c990612910565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611a40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a37906129a2565b60405180910390fd5b83600760008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611a8f91906125af565b925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f19350505050158015611adc573d6000803e3d6000fd5b508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb86604051611b3a9190611ffc565b60405180910390a35050505050565b600254600454611b59919061256d565b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205412611c25576040518060400160405280600981526020017f7a6b41756374696f6e0000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081611c1f9190612b9d565b50611ca8565b6040518060400160405280600381526020017f4e53430000000000000000000000000000000000000000000000000000000000815250600a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209081611ca69190612b9d565b505b50565b611cb3611df6565b73ffffffffffffffffffffffffffffffffffffffff16611cd1610e23565b73ffffffffffffffffffffffffffffffffffffffff1614611d3057611cf4611df6565b6040517f118cdaa7000000000000000000000000000000000000000000000000000000008152600401611d27919061219c565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001606081525090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611ed282611ea7565b9050919050565b611ee281611ec7565b8114611eed57600080fd5b50565b600081359050611eff81611ed9565b92915050565b600060208284031215611f1b57611f1a611ea2565b5b6000611f2984828501611ef0565b91505092915050565b611f3b81611ec7565b82525050565b6000819050919050565b611f5481611f41565b82525050565b606082016000820151611f706000850182611f32565b506020820151611f836020850182611f4b565b506040820151611f966040850182611f4b565b50505050565b6000606082019050611fb16000830184611f5a565b92915050565b60008115159050919050565b611fcc81611fb7565b82525050565b6000602082019050611fe76000830184611fc3565b92915050565b611ff681611f41565b82525050565b60006020820190506120116000830184611fed565b92915050565b6000819050919050565b61202a81612017565b82525050565b60006020820190506120456000830184612021565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561208557808201518184015260208101905061206a565b60008484015250505050565b6000601f19601f8301169050919050565b60006120ad8261204b565b6120b78185612056565b93506120c7818560208601612067565b6120d081612091565b840191505092915050565b600060208201905081810360008301526120f581846120a2565b905092915050565b61210681611f41565b811461211157600080fd5b50565b600081359050612123816120fd565b92915050565b60006020828403121561213f5761213e611ea2565b5b600061214d84828501612114565b91505092915050565b61215f81611ec7565b82525050565b600060608201905061217a6000830186612156565b6121876020830185611fed565b6121946040830184611fed565b949350505050565b60006020820190506121b16000830184612156565b92915050565b60a0820160008201516121cd6000850182611f4b565b5060208201516121e06020850182611f4b565b5060408201516121f36040850182611f4b565b5060608201516122066060850182611f4b565b5060808201516122196080850182611f4b565b50505050565b600060a08201905061223460008301846121b7565b92915050565b61224381612017565b82525050565b600082825260208201905092915050565b60006122658261204b565b61226f8185612249565b935061227f818560208601612067565b61228881612091565b840191505092915050565b60006080830160008301516122ab6000860182611f32565b5060208301516122be602086018261223a565b5060408301516122d16040860182611f4b565b50606083015184820360608601526122e9828261225a565b9150508091505092915050565b600060208201905081810360008301526123108184612293565b905092915050565b60008060006060848603121561233157612330611ea2565b5b600061233f86828701611ef0565b935050602061235086828701612114565b925050604061236186828701611ef0565b9150509250925092565b6000806040838503121561238257612381611ea2565b5b600061239085828601611ef0565b92505060206123a185828601612114565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006123e582612017565b91506123f083612017565b925082820190508281121560008312168382126000841215161715612418576124176123ab565b5b92915050565b7f4f6e6c79207468652061646d696e2063616e2063616c6c20746869732066756e60008201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b600061247a602683612056565b91506124858261241e565b604082019050919050565b600060208201905081810360008301526124a98161246d565b9050919050565b60006124bb82611f41565b91506124c683611f41565b92508282039050818111156124de576124dd6123ab565b5b92915050565b60006040820190506124f96000830185611fed565b6125066020830184612156565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061255457607f821691505b6020821081036125675761256661250d565b5b50919050565b600061257882611f41565b915061258383611f41565b925082820261259181611f41565b915082820484148315176125a8576125a76123ab565b5b5092915050565b60006125ba82612017565b91506125c583612017565b92508282039050818112600084121682821360008512151617156125ec576125eb6123ab565b5b92915050565b60006125fd82611f41565b915061260883611f41565b92508282019050808211156126205761261f6123ab565b5b92915050565b7f536c617368656400000000000000000000000000000000000000000000000000600082015250565b600061265c600783612056565b915061266782612626565b602082019050919050565b60006060820190506126876000830185611fed565b6126946020830184611fed565b81810360408301526126a58161264f565b90509392505050565b7f4e6f7420656e6f7567682062616c616e636520666f7220636f6c6c617465726160008201527f6c00000000000000000000000000000000000000000000000000000000000000602082015250565b600061270a602183612056565b9150612715826126ae565b604082019050919050565b60006020820190508181036000830152612739816126fd565b9050919050565b60006040820190506127556000830185612156565b6127626020830184611fed565b9392505050565b600081519050919050565b600081905092915050565b600061278a82612769565b6127948185612774565b93506127a4818560208601612067565b80840191505092915050565b60006127bc828461277f565b915081905092915050565b7f4661696c656420746f2063616c6c206765744f776e65722066756e6374696f6e60008201527f206f662043504163636f756e7400000000000000000000000000000000000000602082015250565b6000612823602d83612056565b915061282e826127c7565b604082019050919050565b6000602082019050818103600083015261285281612816565b9050919050565b600061286482611ea7565b9050919050565b61287481612859565b811461287f57600080fd5b50565b6000815190506128918161286b565b92915050565b6000602082840312156128ad576128ac611ea2565b5b60006128bb84828501612882565b91505092915050565b7f576974686472617720616d6f756e7420657863656564732062616c616e636500600082015250565b60006128fa601f83612056565b9150612905826128c4565b602082019050919050565b60006020820190508181036000830152612929816128ed565b9050919050565b7f4f6e6c792043502773206f776e65722063616e2077697468647261772074686560008201527f20636f6c6c61746572616c2066756e6473000000000000000000000000000000602082015250565b600061298c603183612056565b915061299782612930565b604082019050919050565b600060208201905081810360008301526129bb8161297f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302612a537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612a16565b612a5d8683612a16565b95508019841693508086168417925050509392505050565b6000819050919050565b6000612a9a612a95612a9084611f41565b612a75565b611f41565b9050919050565b6000819050919050565b612ab483612a7f565b612ac8612ac082612aa1565b848454612a23565b825550505050565b600090565b612add612ad0565b612ae8818484612aab565b505050565b5b81811015612b0c57612b01600082612ad5565b600181019050612aee565b5050565b601f821115612b5157612b22816129f1565b612b2b84612a06565b81016020851015612b3a578190505b612b4e612b4685612a06565b830182612aed565b50505b505050565b600082821c905092915050565b6000612b7460001984600802612b56565b1980831691505092915050565b6000612b8d8383612b63565b9150826002028217905092915050565b612ba68261204b565b67ffffffffffffffff811115612bbf57612bbe6129c2565b5b612bc9825461253c565b612bd4828285612b10565b600060209050601f831160018114612c075760008415612bf5578287015190505b612bff8582612b81565b865550612c67565b601f198416612c15866129f1565b60005b82811015612c3d57848901518255600182019150602085019450602081019050612c18565b86831015612c5a5784890151612c56601f891682612b63565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220f6fc876c8cceefb175d874063e75d24a5bc5f3684a2fa35df9a68881ccc1d52b64736f6c63430008140033",
}

// CollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use CollateralMetaData.ABI instead.
var CollateralABI = CollateralMetaData.ABI

// CollateralBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CollateralMetaData.Bin instead.
var CollateralBin = CollateralMetaData.Bin

// DeployCollateral deploys a new Ethereum contract, binding an instance of Collateral to it.
func DeployCollateral(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Collateral, error) {
	parsed, err := CollateralMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CollateralBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Collateral{CollateralCaller: CollateralCaller{contract: contract}, CollateralTransactor: CollateralTransactor{contract: contract}, CollateralFilterer: CollateralFilterer{contract: contract}}, nil
}

// Collateral is an auto generated Go binding around an Ethereum contract.
type Collateral struct {
	CollateralCaller     // Read-only binding to the contract
	CollateralTransactor // Write-only binding to the contract
	CollateralFilterer   // Log filterer for contract events
}

// CollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollateralSession struct {
	Contract     *Collateral       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollateralCallerSession struct {
	Contract *CollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollateralTransactorSession struct {
	Contract     *CollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollateralRaw struct {
	Contract *Collateral // Generic contract binding to access the raw methods on
}

// CollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollateralCallerRaw struct {
	Contract *CollateralCaller // Generic read-only contract binding to access the raw methods on
}

// CollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollateralTransactorRaw struct {
	Contract *CollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollateral creates a new instance of Collateral, bound to a specific deployed contract.
func NewCollateral(address common.Address, backend bind.ContractBackend) (*Collateral, error) {
	contract, err := bindCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Collateral{CollateralCaller: CollateralCaller{contract: contract}, CollateralTransactor: CollateralTransactor{contract: contract}, CollateralFilterer: CollateralFilterer{contract: contract}}, nil
}

// NewCollateralCaller creates a new read-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralCaller(address common.Address, caller bind.ContractCaller) (*CollateralCaller, error) {
	contract, err := bindCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralCaller{contract: contract}, nil
}

// NewCollateralTransactor creates a new write-only instance of Collateral, bound to a specific deployed contract.
func NewCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*CollateralTransactor, error) {
	contract, err := bindCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollateralTransactor{contract: contract}, nil
}

// NewCollateralFilterer creates a new log filterer instance of Collateral, bound to a specific deployed contract.
func NewCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*CollateralFilterer, error) {
	contract, err := bindCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollateralFilterer{contract: contract}, nil
}

// bindCollateral binds a generic wrapper to an already deployed contract.
func bindCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.CollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.CollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Collateral *CollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Collateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Collateral *CollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Collateral *CollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Collateral.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_Collateral *CollateralCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_Collateral *CollateralSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Balances(&_Collateral.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(int256)
func (_Collateral *CollateralCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.Balances(&_Collateral.CallOpts, arg0)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_Collateral *CollateralCaller) BaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "baseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_Collateral *CollateralSession) BaseCollateral() (*big.Int, error) {
	return _Collateral.Contract.BaseCollateral(&_Collateral.CallOpts)
}

// BaseCollateral is a free data retrieval call binding the contract method 0x9b5ddf09.
//
// Solidity: function baseCollateral() view returns(uint256)
func (_Collateral *CollateralCallerSession) BaseCollateral() (*big.Int, error) {
	return _Collateral.Contract.BaseCollateral(&_Collateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_Collateral *CollateralCaller) CollateralRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "collateralRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_Collateral *CollateralSession) CollateralRatio() (*big.Int, error) {
	return _Collateral.Contract.CollateralRatio(&_Collateral.CallOpts)
}

// CollateralRatio is a free data retrieval call binding the contract method 0xb4eae1cb.
//
// Solidity: function collateralRatio() view returns(uint256)
func (_Collateral *CollateralCallerSession) CollateralRatio() (*big.Int, error) {
	return _Collateral.Contract.CollateralRatio(&_Collateral.CallOpts)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_Collateral *CollateralCaller) CpInfo(opts *bind.CallOpts, cpAddress common.Address) (ECPCollateralCPInfo, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "cpInfo", cpAddress)

	if err != nil {
		return *new(ECPCollateralCPInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralCPInfo)).(*ECPCollateralCPInfo)

	return out0, err

}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_Collateral *CollateralSession) CpInfo(cpAddress common.Address) (ECPCollateralCPInfo, error) {
	return _Collateral.Contract.CpInfo(&_Collateral.CallOpts, cpAddress)
}

// CpInfo is a free data retrieval call binding the contract method 0xa664c216.
//
// Solidity: function cpInfo(address cpAddress) view returns((address,int256,uint256,string))
func (_Collateral *CollateralCallerSession) CpInfo(cpAddress common.Address) (ECPCollateralCPInfo, error) {
	return _Collateral.Contract.CpInfo(&_Collateral.CallOpts, cpAddress)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_Collateral *CollateralCaller) CpStatus(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "cpStatus", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_Collateral *CollateralSession) CpStatus(arg0 common.Address) (string, error) {
	return _Collateral.Contract.CpStatus(&_Collateral.CallOpts, arg0)
}

// CpStatus is a free data retrieval call binding the contract method 0x3fe65177.
//
// Solidity: function cpStatus(address ) view returns(string)
func (_Collateral *CollateralCallerSession) CpStatus(arg0 common.Address) (string, error) {
	return _Collateral.Contract.CpStatus(&_Collateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCaller) FrozenBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "frozenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.FrozenBalance(&_Collateral.CallOpts, arg0)
}

// FrozenBalance is a free data retrieval call binding the contract method 0x266565a9.
//
// Solidity: function frozenBalance(address ) view returns(uint256)
func (_Collateral *CollateralCallerSession) FrozenBalance(arg0 common.Address) (*big.Int, error) {
	return _Collateral.Contract.FrozenBalance(&_Collateral.CallOpts, arg0)
}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_Collateral *CollateralCaller) GetBaseCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "getBaseCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_Collateral *CollateralSession) GetBaseCollateral() (*big.Int, error) {
	return _Collateral.Contract.GetBaseCollateral(&_Collateral.CallOpts)
}

// GetBaseCollateral is a free data retrieval call binding the contract method 0x70b72944.
//
// Solidity: function getBaseCollateral() view returns(uint256)
func (_Collateral *CollateralCallerSession) GetBaseCollateral() (*big.Int, error) {
	return _Collateral.Contract.GetBaseCollateral(&_Collateral.CallOpts)
}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Collateral *CollateralCaller) GetECPCollateralInfo(opts *bind.CallOpts) (ECPCollateralContractInfo, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "getECPCollateralInfo")

	if err != nil {
		return *new(ECPCollateralContractInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralContractInfo)).(*ECPCollateralContractInfo)

	return out0, err

}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Collateral *CollateralSession) GetECPCollateralInfo() (ECPCollateralContractInfo, error) {
	return _Collateral.Contract.GetECPCollateralInfo(&_Collateral.CallOpts)
}

// GetECPCollateralInfo is a free data retrieval call binding the contract method 0x9939cd18.
//
// Solidity: function getECPCollateralInfo() view returns((uint256,uint256,uint256,uint256,uint256))
func (_Collateral *CollateralCallerSession) GetECPCollateralInfo() (ECPCollateralContractInfo, error) {
	return _Collateral.Contract.GetECPCollateralInfo(&_Collateral.CallOpts)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x0b4b73d1.
//
// Solidity: function getTaskInfo(address taskContractAddress) view returns((address,uint256,uint256))
func (_Collateral *CollateralCaller) GetTaskInfo(opts *bind.CallOpts, taskContractAddress common.Address) (ECPCollateralTask, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "getTaskInfo", taskContractAddress)

	if err != nil {
		return *new(ECPCollateralTask), err
	}

	out0 := *abi.ConvertType(out[0], new(ECPCollateralTask)).(*ECPCollateralTask)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0x0b4b73d1.
//
// Solidity: function getTaskInfo(address taskContractAddress) view returns((address,uint256,uint256))
func (_Collateral *CollateralSession) GetTaskInfo(taskContractAddress common.Address) (ECPCollateralTask, error) {
	return _Collateral.Contract.GetTaskInfo(&_Collateral.CallOpts, taskContractAddress)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0x0b4b73d1.
//
// Solidity: function getTaskInfo(address taskContractAddress) view returns((address,uint256,uint256))
func (_Collateral *CollateralCallerSession) GetTaskInfo(taskContractAddress common.Address) (ECPCollateralTask, error) {
	return _Collateral.Contract.GetTaskInfo(&_Collateral.CallOpts, taskContractAddress)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Collateral.Contract.IsAdmin(&_Collateral.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_Collateral *CollateralCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _Collateral.Contract.IsAdmin(&_Collateral.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Collateral *CollateralCallerSession) Owner() (common.Address, error) {
	return _Collateral.Contract.Owner(&_Collateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_Collateral *CollateralCaller) SlashRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "slashRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_Collateral *CollateralSession) SlashRatio() (*big.Int, error) {
	return _Collateral.Contract.SlashRatio(&_Collateral.CallOpts)
}

// SlashRatio is a free data retrieval call binding the contract method 0xd27ca89b.
//
// Solidity: function slashRatio() view returns(uint256)
func (_Collateral *CollateralCallerSession) SlashRatio() (*big.Int, error) {
	return _Collateral.Contract.SlashRatio(&_Collateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_Collateral *CollateralCaller) SlashedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "slashedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_Collateral *CollateralSession) SlashedFunds() (*big.Int, error) {
	return _Collateral.Contract.SlashedFunds(&_Collateral.CallOpts)
}

// SlashedFunds is a free data retrieval call binding the contract method 0x6f99f15c.
//
// Solidity: function slashedFunds() view returns(uint256)
func (_Collateral *CollateralCallerSession) SlashedFunds() (*big.Int, error) {
	return _Collateral.Contract.SlashedFunds(&_Collateral.CallOpts)
}

// TaskBalance is a free data retrieval call binding the contract method 0xf31cc883.
//
// Solidity: function taskBalance() view returns(uint256)
func (_Collateral *CollateralCaller) TaskBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "taskBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaskBalance is a free data retrieval call binding the contract method 0xf31cc883.
//
// Solidity: function taskBalance() view returns(uint256)
func (_Collateral *CollateralSession) TaskBalance() (*big.Int, error) {
	return _Collateral.Contract.TaskBalance(&_Collateral.CallOpts)
}

// TaskBalance is a free data retrieval call binding the contract method 0xf31cc883.
//
// Solidity: function taskBalance() view returns(uint256)
func (_Collateral *CollateralCallerSession) TaskBalance() (*big.Int, error) {
	return _Collateral.Contract.TaskBalance(&_Collateral.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x77c237fd.
//
// Solidity: function tasks(address ) view returns(address cpAccountAddress, uint256 collateral, uint256 status)
func (_Collateral *CollateralCaller) Tasks(opts *bind.CallOpts, arg0 common.Address) (struct {
	CpAccountAddress common.Address
	Collateral       *big.Int
	Status           *big.Int
}, error) {
	var out []interface{}
	err := _Collateral.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		CpAccountAddress common.Address
		Collateral       *big.Int
		Status           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CpAccountAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x77c237fd.
//
// Solidity: function tasks(address ) view returns(address cpAccountAddress, uint256 collateral, uint256 status)
func (_Collateral *CollateralSession) Tasks(arg0 common.Address) (struct {
	CpAccountAddress common.Address
	Collateral       *big.Int
	Status           *big.Int
}, error) {
	return _Collateral.Contract.Tasks(&_Collateral.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x77c237fd.
//
// Solidity: function tasks(address ) view returns(address cpAccountAddress, uint256 collateral, uint256 status)
func (_Collateral *CollateralCallerSession) Tasks(arg0 common.Address) (struct {
	CpAccountAddress common.Address
	Collateral       *big.Int
	Status           *big.Int
}, error) {
	return _Collateral.Contract.Tasks(&_Collateral.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralTransactor) AddAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "addAdmin", newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.AddAdmin(&_Collateral.TransactOpts, newAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address newAdmin) returns()
func (_Collateral *CollateralTransactorSession) AddAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.AddAdmin(&_Collateral.TransactOpts, newAdmin)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_Collateral *CollateralTransactor) Deposit(opts *bind.TransactOpts, cpAccount common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "deposit", cpAccount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_Collateral *CollateralSession) Deposit(cpAccount common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, cpAccount)
}

// Deposit is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address cpAccount) payable returns()
func (_Collateral *CollateralTransactorSession) Deposit(cpAccount common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.Deposit(&_Collateral.TransactOpts, cpAccount)
}

// LockCollateral is a paid mutator transaction binding the contract method 0xe0ed539d.
//
// Solidity: function lockCollateral(address cp, uint256 collateral, address taskContractAddress) returns()
func (_Collateral *CollateralTransactor) LockCollateral(opts *bind.TransactOpts, cp common.Address, collateral *big.Int, taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "lockCollateral", cp, collateral, taskContractAddress)
}

// LockCollateral is a paid mutator transaction binding the contract method 0xe0ed539d.
//
// Solidity: function lockCollateral(address cp, uint256 collateral, address taskContractAddress) returns()
func (_Collateral *CollateralSession) LockCollateral(cp common.Address, collateral *big.Int, taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.LockCollateral(&_Collateral.TransactOpts, cp, collateral, taskContractAddress)
}

// LockCollateral is a paid mutator transaction binding the contract method 0xe0ed539d.
//
// Solidity: function lockCollateral(address cp, uint256 collateral, address taskContractAddress) returns()
func (_Collateral *CollateralTransactorSession) LockCollateral(cp common.Address, collateral *big.Int, taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.LockCollateral(&_Collateral.TransactOpts, cp, collateral, taskContractAddress)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.RemoveAdmin(&_Collateral.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Collateral *CollateralTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.RemoveAdmin(&_Collateral.TransactOpts, admin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collateral.Contract.RenounceOwnership(&_Collateral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Collateral *CollateralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Collateral.Contract.RenounceOwnership(&_Collateral.TransactOpts)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_Collateral *CollateralTransactor) SetBaseCollateral(opts *bind.TransactOpts, _baseCollateral *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "setBaseCollateral", _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_Collateral *CollateralSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetBaseCollateral(&_Collateral.TransactOpts, _baseCollateral)
}

// SetBaseCollateral is a paid mutator transaction binding the contract method 0x7f58a7e5.
//
// Solidity: function setBaseCollateral(uint256 _baseCollateral) returns()
func (_Collateral *CollateralTransactorSession) SetBaseCollateral(_baseCollateral *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetBaseCollateral(&_Collateral.TransactOpts, _baseCollateral)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_Collateral *CollateralTransactor) SetCollateralRatio(opts *bind.TransactOpts, _collateralRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "setCollateralRatio", _collateralRatio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_Collateral *CollateralSession) SetCollateralRatio(_collateralRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetCollateralRatio(&_Collateral.TransactOpts, _collateralRatio)
}

// SetCollateralRatio is a paid mutator transaction binding the contract method 0x6060663e.
//
// Solidity: function setCollateralRatio(uint256 _collateralRatio) returns()
func (_Collateral *CollateralTransactorSession) SetCollateralRatio(_collateralRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetCollateralRatio(&_Collateral.TransactOpts, _collateralRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_Collateral *CollateralTransactor) SetSlashRatio(opts *bind.TransactOpts, _slashRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "setSlashRatio", _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_Collateral *CollateralSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetSlashRatio(&_Collateral.TransactOpts, _slashRatio)
}

// SetSlashRatio is a paid mutator transaction binding the contract method 0xce3518aa.
//
// Solidity: function setSlashRatio(uint256 _slashRatio) returns()
func (_Collateral *CollateralTransactorSession) SetSlashRatio(_slashRatio *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.SetSlashRatio(&_Collateral.TransactOpts, _slashRatio)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x9d1fef81.
//
// Solidity: function slashCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralTransactor) SlashCollateral(opts *bind.TransactOpts, taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "slashCollateral", taskContractAddress)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x9d1fef81.
//
// Solidity: function slashCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralSession) SlashCollateral(taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.SlashCollateral(&_Collateral.TransactOpts, taskContractAddress)
}

// SlashCollateral is a paid mutator transaction binding the contract method 0x9d1fef81.
//
// Solidity: function slashCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralTransactorSession) SlashCollateral(taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.SlashCollateral(&_Collateral.TransactOpts, taskContractAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Collateral *CollateralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.TransferOwnership(&_Collateral.TransactOpts, newOwner)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralTransactor) UnlockCollateral(opts *bind.TransactOpts, taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "unlockCollateral", taskContractAddress)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralSession) UnlockCollateral(taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UnlockCollateral(&_Collateral.TransactOpts, taskContractAddress)
}

// UnlockCollateral is a paid mutator transaction binding the contract method 0x3f001fd9.
//
// Solidity: function unlockCollateral(address taskContractAddress) returns()
func (_Collateral *CollateralTransactorSession) UnlockCollateral(taskContractAddress common.Address) (*types.Transaction, error) {
	return _Collateral.Contract.UnlockCollateral(&_Collateral.TransactOpts, taskContractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_Collateral *CollateralTransactor) Withdraw(opts *bind.TransactOpts, cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.contract.Transact(opts, "withdraw", cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_Collateral *CollateralSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, cpAccount, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address cpAccount, uint256 amount) returns()
func (_Collateral *CollateralTransactorSession) Withdraw(cpAccount common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Collateral.Contract.Withdraw(&_Collateral.TransactOpts, cpAccount, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Collateral.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralSession) Receive() (*types.Transaction, error) {
	return _Collateral.Contract.Receive(&_Collateral.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Collateral *CollateralTransactorSession) Receive() (*types.Transaction, error) {
	return _Collateral.Contract.Receive(&_Collateral.TransactOpts)
}

// CollateralCollateralAdjustedIterator is returned from FilterCollateralAdjusted and is used to iterate over the raw logs and unpacked data for CollateralAdjusted events raised by the Collateral contract.
type CollateralCollateralAdjustedIterator struct {
	Event *CollateralCollateralAdjusted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralCollateralAdjustedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralCollateralAdjusted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralCollateralAdjusted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralCollateralAdjustedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralCollateralAdjustedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralCollateralAdjusted represents a CollateralAdjusted event raised by the Collateral contract.
type CollateralCollateralAdjusted struct {
	Cp            common.Address
	FrozenAmount  *big.Int
	BalanceAmount *big.Int
	Operation     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralAdjusted is a free log retrieval operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_Collateral *CollateralFilterer) FilterCollateralAdjusted(opts *bind.FilterOpts, cp []common.Address) (*CollateralCollateralAdjustedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return &CollateralCollateralAdjustedIterator{contract: _Collateral.contract, event: "CollateralAdjusted", logs: logs, sub: sub}, nil
}

// WatchCollateralAdjusted is a free log subscription operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_Collateral *CollateralFilterer) WatchCollateralAdjusted(opts *bind.WatchOpts, sink chan<- *CollateralCollateralAdjusted, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "CollateralAdjusted", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralCollateralAdjusted)
				if err := _Collateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralAdjusted is a log parse operation binding the contract event 0x42f1a8a1aee108e84e4eabfaa8d51d7baaa1a02d482295297883a44b2debd3e4.
//
// Solidity: event CollateralAdjusted(address indexed cp, uint256 frozenAmount, uint256 balanceAmount, string operation)
func (_Collateral *CollateralFilterer) ParseCollateralAdjusted(log types.Log) (*CollateralCollateralAdjusted, error) {
	event := new(CollateralCollateralAdjusted)
	if err := _Collateral.contract.UnpackLog(event, "CollateralAdjusted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralCollateralLockedIterator is returned from FilterCollateralLocked and is used to iterate over the raw logs and unpacked data for CollateralLocked events raised by the Collateral contract.
type CollateralCollateralLockedIterator struct {
	Event *CollateralCollateralLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralCollateralLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralCollateralLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralCollateralLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralCollateralLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralCollateralLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralCollateralLocked represents a CollateralLocked event raised by the Collateral contract.
type CollateralCollateralLocked struct {
	Cp                  common.Address
	CollateralAmount    *big.Int
	TaskContractAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollateralLocked is a free log retrieval operation binding the contract event 0x56b60085244fe1afa467f92231debadc8bcb6127e9dbf35dbb58ca8406a423fe.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) FilterCollateralLocked(opts *bind.FilterOpts, cp []common.Address) (*CollateralCollateralLockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &CollateralCollateralLockedIterator{contract: _Collateral.contract, event: "CollateralLocked", logs: logs, sub: sub}, nil
}

// WatchCollateralLocked is a free log subscription operation binding the contract event 0x56b60085244fe1afa467f92231debadc8bcb6127e9dbf35dbb58ca8406a423fe.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) WatchCollateralLocked(opts *bind.WatchOpts, sink chan<- *CollateralCollateralLocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "CollateralLocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralCollateralLocked)
				if err := _Collateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralLocked is a log parse operation binding the contract event 0x56b60085244fe1afa467f92231debadc8bcb6127e9dbf35dbb58ca8406a423fe.
//
// Solidity: event CollateralLocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) ParseCollateralLocked(log types.Log) (*CollateralCollateralLocked, error) {
	event := new(CollateralCollateralLocked)
	if err := _Collateral.contract.UnpackLog(event, "CollateralLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralCollateralSlashedIterator is returned from FilterCollateralSlashed and is used to iterate over the raw logs and unpacked data for CollateralSlashed events raised by the Collateral contract.
type CollateralCollateralSlashedIterator struct {
	Event *CollateralCollateralSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralCollateralSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralCollateralSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralCollateralSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralCollateralSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralCollateralSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralCollateralSlashed represents a CollateralSlashed event raised by the Collateral contract.
type CollateralCollateralSlashed struct {
	Cp                  common.Address
	Amount              *big.Int
	TaskContractAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollateralSlashed is a free log retrieval operation binding the contract event 0x0f51f32e6325a6b213a7e9459df2a4564f058d9dca8309ff9b2508f6a83cf593.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, address taskContractAddress)
func (_Collateral *CollateralFilterer) FilterCollateralSlashed(opts *bind.FilterOpts, cp []common.Address) (*CollateralCollateralSlashedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return &CollateralCollateralSlashedIterator{contract: _Collateral.contract, event: "CollateralSlashed", logs: logs, sub: sub}, nil
}

// WatchCollateralSlashed is a free log subscription operation binding the contract event 0x0f51f32e6325a6b213a7e9459df2a4564f058d9dca8309ff9b2508f6a83cf593.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, address taskContractAddress)
func (_Collateral *CollateralFilterer) WatchCollateralSlashed(opts *bind.WatchOpts, sink chan<- *CollateralCollateralSlashed, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "CollateralSlashed", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralCollateralSlashed)
				if err := _Collateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralSlashed is a log parse operation binding the contract event 0x0f51f32e6325a6b213a7e9459df2a4564f058d9dca8309ff9b2508f6a83cf593.
//
// Solidity: event CollateralSlashed(address indexed cp, uint256 amount, address taskContractAddress)
func (_Collateral *CollateralFilterer) ParseCollateralSlashed(log types.Log) (*CollateralCollateralSlashed, error) {
	event := new(CollateralCollateralSlashed)
	if err := _Collateral.contract.UnpackLog(event, "CollateralSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralCollateralUnlockedIterator is returned from FilterCollateralUnlocked and is used to iterate over the raw logs and unpacked data for CollateralUnlocked events raised by the Collateral contract.
type CollateralCollateralUnlockedIterator struct {
	Event *CollateralCollateralUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralCollateralUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralCollateralUnlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralCollateralUnlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralCollateralUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralCollateralUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralCollateralUnlocked represents a CollateralUnlocked event raised by the Collateral contract.
type CollateralCollateralUnlocked struct {
	Cp                  common.Address
	CollateralAmount    *big.Int
	TaskContractAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollateralUnlocked is a free log retrieval operation binding the contract event 0x817d69f45cb0984f7141c288a3bea0d2dd7bc075758f311056d0aa31798614fa.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) FilterCollateralUnlocked(opts *bind.FilterOpts, cp []common.Address) (*CollateralCollateralUnlockedIterator, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return &CollateralCollateralUnlockedIterator{contract: _Collateral.contract, event: "CollateralUnlocked", logs: logs, sub: sub}, nil
}

// WatchCollateralUnlocked is a free log subscription operation binding the contract event 0x817d69f45cb0984f7141c288a3bea0d2dd7bc075758f311056d0aa31798614fa.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) WatchCollateralUnlocked(opts *bind.WatchOpts, sink chan<- *CollateralCollateralUnlocked, cp []common.Address) (event.Subscription, error) {

	var cpRule []interface{}
	for _, cpItem := range cp {
		cpRule = append(cpRule, cpItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "CollateralUnlocked", cpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralCollateralUnlocked)
				if err := _Collateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralUnlocked is a log parse operation binding the contract event 0x817d69f45cb0984f7141c288a3bea0d2dd7bc075758f311056d0aa31798614fa.
//
// Solidity: event CollateralUnlocked(address indexed cp, uint256 collateralAmount, address taskContractAddress)
func (_Collateral *CollateralFilterer) ParseCollateralUnlocked(log types.Log) (*CollateralCollateralUnlocked, error) {
	event := new(CollateralCollateralUnlocked)
	if err := _Collateral.contract.UnpackLog(event, "CollateralUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Collateral contract.
type CollateralDepositIterator struct {
	Event *CollateralDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralDeposit represents a Deposit event raised by the Collateral contract.
type CollateralDeposit struct {
	FundingWallet common.Address
	CpAccount     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_Collateral *CollateralFilterer) FilterDeposit(opts *bind.FilterOpts, fundingWallet []common.Address, cpAccount []common.Address) (*CollateralDepositIterator, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &CollateralDepositIterator{contract: _Collateral.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_Collateral *CollateralFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CollateralDeposit, fundingWallet []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var fundingWalletRule []interface{}
	for _, fundingWalletItem := range fundingWallet {
		fundingWalletRule = append(fundingWalletRule, fundingWalletItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Deposit", fundingWalletRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralDeposit)
				if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed fundingWallet, address indexed cpAccount, uint256 depositAmount)
func (_Collateral *CollateralFilterer) ParseDeposit(log types.Log) (*CollateralDeposit, error) {
	event := new(CollateralDeposit)
	if err := _Collateral.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Collateral contract.
type CollateralOwnershipTransferredIterator struct {
	Event *CollateralOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralOwnershipTransferred represents a OwnershipTransferred event raised by the Collateral contract.
type CollateralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collateral *CollateralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollateralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollateralOwnershipTransferredIterator{contract: _Collateral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collateral *CollateralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollateralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralOwnershipTransferred)
				if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Collateral *CollateralFilterer) ParseOwnershipTransferred(log types.Log) (*CollateralOwnershipTransferred, error) {
	event := new(CollateralOwnershipTransferred)
	if err := _Collateral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTaskCreatedIterator is returned from FilterTaskCreated and is used to iterate over the raw logs and unpacked data for TaskCreated events raised by the Collateral contract.
type CollateralTaskCreatedIterator struct {
	Event *CollateralTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTaskCreated represents a TaskCreated event raised by the Collateral contract.
type CollateralTaskCreated struct {
	TaskContractAddress common.Address
	CpAccountAddress    common.Address
	Collateral          *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTaskCreated is a free log retrieval operation binding the contract event 0x3f9b921f6bb5d577cdf49122202d27f7c50e3cc981c47543dc4a59c00dee5229.
//
// Solidity: event TaskCreated(address indexed taskContractAddress, address cpAccountAddress, uint256 collateral)
func (_Collateral *CollateralFilterer) FilterTaskCreated(opts *bind.FilterOpts, taskContractAddress []common.Address) (*CollateralTaskCreatedIterator, error) {

	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "TaskCreated", taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTaskCreatedIterator{contract: _Collateral.contract, event: "TaskCreated", logs: logs, sub: sub}, nil
}

// WatchTaskCreated is a free log subscription operation binding the contract event 0x3f9b921f6bb5d577cdf49122202d27f7c50e3cc981c47543dc4a59c00dee5229.
//
// Solidity: event TaskCreated(address indexed taskContractAddress, address cpAccountAddress, uint256 collateral)
func (_Collateral *CollateralFilterer) WatchTaskCreated(opts *bind.WatchOpts, sink chan<- *CollateralTaskCreated, taskContractAddress []common.Address) (event.Subscription, error) {

	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "TaskCreated", taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTaskCreated)
				if err := _Collateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCreated is a log parse operation binding the contract event 0x3f9b921f6bb5d577cdf49122202d27f7c50e3cc981c47543dc4a59c00dee5229.
//
// Solidity: event TaskCreated(address indexed taskContractAddress, address cpAccountAddress, uint256 collateral)
func (_Collateral *CollateralFilterer) ParseTaskCreated(log types.Log) (*CollateralTaskCreated, error) {
	event := new(CollateralTaskCreated)
	if err := _Collateral.contract.UnpackLog(event, "TaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralTaskStatusChangedIterator is returned from FilterTaskStatusChanged and is used to iterate over the raw logs and unpacked data for TaskStatusChanged events raised by the Collateral contract.
type CollateralTaskStatusChangedIterator struct {
	Event *CollateralTaskStatusChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralTaskStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralTaskStatusChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralTaskStatusChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralTaskStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralTaskStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralTaskStatusChanged represents a TaskStatusChanged event raised by the Collateral contract.
type CollateralTaskStatusChanged struct {
	TaskContractAddress common.Address
	NewStatus           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTaskStatusChanged is a free log retrieval operation binding the contract event 0x43a64df649a51c960307a3b5a4ec25a67c76629f0881669a214d162ae4ad2b5f.
//
// Solidity: event TaskStatusChanged(address indexed taskContractAddress, uint256 newStatus)
func (_Collateral *CollateralFilterer) FilterTaskStatusChanged(opts *bind.FilterOpts, taskContractAddress []common.Address) (*CollateralTaskStatusChangedIterator, error) {

	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "TaskStatusChanged", taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &CollateralTaskStatusChangedIterator{contract: _Collateral.contract, event: "TaskStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTaskStatusChanged is a free log subscription operation binding the contract event 0x43a64df649a51c960307a3b5a4ec25a67c76629f0881669a214d162ae4ad2b5f.
//
// Solidity: event TaskStatusChanged(address indexed taskContractAddress, uint256 newStatus)
func (_Collateral *CollateralFilterer) WatchTaskStatusChanged(opts *bind.WatchOpts, sink chan<- *CollateralTaskStatusChanged, taskContractAddress []common.Address) (event.Subscription, error) {

	var taskContractAddressRule []interface{}
	for _, taskContractAddressItem := range taskContractAddress {
		taskContractAddressRule = append(taskContractAddressRule, taskContractAddressItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "TaskStatusChanged", taskContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralTaskStatusChanged)
				if err := _Collateral.contract.UnpackLog(event, "TaskStatusChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskStatusChanged is a log parse operation binding the contract event 0x43a64df649a51c960307a3b5a4ec25a67c76629f0881669a214d162ae4ad2b5f.
//
// Solidity: event TaskStatusChanged(address indexed taskContractAddress, uint256 newStatus)
func (_Collateral *CollateralFilterer) ParseTaskStatusChanged(log types.Log) (*CollateralTaskStatusChanged, error) {
	event := new(CollateralTaskStatusChanged)
	if err := _Collateral.contract.UnpackLog(event, "TaskStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollateralWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Collateral contract.
type CollateralWithdrawIterator struct {
	Event *CollateralWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CollateralWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollateralWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CollateralWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CollateralWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollateralWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollateralWithdraw represents a Withdraw event raised by the Collateral contract.
type CollateralWithdraw struct {
	CpOwner        common.Address
	CpAccount      common.Address
	WithdrawAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) FilterWithdraw(opts *bind.FilterOpts, cpOwner []common.Address, cpAccount []common.Address) (*CollateralWithdrawIterator, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _Collateral.contract.FilterLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return &CollateralWithdrawIterator{contract: _Collateral.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CollateralWithdraw, cpOwner []common.Address, cpAccount []common.Address) (event.Subscription, error) {

	var cpOwnerRule []interface{}
	for _, cpOwnerItem := range cpOwner {
		cpOwnerRule = append(cpOwnerRule, cpOwnerItem)
	}
	var cpAccountRule []interface{}
	for _, cpAccountItem := range cpAccount {
		cpAccountRule = append(cpAccountRule, cpAccountItem)
	}

	logs, sub, err := _Collateral.contract.WatchLogs(opts, "Withdraw", cpOwnerRule, cpAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollateralWithdraw)
				if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed cpOwner, address indexed cpAccount, uint256 withdrawAmount)
func (_Collateral *CollateralFilterer) ParseWithdraw(log types.Log) (*CollateralWithdraw, error) {
	event := new(CollateralWithdraw)
	if err := _Collateral.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
