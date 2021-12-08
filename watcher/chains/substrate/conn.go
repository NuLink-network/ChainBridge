package substrate

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/ethereum/go-ethereum/log"
)

type Connection struct {
	api  *gsrpc.SubstrateAPI
	url  string   // API endpoint
	stop chan int // Signals system shutdown, should be observed in all selects and loops
}

func NewConnection(url string) *Connection {
	return &Connection{url: url}
}

func (c *Connection) Connect() error {
	log.Info("Connecting to substrate chain...", "url", c.url)
	api, err := gsrpc.NewSubstrateAPI(c.url)
	if err != nil {
		return err
	}
	c.api = api

	return nil
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	close(c.stop)
}
