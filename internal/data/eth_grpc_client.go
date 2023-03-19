package data

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthGrpcClient struct {
	*ethclient.Client
}

func NewEthClient(endPoint string) (*EthGrpcClient, func(), error) {
	conn, err := ethclient.Dial(endPoint)
	if err != nil {
		return nil, func() {}, err
	}
	return &EthGrpcClient{conn}, func() {
		conn.Close()
	}, nil
}
