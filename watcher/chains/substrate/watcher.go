package substrate

import "github.com/centrifuge/go-substrate-rpc-client/v4/types"

func (c *Connection) RegisterWatcher() error {
	// todo construct request parameters

	return c.SubmitTx(RegisterWatcher)
}

type StakeInfo struct {
	Coinbase      [32]byte
	WorkBase      [32]byte
	IsWork        bool
	LockedBalance types.UCompact
	WorkCount     uint32
}

func (c *Connection) UpdateStakeInfo(stakeInfos []*StakeInfo) error {
	s, err := types.EncodeToBytes(stakeInfos)
	if err != nil {
		return err
	}

	return c.SubmitTx(UpdateStakeInfo, s)
}
