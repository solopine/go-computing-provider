package main

import (
	"context"
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
	router.GET("/cp", computing.StatisticalSources)

	router.GET("/host/info", computing.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", computing.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", computing.RedeployJob)
	router.DELETE("/lagrange/jobs", computing.CancelJob)
	router.POST("/lagrange/jobs/renew", computing.ReNewJob)
	router.GET("/lagrange/spaces/log", computing.GetSpaceLog)
	router.POST("/lagrange/cp/proof", computing.DoProof)
	router.GET("/lagrange/cp/whiteList", computing.WhiteList)
	router.GET("/lagrange/job/:job_uuid", computing.GetJobStatus)

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

		var fcpCollateralBalance, ecpCollateralBalance, ownerBalance string
		var contractAddress, ownerAddress, workerAddress, beneficiaryAddress, taskTypes, chainNodeId, version string

		cpStub, err := account.NewAccountStub(client)
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount failed, error: %v", err)
			}

			for _, taskType := range cpAccount.TaskTypes {
				switch taskType {
				case 1:
					taskTypes += "Fil-C2,"
				case 2:
					taskTypes += "Aleo,"
				case 3:
					taskTypes += "Ai,"
				}
			}
			if taskTypes != "" {
				taskTypes = taskTypes[:len(taskTypes)-1]
			}

			contractAddress = cpStub.ContractAddress
			ownerAddress = cpAccount.OwnerAddress
			workerAddress = cpAccount.WorkerAddress
			beneficiaryAddress = cpAccount.Beneficiary
			chainNodeId = cpAccount.NodeId
			version = cpAccount.Version
		}

		ownerBalance, err = wallet.Balance(context.TODO(), client, ownerAddress)
		fcpCollateralStub, err := collateral.NewCollateralStub(client, collateral.WithPublicKey(ownerAddress))
		if err == nil {
			fcpCollateralBalance, err = fcpCollateralStub.Balances()
		}

		ecpCollateral, err := account.NewCollateralStub(client, account.WithPublicKey(ownerAddress))
		if err == nil {
			cpCollateralInfo, err := ecpCollateral.CpInfo()
			if err == nil {
				ecpCollateralBalance = cpCollateralInfo.CollateralBalance
			}
		}

		var domain = conf.GetConfig().API.Domain
		if strings.HasPrefix(domain, ".") {
			domain = domain[1:]
		}
		var taskData [][]string

		taskData = append(taskData, []string{"Owner:", ownerAddress})
		taskData = append(taskData, []string{"Contract Address:", contractAddress})
		taskData = append(taskData, []string{"Contract Version:", version})
		taskData = append(taskData, []string{"Multi-Address:", conf.GetConfig().API.MultiAddress})
		taskData = append(taskData, []string{"Node ID:", localNodeId})
		taskData = append(taskData, []string{"Task Types:", taskTypes})
		taskData = append(taskData, []string{"ECP:"})
		taskData = append(taskData, []string{"   Worker Address:", workerAddress})
		taskData = append(taskData, []string{"   Beneficiary Address:", beneficiaryAddress})
		taskData = append(taskData, []string{"   Available(SWAN-ETH):", ownerBalance})
		taskData = append(taskData, []string{"   Collateral(SWAN-ETH):", ecpCollateralBalance})
		taskData = append(taskData, []string{"FCP:"})
		taskData = append(taskData, []string{"   Domain:", domain})
		taskData = append(taskData, []string{"   Running deployments:", strconv.Itoa(count)})
		taskData = append(taskData, []string{"   Available(SWAN-ETH):", ownerBalance})
		taskData = append(taskData, []string{"   Collateral(SWAN-ETH):", fcpCollateralBalance})

		var rowColorList []RowColor
		if taskTypes != "" {
			var rowColor []tablewriter.Colors
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			rowColorList = append(rowColorList, RowColor{
				row:    5,
				column: []int{1},
				color:  rowColor,
			})
		}
		header := []string{"Name:", conf.GetConfig().API.NodeName}
		NewVisualTable(header, taskData, rowColorList).Generate(false)
		if localNodeId != chainNodeId {
			fmt.Printf("NodeId mismatch, local node id: %s, chain node id: %s.\n", localNodeId, chainNodeId)
		}
		return nil
	},
}

