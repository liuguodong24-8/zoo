package data

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

func (cli *EthGrpcClient) GetPrivate(c context.Context, privateKey string) (*bind.TransactOpts, error) {
	priKeyECD, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	chainId, err := cli.ChainID(c)
	if err != nil {
		return nil, err
	}

	param, err := bind.NewKeyedTransactorWithChainID(priKeyECD, chainId)
	if err != nil {
		return nil, err
	}

	return param, nil
}
