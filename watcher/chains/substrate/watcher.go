package substrate

func (c *Connection) RegisterWatcher() error {
	// todo construct request parameters

	return c.SubmitTx(RegisterWatcher)
}
