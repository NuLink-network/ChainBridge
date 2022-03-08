package substrate

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"strings"

	"github.com/centrifuge/go-substrate-rpc-client/v4/hash"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/xxhash"
)

func CreateStoreKey(prefix, method string, publicKey []byte) (types.StorageKey, error) {
	expectedKeyBuilder := strings.Builder{}
	hexStr, err := types.Hex(xxhash.New128([]byte(prefix)).Sum(nil))
	if err != nil {
		return nil, err
	}
	expectedKeyBuilder.WriteString(hexStr)
	hexStr, err = types.Hex(xxhash.New128([]byte(method)).Sum(nil))
	if err != nil {
		return nil, err
	}
	expectedKeyBuilder.WriteString(strings.TrimPrefix(hexStr, "0x"))

	if publicKey == nil {
		return types.NewStorageKey(types.MustHexDecodeString(expectedKeyBuilder.String())), nil
	}

	accountIdHasher, err := hash.NewBlake2b128Concat(nil)
	if err != nil {
		return nil, err
	}
	accountIdSerialized := types.MustHexDecodeString(types.HexEncodeToString(publicKey))
	_, err = accountIdHasher.Write(accountIdSerialized)
	if err != nil {
		return nil, err
	}
	hexStr, err = types.Hex(accountIdHasher.Sum(nil))
	if err != nil {
		return nil, err
	}
	expectedKeyBuilder.WriteString(strings.TrimPrefix(hexStr, "0x"))

	return types.NewStorageKey(types.MustHexDecodeString(expectedKeyBuilder.String())), nil
}

func (c *Connection) isExistWatcher() (bool, error) {
	keyPrefix, err := CreateStoreKey(NuProxy, Watchers, nil)
	if err != nil {
		return false, err
	}
	keys, err := c.API.RPC.State.GetKeysLatest(keyPrefix)
	if err != nil {
		return false, err
	}
	return len(keys) > 0, nil
}

func (c *Connection) isWatcher() (bool, error) {
	key, err := CreateStoreKey(NuProxy, Watchers, c.Key.PublicKey)
	if err != nil {
		return false, err
	}

	var n uint32
	ok, err := c.API.RPC.State.GetStorageLatest(key, &n)
	if err != nil {
		return false, fmt.Errorf("failed to get the latest storage, err: %v", err)
	}
	if !ok {
		return false, nil
	}
	return n > 0, nil
}

func (c *Connection) RegisterWatcher() error {
	w, err := c.isWatcher()
	if err != nil {
		return err
	}
	if w {
		log.Info("repeat register watcher", "watcher", c.Key.Address)
		return nil
	}

	exist, err := c.isExistWatcher()
	if err != nil {
		return err
	}
	if exist {
		return errors.New("watcher already exists")
	}

	return c.SubmitTx(RegisterWatcher)
}
