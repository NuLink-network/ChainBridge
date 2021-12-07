package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type Connection struct {
	endpoint string
	http     bool
	client   *ethclient.Client
	stop     chan struct{}
}

func NewConnection(endpoint string, http bool) *Connection {
	return &Connection{
		endpoint: endpoint,
		http:     http,
	}
}

// Connect starts the ethereum WS connection
func (c *Connection) Connect() error {
	log.Info("Connecting to ethereum chain...", "url", c.endpoint)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if c.http {
		rpcClient, err = rpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), c.endpoint)
	}
	if err != nil {
		return err
	}
	c.client = ethclient.NewClient(rpcClient)
	return nil
}

func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.client != nil {
		c.client.Close()
	}
	close(c.stop)
}
