// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package util

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
)

// ZooNftMetaData contains all meta data concerning the ZooNft contract.
var ZooNftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NFT_FACTORY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI_\",\"type\":\"string\"}],\"name\":\"_setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"}],\"name\":\"getNftURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseURIInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"levels\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"categorys\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"URIs\",\"type\":\"string[]\"}],\"name\":\"setMultiNftURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftFactory\",\"type\":\"address\"}],\"name\":\"setNFTFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"name\":\"setNftURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052604051620023e2380380620023e2833981016040819052620000269162000381565b604080518082018252600680825265169bdbd3919560d21b6020808401828152855180870190965292855284015281519192916200006791600291620002db565b5080516200007d906003906020840190620002db565b5050600054610100900460ff16159050808015620000a25750600054600160ff909116105b80620000d25750620000bf30620001e760201b62000cdb1760201c565b158015620000d2575060005460ff166001145b6200013a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff1916600117905580156200015e576000805461ff0019166101001790555b6200016b600083620001f6565b620001987f5eef2d0cdce42042b2a1881ea7f51ddbc3df94d471ed2a17f74c02f01f8301e3600062000206565b8015620001df576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050620003ee565b6001600160a01b03163b151590565b62000202828262000253565b5050565b6000828152600160208190526040808320909101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b60008281526001602090815260408083206001600160a01b038516845290915290205460ff16620002025760008281526001602081815260408084206001600160a01b0386168086529252808420805460ff19169093179092559051339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b828054620002e990620003b1565b90600052602060002090601f0160209004810192826200030d576000855562000358565b82601f106200032857805160ff191683800117855562000358565b8280016001018555821562000358579182015b82811115620003585782518255916020019190600101906200033b565b50620003669291506200036a565b5090565b5b808211156200036657600081556001016200036b565b60006020828403121562000393578081fd5b81516001600160a01b0381168114620003aa578182fd5b9392505050565b600281046001821680620003c657607f821691505b60208210811415620003e857634e487b7160e01b600052602260045260246000fd5b50919050565b611fe480620003fe6000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806355f804b31161010f578063a217fddf116100a2578063cb2d93ce11610071578063cb2d93ce1461042a578063d547741f1461043d578063e60b943a14610450578063e985e9c514610463576101e5565b8063a217fddf146103e9578063a22cb465146103f1578063b88d4fde14610404578063c87b56dd14610417576101e5565b80637a45fe83116100de5780637a45fe83146103945780637a47dc77146103bb57806391d14854146103ce57806395d89b41146103e1576101e5565b806355f804b314610348578063574bd1081461035b5780636352211e1461036e57806370a0823114610381576101e5565b8063248a9ca31161018757806340c10f191161015657806340c10f19146102fc57806342842e0e1461030f57806342966c6814610322578063509f188014610335576101e5565b8063248a9ca31461029f5780632f2ff15d146102c357806336568abe146102d65780633e42c199146102e9576101e5565b8063081812fc116101c3578063081812fc1461023c578063095ea7b31461026757806318160ddd1461027a57806323b872dd1461028c576101e5565b806301538868146101ea57806301ffc9a7146101ff57806306fdde0314610227575b600080fd5b6101fd6101f8366004611bbe565b61049f565b005b61021261020d366004611b53565b6104c3565b60405190151581526020015b60405180910390f35b61022f6104fc565b60405161021e9190611d50565b61024f61024a366004611b19565b61058e565b6040516001600160a01b03909116815260200161021e565b6101fd610275366004611a1d565b6105b5565b600a545b60405190815260200161021e565b6101fd61029a36600461192f565b6106cb565b61027e6102ad366004611b19565b6000908152600160208190526040909120015490565b6101fd6102d1366004611b31565b6106fc565b6101fd6102e4366004611b31565b610722565b6101fd6102f7366004611b8b565b6107a0565b6101fd61030a366004611a1d565b6107b3565b6101fd61031d36600461192f565b6107cd565b6101fd610330366004611b19565b6107e8565b61022f610343366004611c03565b6107f4565b6101fd610356366004611b8b565b610899565b6101fd610369366004611a46565b6108b6565b61024f61037c366004611b19565b61099c565b61027e61038f3660046118e3565b6109fc565b61027e7f5eef2d0cdce42042b2a1881ea7f51ddbc3df94d471ed2a17f74c02f01f8301e381565b6101fd6103c9366004611c24565b610a82565b6102126103dc366004611b31565b610abd565b61022f610ae8565b61027e600081565b6101fd6103ff3660046119e3565b610af7565b6101fd61041236600461196a565b610b02565b61022f610425366004611b19565b610b34565b6101fd6104383660046118e3565b610bdd565b6101fd61044b366004611b31565b610c07565b61022f61045e366004611c03565b610c2d565b6102126104713660046118fd565b6001600160a01b03918216600090815260096020908152604080832093909416825291909152205460ff1690565b600082815260086020908152604090912082516104be92840190611750565b505050565b60006001600160e01b03198216637965db0b60e01b14806104f457506301ffc9a760e01b6001600160e01b03198316145b90505b919050565b60606002805461050b90611f16565b80601f016020809104026020016040519081016040528092919081815260200182805461053790611f16565b80156105845780601f1061055957610100808354040283529160200191610584565b820191906000526020600020905b81548152906001019060200180831161056757829003601f168201915b5050505050905090565b600061059982610cea565b506000908152600760205260409020546001600160a01b031690565b60006105c08261099c565b9050806001600160a01b0316836001600160a01b031614156106335760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b038216148061064f575061064f8133610471565b6106c15760405162461bcd60e51b815260206004820152603d60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000606482015260840161062a565b6104be8383610d49565b6106d53382610db7565b6106f15760405162461bcd60e51b815260040161062a90611d63565b6104be838383610e36565b6000828152600160208190526040909120015461071881610f9a565b6104be8383610fa4565b6001600160a01b03811633146107925760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b606482015260840161062a565b61079c828261100f565b5050565b805161079c906004906020840190611750565b61079c828260405180602001604052806000815250611076565b6104be83838360405180602001604052806000815250610b02565b6107f1816110a9565b50565b600b6020908152600092835260408084209091529082529020805461081890611f16565b80601f016020809104026020016040519081016040528092919081815260200182805461084490611f16565b80156108915780601f1061086657610100808354040283529160200191610891565b820191906000526020600020905b81548152906001019060200180831161087457829003601f168201915b505050505081565b6108a4600033610abd565b6108ad57600080fd5b6107f1816107a0565b6108c1600033610abd565b6108ca57600080fd5b60005b8351811015610996578181815181106108f657634e487b7160e01b600052603260045260246000fd5b6020026020010151600b600086848151811061092257634e487b7160e01b600052603260045260246000fd5b60200260200101518152602001908152602001600020600085848151811061095a57634e487b7160e01b600052603260045260246000fd5b602002602001015181526020019081526020016000209080519060200190610983929190611750565b508061098e81611f51565b9150506108cd565b50505050565b6000818152600560205260408120546001600160a01b0316806104f45760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161062a565b60006001600160a01b038216610a665760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b606482015260840161062a565b506001600160a01b031660009081526006602052604090205490565b610a8d600033610abd565b610a9657600080fd5b6000838152600b602090815260408083208584528252909120825161099692840190611750565b60009182526001602090815260408084206001600160a01b0393909316845291905290205460ff1690565b60606003805461050b90611f16565b61079c33838361113f565b610b0c3383610db7565b610b285760405162461bcd60e51b815260040161062a90611d63565b6109968484848461120e565b6060610b3f82610cea565b60008281526008602052604090208054610b5890611f16565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8490611f16565b8015610bd15780601f10610ba657610100808354040283529160200191610bd1565b820191906000526020600020905b815481529060010190602001808311610bb457829003601f168201915b50505050509050919050565b6107f17f5eef2d0cdce42042b2a1881ea7f51ddbc3df94d471ed2a17f74c02f01f8301e3826106fc565b60008281526001602081905260409091200154610c2381610f9a565b6104be838361100f565b6000828152600b602090815260408083208484529091529020805460609190610c5590611f16565b80601f0160208091040260200160405190810160405280929190818152602001828054610c8190611f16565b8015610cce5780601f10610ca357610100808354040283529160200191610cce565b820191906000526020600020905b815481529060010190602001808311610cb157829003601f168201915b5050505050905092915050565b6001600160a01b03163b151590565b6000818152600560205260409020546001600160a01b03166107f15760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b604482015260640161062a565b600081815260076020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610d7e8261099c565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b600080610dc38361099c565b9050806001600160a01b0316846001600160a01b03161480610e0a57506001600160a01b0380821660009081526009602090815260408083209388168352929052205460ff165b80610e2e5750836001600160a01b0316610e238461058e565b6001600160a01b0316145b949350505050565b826001600160a01b0316610e498261099c565b6001600160a01b031614610e6f5760405162461bcd60e51b815260040161062a90611e02565b6001600160a01b038216610ed15760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b606482015260840161062a565b826001600160a01b0316610ee48261099c565b6001600160a01b031614610f0a5760405162461bcd60e51b815260040161062a90611e02565b600081815260076020908152604080832080546001600160a01b03199081169091556001600160a01b0387811680865260068552838620805460001901905590871680865283862080546001019055868652600590945282852080549092168417909155905184937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a46104be565b6107f18133611241565b610fae8282610abd565b61079c5760008281526001602081815260408084206001600160a01b0386168086529252808420805460ff19169093179092559051339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6110198282610abd565b1561079c5760008281526001602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b611080838361129a565b61108d6000848484611448565b6104be5760405162461bcd60e51b815260040161062a90611db0565b60006110b48261099c565b90506110bf8261099c565b600083815260076020908152604080832080546001600160a01b03199081169091556001600160a01b0385168085526006845282852080546000190190558785526005909352818420805490911690555192935084927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a461079c565b816001600160a01b0316836001600160a01b031614156111a15760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015260640161062a565b6001600160a01b03838116600081815260096020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b611219848484610e36565b61122584848484611448565b6109965760405162461bcd60e51b815260040161062a90611db0565b61124b8282610abd565b61079c5761125881611555565b611263836020611567565b604051602001611274929190611c9e565b60408051601f198184030181529082905262461bcd60e51b825261062a91600401611d50565b6001600160a01b0382166112f05760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015260640161062a565b6000818152600560205260409020546001600160a01b0316156113555760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161062a565b6000818152600560205260409020546001600160a01b0316156113ba5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015260640161062a565b6001600160a01b0382166000818152600660209081526040808320805460010190558483526005909152812080546001600160a01b031916909217909155600a80549161140683611f51565b909155505060405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a461079c565b60006001600160a01b0384163b1561154a57604051630a85bd0160e11b81526001600160a01b0385169063150b7a029061148c903390899088908890600401611d13565b602060405180830381600087803b1580156114a657600080fd5b505af19250505080156114d6575060408051601f3d908101601f191682019092526114d391810190611b6f565b60015b611530573d808015611504576040519150601f19603f3d011682016040523d82523d6000602084013e611509565b606091505b5080516115285760405162461bcd60e51b815260040161062a90611db0565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610e2e565b506001949350505050565b60606104f46001600160a01b03831660145b60606000611576836002611eb4565b611581906002611e9c565b67ffffffffffffffff8111156115a757634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156115d1576020820181803683370190505b509050600360fc1b816000815181106115fa57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061163757634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350600061165b846002611eb4565b611666906001611e9c565b90505b60018111156116fa576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106116a857634e487b7160e01b600052603260045260246000fd5b1a60f81b8282815181106116cc57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535060049490941c936116f381611eff565b9050611669565b5083156117495760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161062a565b9392505050565b82805461175c90611f16565b90600052602060002090601f01602090048101928261177e57600085556117c4565b82601f1061179757805160ff19168380011785556117c4565b828001600101855582156117c4579182015b828111156117c45782518255916020019190600101906117a9565b506117d09291506117d4565b5090565b5b808211156117d057600081556001016117d5565b600067ffffffffffffffff83111561180357611803611f82565b611816601f8401601f1916602001611e47565b905082815283838301111561182a57600080fd5b828260208301376000602084830101529392505050565b80356001600160a01b03811681146104f757600080fd5b600082601f830112611868578081fd5b8135602061187d61187883611e78565b611e47565b8281528181019085830183850287018401881015611899578586fd5b855b858110156118b75781358452928401929084019060010161189b565b5090979650505050505050565b600082601f8301126118d4578081fd5b611749838335602085016117e9565b6000602082840312156118f4578081fd5b61174982611841565b6000806040838503121561190f578081fd5b61191883611841565b915061192660208401611841565b90509250929050565b600080600060608486031215611943578081fd5b61194c84611841565b925061195a60208501611841565b9150604084013590509250925092565b6000806000806080858703121561197f578081fd5b61198885611841565b935061199660208601611841565b925060408501359150606085013567ffffffffffffffff8111156119b8578182fd5b8501601f810187136119c8578182fd5b6119d7878235602084016117e9565b91505092959194509250565b600080604083850312156119f5578182fd5b6119fe83611841565b915060208301358015158114611a12578182fd5b809150509250929050565b60008060408385031215611a2f578182fd5b611a3883611841565b946020939093013593505050565b600080600060608486031215611a5a578283fd5b833567ffffffffffffffff80821115611a71578485fd5b611a7d87838801611858565b9450602091508186013581811115611a93578485fd5b611a9f88828901611858565b945050604086013581811115611ab3578384fd5b86019050601f81018713611ac5578283fd5b8035611ad361187882611e78565b81815283810190838501865b84811015611b0857611af68c8884358901016118c4565b84529286019290860190600101611adf565b505080955050505050509250925092565b600060208284031215611b2a578081fd5b5035919050565b60008060408385031215611b43578182fd5b8235915061192660208401611841565b600060208284031215611b64578081fd5b813561174981611f98565b600060208284031215611b80578081fd5b815161174981611f98565b600060208284031215611b9c578081fd5b813567ffffffffffffffff811115611bb2578182fd5b610e2e848285016118c4565b60008060408385031215611bd0578182fd5b82359150602083013567ffffffffffffffff811115611bed578182fd5b611bf9858286016118c4565b9150509250929050565b60008060408385031215611c15578182fd5b50508035926020909101359150565b600080600060608486031215611c38578081fd5b8335925060208401359150604084013567ffffffffffffffff811115611c5c578182fd5b611c68868287016118c4565b9150509250925092565b60008151808452611c8a816020860160208601611ed3565b601f01601f19169290920160200192915050565b60007f416363657373436f6e74726f6c3a206163636f756e742000000000000000000082528351611cd6816017850160208801611ed3565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611d07816028840160208801611ed3565b01602801949350505050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611d4690830184611c72565b9695505050505050565b6000602082526117496020830184611c72565b6020808252602d908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526c1c881bdc88185c1c1c9bdd9959609a1b606082015260800190565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b60208082526025908201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060408201526437bbb732b960d91b606082015260800190565b604051601f8201601f1916810167ffffffffffffffff81118282101715611e7057611e70611f82565b604052919050565b600067ffffffffffffffff821115611e9257611e92611f82565b5060209081020190565b60008219821115611eaf57611eaf611f6c565b500190565b6000816000190483118215151615611ece57611ece611f6c565b500290565b60005b83811015611eee578181015183820152602001611ed6565b838111156109965750506000910152565b600081611f0e57611f0e611f6c565b506000190190565b600281046001821680611f2a57607f821691505b60208210811415611f4b57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611f6557611f65611f6c565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160e01b0319811681146107f157600080fdfea2646970667358221220dad9c2d4cd648a67b139abc0f668455a2036c13d65662a7a7ef3fda77e7472c264736f6c63430008020033",
}

