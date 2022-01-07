package substrate

import (
	"fmt"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/log"
)

type Connection struct {
	API  *gsrpc.SubstrateAPI
	URL  string                 // API endpoint
	Key  *signature.KeyringPair // Keyring used for signing
	Stop chan struct{}          // Signals system shutdown, should be observed in all selects and loops
}

func NewConnection(url string, key *signature.KeyringPair, stop chan struct{}) *Connection {
	return &Connection{
		URL:  url,
		Key:  key,
		Stop: stop,
	}
}

func (c *Connection) Connect() error {
	log.Info("Connecting to substrate chain...", "url", c.URL)
	api, err := gsrpc.NewSubstrateAPI(c.URL)
	if err != nil {
		return err
	}
	c.API = api

	return nil
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {

}

func (c *Connection) SubmitTx(method Method, args ...interface{}) error {
	//c.Key = &signature.TestKeyringPairAlice
	log.Info("Submitting substrate call...", "method", method, "sender", c.Key.Address)

	meta, err := c.API.RPC.State.GetMetadataLatest()
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

	genesisHash, err := c.API.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return fmt.Errorf("failed to get the genesis hash, err: %v", genesisHash)
	}
	// Get latest runtime version
	rv, err := c.API.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return err
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", c.Key.PublicKey, nil)
	if err != nil {
		return fmt.Errorf("create storage key failed, err: %v", err)
	}

	var accountInfo types.AccountInfo
	ok, err := c.API.RPC.State.GetStorageLatest(key, &accountInfo)
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

	err = ext.Sign(*c.Key, opts)
	if err != nil {
		return err
	}

	// Send the extrinsic
	hash, err := c.API.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		return fmt.Errorf("submit of extrinsic failed: %v", err)
	}
	log.Info("submit extrinsic succeeded", "hash", hash.Hex())

	return nil
}