var stateInfoCmd = &cli.Command{
	Name:  "state-info",
	Usage: "Print computing-provider chain info",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "chain",
			Usage: "specify the account to send funds from",
			Value: conf.DefaultRpc,
		},
		&cli.StringFlag{
			Name:  "contract",
			Usage: "specify the account contract",
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

		chain := cctx.String("chain")
		if strings.TrimSpace(chain) == "" {
			return fmt.Errorf("chain field cannot be empty")
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

		var ownerAddress, contractAddress, beneficiaryAddress, chainNodeId,
			multiAddress, workerAddress, taskTypes, version string

		cpStub, err := account.NewAccountStub(client, account.WithContractAddress(cctx.String("contract")))
		if err == nil {
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				err = fmt.Errorf("get cpAccount failed, error: %v", err)
			}

			contractAddress = cpStub.ContractAddress
			chainNodeId = cpAccount.NodeId
			multiAddress = strings.Join(cpAccount.MultiAddresses, ",")
			ownerAddress = cpAccount.OwnerAddress
			beneficiaryAddress = cpAccount.Beneficiary
			workerAddress = cpAccount.WorkerAddress
			version = cpAccount.Version

			for _, taskType := range cpAccount.TaskTypes {
				switch taskType {
				case 1:
					taskTypes += "Fil-C2,"

				case 2:
					taskTypes += "Aleo,"
				}
			}
			if taskTypes != "" {
				taskTypes = taskTypes[:len(taskTypes)-1]
			}
		}

		var taskData [][]string
		taskData = append(taskData, []string{"Node ID:", chainNodeId})
		taskData = append(taskData, []string{"Multi-Address:", multiAddress})
		taskData = append(taskData, []string{"Contract Address:", contractAddress})
		taskData = append(taskData, []string{"Contract Version:", version})
		taskData = append(taskData, []string{"Task Types:", taskTypes})
		taskData = append(taskData, []string{"Worker Address:", workerAddress})
		taskData = append(taskData, []string{"Beneficiary Address:", beneficiaryAddress})

		var rowColorList []RowColor
		if taskTypes != "" {
			var rowColor []tablewriter.Colors
			rowColor = []tablewriter.Colors{{tablewriter.Bold, tablewriter.FgGreenColor}}
			rowColorList = append(rowColorList, RowColor{
				row:    4,
				column: []int{1},
				color:  rowColor,
			})
		}
		header := []string{"Owner:", ownerAddress}
		NewVisualTable(header, taskData, rowColorList).Generate(false)
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
			Usage: "The cp listens on port",
			Value: 9085,
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
		changeWorkerAddressCmd,
		changeBeneficiaryAddressCmd,
		changeTaskTypesCmd,
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
			Name:  "workerAddress",
			Usage: "Specify a workerAddress",
		},
		&cli.StringFlag{
			Name:  "beneficiaryAddress",
			Usage: "Specify a beneficiaryAddress to receive rewards. If not specified, use the same address as ownerAddress",
		},
		&cli.StringFlag{
			Name:  "task-types",
			Usage: "Task types of CP (1:Fil-C2, 2:Aleo, 3:Ai), separated by commas",
		},
	},
	Action: func(cctx *cli.Context) error {
		ownerAddress := cctx.String("ownerAddress")
		if strings.TrimSpace(ownerAddress) == "" {
			return fmt.Errorf("ownerAddress is not empty")
		}

		workerAddress := cctx.String("workerAddress")
		if strings.TrimSpace(workerAddress) == "" {
			return fmt.Errorf("workerAddress is not empty")
		}

		beneficiaryAddress := cctx.String("beneficiaryAddress")
		if strings.TrimSpace(beneficiaryAddress) == "" {
			beneficiaryAddress = ownerAddress
		}

		taskTypes := strings.TrimSpace(cctx.String("task-types"))
		if strings.TrimSpace(taskTypes) == "" {
			return fmt.Errorf("taskTypes is not empty")
		}

		var taskTypesUint []uint8
		if strings.Index(taskTypes, ",") > 0 {
			for _, taskT := range strings.Split(taskTypes, ",") {
				tt, _ := strconv.ParseUint(taskT, 10, 64)
				if tt != 1 && tt != 2 && tt != 3 {
					return fmt.Errorf("TaskTypes supports 1, 2, 3")
				}
				taskTypesUint = append(taskTypesUint, uint8(tt))
			}
		} else {
			tt, _ := strconv.ParseUint(taskTypes, 10, 64)
			if tt != 1 && tt != 2 && tt != 3 {
				return fmt.Errorf("TaskTypes supports 1, 2, 3")
			}
		}

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, true); err != nil {
			logs.GetLogger().Fatal(err)
		}
		return createAccount(cpRepoPath, ownerAddress, beneficiaryAddress, workerAddress, taskTypesUint)
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
			return fmt.Errorf("ownerAddress is not empty")
		}

		if cctx.NArg() != 1 {
			return fmt.Errorf(" Requires a ownerAddress")
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
		changeBeneficiaryAddressTx, err := cpStub.ChangeBeneficiary(common.HexToAddress(beneficiaryAddress))
		if err != nil {
			logs.GetLogger().Errorf("change beneficiary address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("changeBeneficiaryAddress Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeWorkerAddressCmd = &cli.Command{
	Name:      "changeWorkerAddress",
	Usage:     "Update workerAddress of CP",
	ArgsUsage: "[workerAddress]",
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

		workerAddress := cctx.Args().Get(0)
		if strings.TrimSpace(workerAddress) == "" {
			return fmt.Errorf("failed to parse target address: %s", workerAddress)
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
		changeBeneficiaryAddressTx, err := cpStub.ChangeWorkerAddress(common.HexToAddress(workerAddress))
		if err != nil {
			logs.GetLogger().Errorf("change worker address tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("changeWorkerAddress Transaction hash: %s \n", changeBeneficiaryAddressTx)
		return nil
	},
}

var changeTaskTypesCmd = &cli.Command{
	Name:      "changeTaskTypes",
	Usage:     "Update taskTypes of CP (1:Fil-C2, 2:Aleo, 3: Ai), separated by commas",
	ArgsUsage: "[TaskTypes]",
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

		taskTypes := strings.TrimSpace(cctx.Args().Get(0))
		if strings.TrimSpace(taskTypes) == "" {
			return fmt.Errorf("taskTypes is not empty")
		}

		var taskTypesUint []uint8
		if strings.Index(taskTypes, ",") > 0 {
			for _, taskT := range strings.Split(taskTypes, ",") {
				tt, _ := strconv.ParseUint(taskT, 10, 64)
				if tt != 1 && tt != 2 && tt != 3 {
					return fmt.Errorf("TaskTypes supports 1, 2, 3")
				}
				taskTypesUint = append(taskTypesUint, uint8(tt))
			}
		} else {
			tt, _ := strconv.ParseUint(taskTypes, 10, 64)
			if tt != 1 && tt != 2 && tt != 3 {
				return fmt.Errorf("TaskTypes supports 1, 2, 3")
			}
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

		changeTaskTypesTx, err := cpStub.ChangeTaskTypes(taskTypesUint)
		if err != nil {
			logs.GetLogger().Errorf("change TaskTypes tx failed, error: %v,", err)
			return err
		}
		fmt.Printf("ChangeTaskTypes Transaction hash: %s \n", changeTaskTypesTx)
		return nil
	},
}