// ZooNftABI is the input ABI used to generate the binding from.
// Deprecated: Use ZooNftMetaData.ABI instead.
var ZooNftABI = ZooNftMetaData.ABI

// ZooNftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZooNftMetaData.Bin instead.
var ZooNftBin = ZooNftMetaData.Bin

// DeployZooNft deploys a new Ethereum contract, binding an instance of ZooNft to it.
func DeployZooNft(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address) (common.Address, *types.Transaction, *ZooNft, error) {
	parsed, err := ZooNftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZooNftBin), backend, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZooNft{ZooNftCaller: ZooNftCaller{contract: contract}, ZooNftTransactor: ZooNftTransactor{contract: contract}, ZooNftFilterer: ZooNftFilterer{contract: contract}}, nil
}

// ZooNft is an auto generated Go binding around an Ethereum contract.
type ZooNft struct {
	ZooNftCaller     // Read-only binding to the contract
	ZooNftTransactor // Write-only binding to the contract
	ZooNftFilterer   // Log filterer for contract events
}

// ZooNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZooNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZooNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZooNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZooNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZooNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZooNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZooNftSession struct {
	Contract     *ZooNft           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZooNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZooNftCallerSession struct {
	Contract *ZooNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZooNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZooNftTransactorSession struct {
	Contract     *ZooNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZooNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZooNftRaw struct {
	Contract *ZooNft // Generic contract binding to access the raw methods on
}

// ZooNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZooNftCallerRaw struct {
	Contract *ZooNftCaller // Generic read-only contract binding to access the raw methods on
}

// ZooNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZooNftTransactorRaw struct {
	Contract *ZooNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZooNft creates a new instance of ZooNft, bound to a specific deployed contract.
func NewZooNft(address common.Address, backend bind.ContractBackend) (*ZooNft, error) {
	contract, err := bindZooNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZooNft{ZooNftCaller: ZooNftCaller{contract: contract}, ZooNftTransactor: ZooNftTransactor{contract: contract}, ZooNftFilterer: ZooNftFilterer{contract: contract}}, nil
}

// NewZooNftCaller creates a new read-only instance of ZooNft, bound to a specific deployed contract.
func NewZooNftCaller(address common.Address, caller bind.ContractCaller) (*ZooNftCaller, error) {
	contract, err := bindZooNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZooNftCaller{contract: contract}, nil
}

// NewZooNftTransactor creates a new write-only instance of ZooNft, bound to a specific deployed contract.
func NewZooNftTransactor(address common.Address, transactor bind.ContractTransactor) (*ZooNftTransactor, error) {
	contract, err := bindZooNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZooNftTransactor{contract: contract}, nil
}

// NewZooNftFilterer creates a new log filterer instance of ZooNft, bound to a specific deployed contract.
func NewZooNftFilterer(address common.Address, filterer bind.ContractFilterer) (*ZooNftFilterer, error) {
	contract, err := bindZooNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZooNftFilterer{contract: contract}, nil
}

// bindZooNft binds a generic wrapper to an already deployed contract.
func bindZooNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZooNftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZooNft *ZooNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZooNft.Contract.ZooNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZooNft *ZooNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZooNft.Contract.ZooNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZooNft *ZooNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZooNft.Contract.ZooNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZooNft *ZooNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZooNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZooNft *ZooNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZooNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZooNft *ZooNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZooNft.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZooNft.Contract.DEFAULTADMINROLE(&_ZooNft.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZooNft.Contract.DEFAULTADMINROLE(&_ZooNft.CallOpts)
}

// NFTFACTORYROLE is a free data retrieval call binding the contract method 0x7a45fe83.
//
// Solidity: function NFT_FACTORY_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftCaller) NFTFACTORYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "NFT_FACTORY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NFTFACTORYROLE is a free data retrieval call binding the contract method 0x7a45fe83.
//
// Solidity: function NFT_FACTORY_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftSession) NFTFACTORYROLE() ([32]byte, error) {
	return _ZooNft.Contract.NFTFACTORYROLE(&_ZooNft.CallOpts)
}

