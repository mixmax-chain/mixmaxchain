package systemcontract

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/params"
	//	"math/big"
	"strings"
)

const valid_user_abi = `[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "addr",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "flag",
				"type": "uint256"
			}
		],
		"name": "add_black_list",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address[]",
				"name": "addrs",
				"type": "address[]"
			},
			{
				"internalType": "uint256[]",
				"name": "flags",
				"type": "uint256[]"
			}
		],
		"name": "bulk_add_black_list",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "init",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "addr",
				"type": "address"
			}
		],
		"name": "checkValidUser",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`

var (
	ValidatorsUserAddr = common.HexToAddress("0xFC3e5d6C419d3d855B8a2d256F76de4af3EFdC47")
	Valid_user_abi2    abi.ABI
)

func init() {
	Valid_user_abi2, _ = abi.JSON(strings.NewReader(valid_user_abi))
	//	valid_user_abi2 = a
}
