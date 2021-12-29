package substrate

import (
	"math/big"
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/common"
)

var api *gsrpc.SubstrateAPI

func init() {
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
				Key: &signature.TestKeyringPairAlice,
			}
			if err := c.SubmitTx(RegisterWatcher); (err != nil) != tt.wantErr {
				t.Errorf("RegisterWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnection_UpdateStakeInfo(t *testing.T) {
	bob, err := types.NewAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	if err != nil {
		panic(err)
	}
	address := common.HexToAddress("0xa7f6c9a5052a08a14ff0e3349094b6efbc591ea4")

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
				key: &signature.TestKeyringPairAlice,
			},
			args: args{
				stakeInfos: []*StakeInfo{
					{
						Coinbase:      bob.AsAccountID,
						WorkBase:      address[:],
						IsWork:        true,
						LockedBalance: types.NewU128(*big.NewInt(123456)),
						WorkCount:     10,
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
	phrase := " indoor height cinnamon parent kite oxygen dolphin output pet bitter joke grain"
	seed := "0xe1d5a01954b8320d8c5ceb88199487b5a3821bbc4b520286360a71a946f22c33"
	kp1, err := signature.KeyringPairFromSecret(phrase, 2)
	if err != nil {
		t.Error(err)
	}
	t.Logf("kp1: %+v\n", kp1)

	kp2, err := signature.KeyringPairFromSecret(seed, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("kp2: %+v\n", kp2)
}
