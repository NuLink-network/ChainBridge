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
			if err := c.RegisterWatcher(); (err != nil) != tt.wantErr {
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
						WorkBase:      common.BigToHash(big.NewInt(123)),
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
			if err := c.UpdateStakeInfo(tt.args.stakeInfos); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStakeInfo() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
