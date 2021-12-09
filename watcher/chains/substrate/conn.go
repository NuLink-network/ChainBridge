package substrate

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
	"github.com/centrifuge/go-substrate-rpc-client/types"
	"github.com/ethereum/go-ethereum/log"

	"github.com/NuLink-network/watcher/watcher/utils"
)

type Connection struct {
	api  *gsrpc.SubstrateAPI
	url  string                 // API endpoint
	key  *signature.KeyringPair // Keyring used for signing
	stop chan int               // Signals system shutdown, should be observed in all selects and loops
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

func (c *Connection) SubmitTx(method utils.Method, args ...interface{}) error {
	log.Info("Submitting substrate call...", "method", method, "sender", c.key.Address)

	meta, err := c.api.RPC.State.GetMetadataLatest()
	if err != nil {
		return fmt.Errorf("failed get the latest metadata, err: %v", err)
	}

	// Create call and extrinsic
	call, err := types.NewCall(meta, string(method), args...)
	if err != nil {
		return fmt.Errorf("failed to construct call, err: %v", err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(call)

	genesisHash, err := c.api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return fmt.Errorf("failed to get the genesis hash, err: %v", genesisHash)
	}
	// Get latest runtime version
	rv, err := c.api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", c.key.PublicKey, nil)
	if err != nil {
		return fmt.Errorf("create storage key failed, err: %v", err)
	}

	var accountInfo types.AccountInfo
	ok, err := c.api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return fmt.Errorf("failed to get the latest storage, err: %v", err)
	}

	nonce := uint32(accountInfo.Nonce)

	// Sign the extrinsic
	opts := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}

	err = ext.Sign(*c.key, opts)
	if err != nil {
		return err
	}

	// Send the extrinsic
	hash, err := c.api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		return fmt.Errorf("submit of extrinsic failed: %v", err)
	}
	log.Info("submit extrinsic succeeded", "hash", hash)

	return nil
}
