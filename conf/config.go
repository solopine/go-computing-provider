package conf

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var config *ComputeNode

const (
	DefaultRpc = "swan"
)

// ComputeNode is a compute node config
type ComputeNode struct {
	API      API
	UBI      UBI
	LOG      LOG
	HUB      HUB
	MCS      MCS
	Registry Registry
	RPC      RPC
	CONTRACT CONTRACT
}

type API struct {
	Port            int
	MultiAddress    string
	Domain          string
	NodeName        string
	WalletWhiteList string
}
type UBI struct {
	UbiTask     bool
	UbiEnginePk string
	UbiUrl      string
}

type LOG struct {
	CrtFile string
	KeyFile string
}

type HUB struct {
	ServerUrl        string
	AccessToken      string
	BalanceThreshold float64
	OrchestratorPk   string
	VerifySign       bool
}

type MCS struct {
	ApiKey        string
	AccessToken   string
	BucketName    string
	Network       string
	FileCachePath string
}

type Registry struct {
	ServerAddress string
	UserName      string
	Password      string
}

type RPC struct {
	SwanTestnet string `toml:"SWAN_TESTNET"`
	SwanMainnet string `toml:"SWAN_MAINNET"`
}

type CONTRACT struct {
	SwanToken    string `toml:"SWAN_CONTRACT"`
	Collateral   string `toml:"SWAN_COLLATERAL_CONTRACT"`
	Register     string `toml:"REGISTER_CP_CONTRACT"`
	ZkCollateral string `toml:"ZK_COLLATERAL_CONTRACT"`
}

func GetRpcByName(rpcName string) (string, error) {
	var rpc string
	switch rpcName {
	case DefaultRpc:
		rpc = GetConfig().RPC.SwanTestnet
		break
	}
	return rpc, nil
}

func InitConfig(cpRepoPath string, standalone bool) error {
	configFile := filepath.Join(cpRepoPath, "config.toml")
	metaData, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		return fmt.Errorf("failed load config file, path: %s, error: %w", configFile, err)
	}
	if standalone {
		if !requiredFieldsAreGivenForSeparate(metaData) {
			log.Fatal("Required fields not given")
		}
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("Required fields not given")
		}
	}
	return nil
}

func GetConfig() *ComputeNode {
	return config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"LOG"},
		{"UBI"},
		{"HUB"},
		{"MCS"},
		{"Registry"},
		{"RPC"},
		{"CONTRACT"},

		{"API", "MultiAddress"},
		{"API", "Domain"},
		{"API", "NodeName"},

		{"LOG", "CrtFile"},
		{"LOG", "KeyFile"},

		{"UBI", "UbiTask"},
		{"UBI", "UbiEnginePk"},
		{"UBI", "UbiUrl"},

		{"HUB", "ServerUrl"},
		{"HUB", "AccessToken"},

		{"MCS", "ApiKey"},
		{"MCS", "BucketName"},
		{"MCS", "Network"},
		{"MCS", "FileCachePath"},

		{"RPC", "SWAN_TESTNET"},

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

func requiredFieldsAreGivenForSeparate(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"API"},
		{"UBI"},
		{"HUB"},

		{"API", "MultiAddress"},
		{"API", "NodeName"},

		{"UBI", "UbiTask"},
		{"UBI", "UbiEnginePk"},
		{"UBI", "UbiUrl"},

		{"RPC", "SWAN_TESTNET"},

		{"CONTRACT", "SWAN_CONTRACT"},
		{"CONTRACT", "SWAN_COLLATERAL_CONTRACT"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("Required fields ", v)
		}
	}

	return true
}

//go:embed config.toml
var configFileContent string

func GenerateRepo(cpRepoPath string) error {
	var configTmpl ComputeNode
	var configFile *os.File
	var err error

	configFilePath := path.Join(cpRepoPath, "config.toml")
	if _, err = os.Stat(configFilePath); os.IsNotExist(err) {
		if _, err = toml.Decode(configFileContent, &configTmpl); err != nil {
			return fmt.Errorf("parse toml data failed, error: %v", err)
		}
		configFile, err = os.Create(configFilePath)
		if err != nil {
			return fmt.Errorf("create config.toml file failed, error: %v", err)
		}
		if err = toml.NewEncoder(configFile).Encode(configTmpl); err != nil {
			return fmt.Errorf("write data to config.toml file failed, error: %v", err)
		}
	}
	return nil
}

func UpdateConfigFile(cpRepoPath string, multiAddress, nodeName string, port int) error {
	var configTmpl ComputeNode
	var configFile *os.File
	var err error

	configFilePath := path.Join(cpRepoPath, "config.toml")
	if _, err = toml.DecodeFile(configFilePath, &configTmpl); err != nil {
		return err
	}
	os.Remove(configFilePath)

	configFile, err = os.Create(configFilePath)
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(multiAddress)) != 0 && !strings.EqualFold(strings.TrimSpace(multiAddress), strings.TrimSpace(configTmpl.API.MultiAddress)) {
		configTmpl.API.MultiAddress = multiAddress
	}

	if len(strings.TrimSpace(nodeName)) != 0 {
		configTmpl.API.NodeName = nodeName
	} else {
		hostname, err := os.Hostname()
		if err != nil {
			return fmt.Errorf("get hostname failed, error: %v", err)
		}
		configTmpl.API.NodeName = hostname
	}

	if port != 0 {
		configTmpl.API.Port = port
	}

	if err = toml.NewEncoder(configFile).Encode(configTmpl); err != nil {
		return err
	}
	return nil
}
