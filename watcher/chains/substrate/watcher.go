package substrate

func (c *Connection) RegisterWatcher() error {
	// todo construct request parameters

	return c.SubmitTx(RegisterWatcher)
}

func (c *Connection) UpdateStakeInfo(stakeInfos []*StakeInfo) error {
	return c.SubmitTx(UpdateStakeInfo, stakeInfos)
}
