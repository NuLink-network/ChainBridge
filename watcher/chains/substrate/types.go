package substrate

import (
	"sort"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type StakeInfo struct {
	Coinbase      [32]byte
	WorkBase      []byte
	IsWork        bool
	LockedBalance types.U128
	WorkCount     uint32
}

type StakeInfos []*StakeInfo

func (s StakeInfos) Len() int      { return len(s) }
func (s StakeInfos) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s StakeInfos) Less(i, j int) bool {
	return s[i].LockedBalance.Int.Cmp(s[j].LockedBalance.Int) > 0
}

func (s StakeInfos) LockedBalanceTop20() []*StakeInfo {
	sort.Sort(s)
	if s.Len() > 20 {
		return s[:20]
	}
	return s
}
