package blockchain

import (
	"eth-api/config"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

var EthClient *ethclient.Client
var RpcClient *rpc.Client

func InitEthC() {

	var err error
	fmt.Println(config.GetInfuraUrl())
	EthClient, err = ethclient.Dial(config.GetInfuraUrl())
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum EthClient: %v", err)
	}

	RpcClient, err = rpc.Dial(config.GetInfuraUrl())
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum RpcClient: %v", err)
	}

}
