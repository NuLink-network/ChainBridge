package substrate

import (
	"math/big"
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/NuLink-network/watcher/watcher/params"
)

var api *gsrpc.SubstrateAPI

func Connect() {
	var err error
	api, err = gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
}

func getKeyringPair() *signature.KeyringPair {
	//phrase := " indoor height cinnamon parent kite oxygen dolphin output pet bitter joke grain"
	seed := "0xe1d5a01954b8320d8c5ceb88199487b5a3821bbc4b520286360a71a946f22c33"

	kp, err := signature.KeyringPairFromSecret(seed, 1)
	if err != nil {
		panic(err)
	}
	return &kp
}

func TestConnection_RegisterWatcher(t *testing.T) {
	Connect()

	type fields struct {
		api  *gsrpc.SubstrateAPI
		url  string
		key  *signature.KeyringPair
		stop chan int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "t-1",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
			if err != nil {
				t.Fatal(err)
			}

			c := &Connection{
				API: api,
				Key: params.Watcher,
			}
			if err := c.SubmitTx(RegisterWatcher); (err != nil) != tt.wantErr {
				t.Errorf("RegisterWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnection_UpdateStakeInfo(t *testing.T) {
	Connect()
	//type StakeInfo struct {
	//	Coinbase      [32]byte
	//	WorkBase      []byte
	//	IsWork        bool
	//	LockedBalance *big.Int
	//	WorkCount     uint32
	//}
	//bob, err := types.NewAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	//if err != nil {
	//	panic(err)
	//}
	//accountID := [32]byte{142, 175, 4, 21, 22, 135, 115, 99, 38, 201, 254, 161, 126, 37, 252, 82, 135, 97, 54, 147, 201, 18, 144, 156, 178, 38, 170, 71, 148, 242, 106, 72}
	address1 := common.HexToAddress("0xa7f6c9a5052a08a14ff0e3349094b6efbc591ea4")
	address2 := common.HexToAddress("0x00192fb10df37c9fb26829eb2cc623cd1bf599e8")

	type fields struct {
		api  *gsrpc.SubstrateAPI
		url  string
		key  *signature.KeyringPair
		stop chan int
	}
	type args struct {
		stakeInfos []*StakeInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "t-1",
			fields: fields{
				api: api,
				key: params.Watcher,
			},
			args: args{
				stakeInfos: []*StakeInfo{
					{
						Coinbase:      types.NewAccountID([]byte(address1.Hex())),
						WorkBase:      address1[:],
						IsWork:        true,
						LockedBalance: types.NewU128(*big.NewInt(111111)),
						WorkCount:     1,
					},
					{
						Coinbase:      types.NewAccountID([]byte(address2.Hex())),
						WorkBase:      address2[:],
						IsWork:        true,
						LockedBalance: types.NewU128(*big.NewInt(222222)),
						WorkCount:     2,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				API: tt.fields.api,
				Key: tt.fields.key,
			}
			if err := c.SubmitTx(UpdateStakeInfo, tt.args.stakeInfos); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStakeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestGetAddressFromAccountID(t *testing.T) {
	// galen
	//seed := "0xe1d5a01954b8320d8c5ceb88199487b5a3821bbc4b520286360a71a946f22c33"
	// watcher
	seed := "0xb233c8411f90fa880d3f23a90af5e89896b73bf8ce95217f310c5f497ab3aaf6"

	kp0, err := signature.KeyringPairFromSecret(seed, 0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("kp0: %+v\n", kp0)

	kp3, err := signature.KeyringPairFromSecret(seed, 3)
	if err != nil {
		t.Error(err)
	}
	t.Logf("kp3: %+v\n", kp3)

	kp100, err := signature.KeyringPairFromSecret(seed, 100)
	if err != nil {
		t.Error(err)
	}
	t.Logf("kp100: %+v\n", kp100)
}

func TestNewAccountFromHexAddress(t *testing.T) {
	addr1 := "0x4f51607f575F13bc661B44499b99A7EA3e0cEA8A"
	addr2 := "0xd55A80F033F7d6AbCac39A936ba2cC7Cb38ccb87"

	accountID1 := types.NewAccountID([]byte(addr1))
	accountID2 := types.NewAccountID([]byte(addr2))
	t.Log("account id 1: ", accountID1)
	t.Log("account id 2: ", accountID2)
}
