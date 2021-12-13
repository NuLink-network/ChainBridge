package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (es EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

const (
	// Deposited event: Deposited(address indexed staker, uint256 value, uint16 periods)
	Deposited EventSig = "Deposited(address,uint256,uint16)"
)
