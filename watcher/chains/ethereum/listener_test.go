package ethereum

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/rlp"
	"io/ioutil"
	"math/big"
	"testing"
)

func TestRLPEncodeAndDecode(t *testing.T) {
	type StakeInfo struct {
		Coinbase      [32]byte
		WorkBase      []byte
		IsWork        bool
		LockedBalance types.U128
		WorkCount     uint32
	}
	bob, _ := types.NewAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	var infos = StakeInfo{
		Coinbase:      bob.AsAccountID,
		WorkBase:      bob.AsAccountID[:],
		IsWork:        true,
		LockedBalance: types.NewU128(*big.NewInt(8888)),
		WorkCount:     0,
	}

	data, err := rlp.EncodeToBytes(infos)
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile("/Users/t/Library/NuLinkWatcher/stake-info.rlp", data, 0664); err != nil {
		t.Fatal(err)
	}

	file, err := ioutil.ReadFile("/Users/t/Library/NuLinkWatcher/stake-info.rlp")
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Log("file is empty...")
		return
	}

	var si StakeInfo
	if err := rlp.DecodeBytes(file, &si); err != nil {
		t.Fatal(err)
	}
	t.Logf("============================== si: %+v\n", si)
}

func TestTestRLPDecode(t *testing.T) {
	type StakeInfo struct {
		Coinbase      [32]byte
		WorkBase      []byte
		IsWork        bool
		LockedBalance types.U128
		WorkCount     uint32
	}

	file, err := ioutil.ReadFile("./stake-info.rlp")
	if err != nil {
		t.Fatal(err)
	}
	if len(file) == 0 {
		t.Log("file is empty...")
		return
	}

	var si StakeInfo
	if err := rlp.DecodeBytes(file, &si); err != nil {
		t.Fatal(err)
	}
	t.Logf("============================== si: %+v\n", si)
}
