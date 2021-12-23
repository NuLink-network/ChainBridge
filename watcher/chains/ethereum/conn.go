package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	URL    string
	Http   bool
	Client *ethclient.Client
	Stop   chan struct{}
}

func NewConnection(endpoint string, http bool, stop chan struct{}) *Connection {
	return &Connection{
		URL:  endpoint,
		Http: http,
		Stop: stop,
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	log.Info("Connecting to ethereum chain...", "url", c.URL)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.Http {
		rpcClient, err = rpc.DialHTTP(c.URL)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), c.URL)
	}
	if err != nil {
		return err
	}
	c.Client = ethclient.NewClient(rpcClient)
	return nil
}

func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.Client != nil {
		c.Client.Close()
	}
	close(c.Stop)
}