// NFTFACTORYROLE is a free data retrieval call binding the contract method 0x7a45fe83.
//
// Solidity: function NFT_FACTORY_ROLE() view returns(bytes32)
func (_ZooNft *ZooNftCallerSession) NFTFACTORYROLE() ([32]byte, error) {
	return _ZooNft.Contract.NFTFACTORYROLE(&_ZooNft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ZooNft *ZooNftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ZooNft *ZooNftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ZooNft.Contract.BalanceOf(&_ZooNft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ZooNft *ZooNftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ZooNft.Contract.BalanceOf(&_ZooNft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ZooNft.Contract.GetApproved(&_ZooNft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ZooNft.Contract.GetApproved(&_ZooNft.CallOpts, tokenId)
}

// GetNftURI is a free data retrieval call binding the contract method 0xe60b943a.
//
// Solidity: function getNftURI(uint256 level, uint256 category) view returns(string)
func (_ZooNft *ZooNftCaller) GetNftURI(opts *bind.CallOpts, level *big.Int, category *big.Int) (string, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "getNftURI", level, category)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNftURI is a free data retrieval call binding the contract method 0xe60b943a.
//
// Solidity: function getNftURI(uint256 level, uint256 category) view returns(string)
func (_ZooNft *ZooNftSession) GetNftURI(level *big.Int, category *big.Int) (string, error) {
	return _ZooNft.Contract.GetNftURI(&_ZooNft.CallOpts, level, category)
}

// GetNftURI is a free data retrieval call binding the contract method 0xe60b943a.
//
// Solidity: function getNftURI(uint256 level, uint256 category) view returns(string)
func (_ZooNft *ZooNftCallerSession) GetNftURI(level *big.Int, category *big.Int) (string, error) {
	return _ZooNft.Contract.GetNftURI(&_ZooNft.CallOpts, level, category)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZooNft *ZooNftCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZooNft *ZooNftSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZooNft.Contract.GetRoleAdmin(&_ZooNft.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZooNft *ZooNftCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZooNft.Contract.GetRoleAdmin(&_ZooNft.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZooNft *ZooNftCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZooNft *ZooNftSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZooNft.Contract.HasRole(&_ZooNft.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZooNft *ZooNftCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZooNft.Contract.HasRole(&_ZooNft.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZooNft *ZooNftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZooNft *ZooNftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ZooNft.Contract.IsApprovedForAll(&_ZooNft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ZooNft *ZooNftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ZooNft.Contract.IsApprovedForAll(&_ZooNft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZooNft *ZooNftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZooNft *ZooNftSession) Name() (string, error) {
	return _ZooNft.Contract.Name(&_ZooNft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZooNft *ZooNftCallerSession) Name() (string, error) {
	return _ZooNft.Contract.Name(&_ZooNft.CallOpts)
}

// NftURI is a free data retrieval call binding the contract method 0x509f1880.
//
// Solidity: function nftURI(uint256 , uint256 ) view returns(string)
func (_ZooNft *ZooNftCaller) NftURI(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "nftURI", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NftURI is a free data retrieval call binding the contract method 0x509f1880.
//
// Solidity: function nftURI(uint256 , uint256 ) view returns(string)
func (_ZooNft *ZooNftSession) NftURI(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _ZooNft.Contract.NftURI(&_ZooNft.CallOpts, arg0, arg1)
}

// NftURI is a free data retrieval call binding the contract method 0x509f1880.
//
// Solidity: function nftURI(uint256 , uint256 ) view returns(string)
func (_ZooNft *ZooNftCallerSession) NftURI(arg0 *big.Int, arg1 *big.Int) (string, error) {
	return _ZooNft.Contract.NftURI(&_ZooNft.CallOpts, arg0, arg1)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ZooNft.Contract.OwnerOf(&_ZooNft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZooNft *ZooNftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ZooNft.Contract.OwnerOf(&_ZooNft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZooNft *ZooNftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZooNft *ZooNftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZooNft.Contract.SupportsInterface(&_ZooNft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZooNft *ZooNftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZooNft.Contract.SupportsInterface(&_ZooNft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZooNft *ZooNftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZooNft *ZooNftSession) Symbol() (string, error) {
	return _ZooNft.Contract.Symbol(&_ZooNft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZooNft *ZooNftCallerSession) Symbol() (string, error) {
	return _ZooNft.Contract.Symbol(&_ZooNft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZooNft *ZooNftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZooNft *ZooNftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ZooNft.Contract.TokenURI(&_ZooNft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZooNft *ZooNftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ZooNft.Contract.TokenURI(&_ZooNft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZooNft *ZooNftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZooNft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZooNft *ZooNftSession) TotalSupply() (*big.Int, error) {
	return _ZooNft.Contract.TotalSupply(&_ZooNft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ZooNft *ZooNftCallerSession) TotalSupply() (*big.Int, error) {
	return _ZooNft.Contract.TotalSupply(&_ZooNft.CallOpts)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string tokenURI_) returns()
func (_ZooNft *ZooNftTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "_setTokenURI", tokenId, tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string tokenURI_) returns()
func (_ZooNft *ZooNftSession) SetTokenURI(tokenId *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetTokenURI(&_ZooNft.TransactOpts, tokenId, tokenURI_)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x01538868.
//
// Solidity: function _setTokenURI(uint256 tokenId, string tokenURI_) returns()
func (_ZooNft *ZooNftTransactorSession) SetTokenURI(tokenId *big.Int, tokenURI_ string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetTokenURI(&_ZooNft.TransactOpts, tokenId, tokenURI_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Approve(&_ZooNft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Approve(&_ZooNft.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ZooNft *ZooNftSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Burn(&_ZooNft.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Burn(&_ZooNft.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.GrantRole(&_ZooNft.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.GrantRole(&_ZooNft.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Mint(&_ZooNft.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.Mint(&_ZooNft.TransactOpts, to, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.RenounceRole(&_ZooNft.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.RenounceRole(&_ZooNft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.RevokeRole(&_ZooNft.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZooNft *ZooNftTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.RevokeRole(&_ZooNft.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.SafeTransferFrom(&_ZooNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.SafeTransferFrom(&_ZooNft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ZooNft *ZooNftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ZooNft *ZooNftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZooNft.Contract.SafeTransferFrom0(&_ZooNft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ZooNft *ZooNftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ZooNft.Contract.SafeTransferFrom0(&_ZooNft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZooNft *ZooNftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZooNft *ZooNftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZooNft.Contract.SetApprovalForAll(&_ZooNft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ZooNft *ZooNftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ZooNft.Contract.SetApprovalForAll(&_ZooNft.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ZooNft *ZooNftTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ZooNft *ZooNftSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetBaseURI(&_ZooNft.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ZooNft *ZooNftTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetBaseURI(&_ZooNft.TransactOpts, _baseURI)
}

// SetBaseURIInfo is a paid mutator transaction binding the contract method 0x3e42c199.
//
// Solidity: function setBaseURIInfo(string baseURI_) returns()
func (_ZooNft *ZooNftTransactor) SetBaseURIInfo(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setBaseURIInfo", baseURI_)
}

// SetBaseURIInfo is a paid mutator transaction binding the contract method 0x3e42c199.
//
// Solidity: function setBaseURIInfo(string baseURI_) returns()
func (_ZooNft *ZooNftSession) SetBaseURIInfo(baseURI_ string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetBaseURIInfo(&_ZooNft.TransactOpts, baseURI_)
}

// SetBaseURIInfo is a paid mutator transaction binding the contract method 0x3e42c199.
//
// Solidity: function setBaseURIInfo(string baseURI_) returns()
func (_ZooNft *ZooNftTransactorSession) SetBaseURIInfo(baseURI_ string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetBaseURIInfo(&_ZooNft.TransactOpts, baseURI_)
}

// SetMultiNftURI is a paid mutator transaction binding the contract method 0x574bd108.
//
// Solidity: function setMultiNftURI(uint256[] levels, uint256[] categorys, string[] URIs) returns()
func (_ZooNft *ZooNftTransactor) SetMultiNftURI(opts *bind.TransactOpts, levels []*big.Int, categorys []*big.Int, URIs []string) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setMultiNftURI", levels, categorys, URIs)
}

// SetMultiNftURI is a paid mutator transaction binding the contract method 0x574bd108.
//
// Solidity: function setMultiNftURI(uint256[] levels, uint256[] categorys, string[] URIs) returns()
func (_ZooNft *ZooNftSession) SetMultiNftURI(levels []*big.Int, categorys []*big.Int, URIs []string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetMultiNftURI(&_ZooNft.TransactOpts, levels, categorys, URIs)
}

// SetMultiNftURI is a paid mutator transaction binding the contract method 0x574bd108.
//
// Solidity: function setMultiNftURI(uint256[] levels, uint256[] categorys, string[] URIs) returns()
func (_ZooNft *ZooNftTransactorSession) SetMultiNftURI(levels []*big.Int, categorys []*big.Int, URIs []string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetMultiNftURI(&_ZooNft.TransactOpts, levels, categorys, URIs)
}

// SetNFTFactory is a paid mutator transaction binding the contract method 0xcb2d93ce.
//
// Solidity: function setNFTFactory(address _nftFactory) returns()
func (_ZooNft *ZooNftTransactor) SetNFTFactory(opts *bind.TransactOpts, _nftFactory common.Address) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setNFTFactory", _nftFactory)
}

// SetNFTFactory is a paid mutator transaction binding the contract method 0xcb2d93ce.
//
// Solidity: function setNFTFactory(address _nftFactory) returns()
func (_ZooNft *ZooNftSession) SetNFTFactory(_nftFactory common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.SetNFTFactory(&_ZooNft.TransactOpts, _nftFactory)
}

// SetNFTFactory is a paid mutator transaction binding the contract method 0xcb2d93ce.
//
// Solidity: function setNFTFactory(address _nftFactory) returns()
func (_ZooNft *ZooNftTransactorSession) SetNFTFactory(_nftFactory common.Address) (*types.Transaction, error) {
	return _ZooNft.Contract.SetNFTFactory(&_ZooNft.TransactOpts, _nftFactory)
}

// SetNftURI is a paid mutator transaction binding the contract method 0x7a47dc77.
//
// Solidity: function setNftURI(uint256 level, uint256 category, string URI) returns()
func (_ZooNft *ZooNftTransactor) SetNftURI(opts *bind.TransactOpts, level *big.Int, category *big.Int, URI string) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "setNftURI", level, category, URI)
}

// SetNftURI is a paid mutator transaction binding the contract method 0x7a47dc77.
//
// Solidity: function setNftURI(uint256 level, uint256 category, string URI) returns()
func (_ZooNft *ZooNftSession) SetNftURI(level *big.Int, category *big.Int, URI string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetNftURI(&_ZooNft.TransactOpts, level, category, URI)
}

// SetNftURI is a paid mutator transaction binding the contract method 0x7a47dc77.
//
// Solidity: function setNftURI(uint256 level, uint256 category, string URI) returns()
func (_ZooNft *ZooNftTransactorSession) SetNftURI(level *big.Int, category *big.Int, URI string) (*types.Transaction, error) {
	return _ZooNft.Contract.SetNftURI(&_ZooNft.TransactOpts, level, category, URI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.TransferFrom(&_ZooNft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ZooNft *ZooNftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ZooNft.Contract.TransferFrom(&_ZooNft.TransactOpts, from, to, tokenId)
}

// ZooNftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZooNft contract.
type ZooNftApprovalIterator struct {
	Event *ZooNftApproval // Event containing the contract specifics and raw log

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
func (it *ZooNftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftApproval)
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
		it.Event = new(ZooNftApproval)
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
func (it *ZooNftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftApproval represents a Approval event raised by the ZooNft contract.
type ZooNftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ZooNftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftApprovalIterator{contract: _ZooNft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZooNftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftApproval)
				if err := _ZooNft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) ParseApproval(log types.Log) (*ZooNftApproval, error) {
	event := new(ZooNftApproval)
	if err := _ZooNft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ZooNft contract.
type ZooNftApprovalForAllIterator struct {
	Event *ZooNftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ZooNftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftApprovalForAll)
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
		it.Event = new(ZooNftApprovalForAll)
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
func (it *ZooNftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftApprovalForAll represents a ApprovalForAll event raised by the ZooNft contract.
type ZooNftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZooNft *ZooNftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ZooNftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftApprovalForAllIterator{contract: _ZooNft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZooNft *ZooNftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ZooNftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftApprovalForAll)
				if err := _ZooNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ZooNft *ZooNftFilterer) ParseApprovalForAll(log types.Log) (*ZooNftApprovalForAll, error) {
	event := new(ZooNftApprovalForAll)
	if err := _ZooNft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZooNft contract.
type ZooNftInitializedIterator struct {
	Event *ZooNftInitialized // Event containing the contract specifics and raw log

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
func (it *ZooNftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftInitialized)
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
		it.Event = new(ZooNftInitialized)
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
func (it *ZooNftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftInitialized represents a Initialized event raised by the ZooNft contract.
type ZooNftInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZooNft *ZooNftFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZooNftInitializedIterator, error) {

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZooNftInitializedIterator{contract: _ZooNft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZooNft *ZooNftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZooNftInitialized) (event.Subscription, error) {

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftInitialized)
				if err := _ZooNft.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZooNft *ZooNftFilterer) ParseInitialized(log types.Log) (*ZooNftInitialized, error) {
	event := new(ZooNftInitialized)
	if err := _ZooNft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZooNft contract.
type ZooNftRoleAdminChangedIterator struct {
	Event *ZooNftRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZooNftRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftRoleAdminChanged)
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
		it.Event = new(ZooNftRoleAdminChanged)
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
func (it *ZooNftRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftRoleAdminChanged represents a RoleAdminChanged event raised by the ZooNft contract.
type ZooNftRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZooNft *ZooNftFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZooNftRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftRoleAdminChangedIterator{contract: _ZooNft.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZooNft *ZooNftFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZooNftRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftRoleAdminChanged)
				if err := _ZooNft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZooNft *ZooNftFilterer) ParseRoleAdminChanged(log types.Log) (*ZooNftRoleAdminChanged, error) {
	event := new(ZooNftRoleAdminChanged)
	if err := _ZooNft.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZooNft contract.
type ZooNftRoleGrantedIterator struct {
	Event *ZooNftRoleGranted // Event containing the contract specifics and raw log

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
func (it *ZooNftRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftRoleGranted)
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
		it.Event = new(ZooNftRoleGranted)
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
func (it *ZooNftRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftRoleGranted represents a RoleGranted event raised by the ZooNft contract.
type ZooNftRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZooNftRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftRoleGrantedIterator{contract: _ZooNft.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZooNftRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftRoleGranted)
				if err := _ZooNft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) ParseRoleGranted(log types.Log) (*ZooNftRoleGranted, error) {
	event := new(ZooNftRoleGranted)
	if err := _ZooNft.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZooNft contract.
type ZooNftRoleRevokedIterator struct {
	Event *ZooNftRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZooNftRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftRoleRevoked)
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
		it.Event = new(ZooNftRoleRevoked)
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
func (it *ZooNftRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftRoleRevoked represents a RoleRevoked event raised by the ZooNft contract.
type ZooNftRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZooNftRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftRoleRevokedIterator{contract: _ZooNft.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZooNftRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftRoleRevoked)
				if err := _ZooNft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZooNft *ZooNftFilterer) ParseRoleRevoked(log types.Log) (*ZooNftRoleRevoked, error) {
	event := new(ZooNftRoleRevoked)
	if err := _ZooNft.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZooNftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZooNft contract.
type ZooNftTransferIterator struct {
	Event *ZooNftTransfer // Event containing the contract specifics and raw log

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
func (it *ZooNftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZooNftTransfer)
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
		it.Event = new(ZooNftTransfer)
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
func (it *ZooNftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZooNftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZooNftTransfer represents a Transfer event raised by the ZooNft contract.
type ZooNftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ZooNftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZooNft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ZooNftTransferIterator{contract: _ZooNft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZooNftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ZooNft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZooNftTransfer)
				if err := _ZooNft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ZooNft *ZooNftFilterer) ParseTransfer(log types.Log) (*ZooNftTransfer, error) {
	event := new(ZooNftTransfer)
	if err := _ZooNft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
