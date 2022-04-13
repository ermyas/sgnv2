package nftbr

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type NFTMsg struct {
	MsgType uint8
	User    Addr
	Nft     Addr
	Id      *big.Int
	Uri     string
}

// for log/debug only
func (n *NFTMsg) String() string {
	return fmt.Sprintf("%d %x %x %s %s", n.MsgType, n.User, n.Nft, n.Id, n.Uri)
}

// raw is abi encoded NFTMsg
func DecodeNFTMsg(raw []byte) *NFTMsg {
	ret := new(NFTMsg)
	// "name": "nftMsg",
	val, _ := nftMsgAbi.Methods["nftMsg"].Inputs.Unpack(raw)
	abi.ConvertType(val[0], ret)
	return ret
}

// add func nftMsg(_in NFTMsg) pure view to NFTBridge.sol and hardhat compile, below is from artifact json
var nftMsgAbi, _ = abi.JSON(strings.NewReader(`
[
  {
	"inputs": [
	  {
		"components": [
		  {
			"internalType": "enum NFTBridge.MsgType",
			"name": "msgType",
			"type": "uint8"
		  },
		  {
			"internalType": "address",
			"name": "user",
			"type": "address"
		  },
		  {
			"internalType": "address",
			"name": "nft",
			"type": "address"
		  },
		  {
			"internalType": "uint256",
			"name": "id",
			"type": "uint256"
		  },
		  {
			"internalType": "string",
			"name": "uri",
			"type": "string"
		  }
		],
		"internalType": "struct NFTBridge.NFTMsg",
		"name": "_in",
		"type": "tuple"
	  }
	],
	"name": "nftMsg",
	"outputs": [],
	"stateMutability": "pure",
	"type": "function"
  }
]`))
