package substrate

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

type StakeInfo struct {
	Coinbase      [32]byte
	WorkBase      [32]byte
	IsWork        bool
	LockedBalance types.U128
	WorkCount     uint32
}
