package account

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestCpStub_GetCpAccountInfo(t *testing.T) {
	client, err := ethclient.Dial("https://saturn-rpc.swanchain.io")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer client.Close()

	stub := &CpStub{}
	cpAccountAddress := common.HexToAddress("0x6182ECDA7f01E80224E591323C16d468adeB07Db")
	taskClient, err := NewAccount(cpAccountAddress, client)
	if err != nil {
		log.Fatalln(err)
		return
	}

	stub.account = taskClient
	stub.client = client
	stub.ContractAddress = "0x6182ECDA7f01E80224E591323C16d468adeB07Db"

	a, err := stub.GetCpAccountInfo()
	fmt.Printf("a: %+v \n", a)

	accountStub, err := NewAccountStub(client, WithContractAddress("0x6182ECDA7f01E80224E591323C16d468adeB07Db"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	b, err := accountStub.GetCpAccountInfo()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf(" b: %+v", b)

}
