package substrate

import (
	"fmt"
	"github.com/NuLink-network/watcher/watcher/params"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/log"
	"os"
	"reflect"
	"testing"
)

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(4)
	log.Root().SetHandler(glogger)
}

func wantKey(t *testing.T, prefix, method string, publicKey []byte) types.StorageKey {
	Connect()
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatalf("failed get the latest metadata, err: %v", err)
	}
	key, err := types.CreateStorageKey(meta, prefix, method, publicKey, nil)
	if err != nil {
		t.Fatalf("create storage key failed, err: %v", err)
	}
	return key
}

func TestCreateStoreKey(t *testing.T) {
	type args struct {
		prefix    string
		method    string
		publicKey []byte
	}
	tests := []struct {
		name     string
		args     args
		wantFunc func(t *testing.T, prefix, method string, publicKey []byte) types.StorageKey
		wantErr  bool
	}{
		{
			name: "t-1",
			args: args{
				prefix:    "NulinkNuproxy",
				method:    "Watchers",
				publicKey: params.Watcher.PublicKey,
			},
			wantFunc: wantKey,
			wantErr:  false,
		},
		{
			name: "t-2",
			args: args{
				prefix:    "NulinkNuproxy",
				method:    "Rewards",
				publicKey: params.Watcher.PublicKey,
			},
			wantFunc: wantKey,
			wantErr:  false,
		},
		{
			name: "t-3",
			args: args{
				prefix:    "System",
				method:    "Account",
				publicKey: params.Watcher.PublicKey,
			},
			wantFunc: wantKey,
			wantErr:  false,
		},
		//{
		//	name: "t-4",
		//	args: args{
		//		prefix:    "NulinkNuproxy",
		//		method:    "Watchers",
		//		publicKey: nil,
		//	},
		//	wantFunc: wantKey,
		//	wantErr:  false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateStoreKey(tt.args.prefix, tt.args.method, tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateStoreKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("============================== got hex: ", got.Hex())
			want := tt.wantFunc(t, tt.args.prefix, tt.args.method, tt.args.publicKey)
			fmt.Println("============================== want hex: ", want.Hex())
			if !reflect.DeepEqual(got, want) {
				t.Errorf("CreateStoreKey() got = %v, want %v", got.Hex(), want.Hex())
			}
		})
	}
}

func TestConnection_isExistWatcher(t *testing.T) {
	Connect()

	type fields struct {
		API *gsrpc.SubstrateAPI
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "t-1",
			fields: fields{
				API: api,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &Connection{
				API: tt.fields.API,
			}
			got, err := c.isExistWatcher()
			if (err != nil) != tt.wantErr {
				t.Errorf("isExistWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isExistWatcher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_isWatcher(t *testing.T) {
	Connect()

	type fields struct {
		API *gsrpc.SubstrateAPI
		URL string
		Key *signature.KeyringPair
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "t-1",
			fields: fields{
				API: api,
				Key: params.Watcher,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "t-2",
			fields: fields{
				API: api,
				Key: &signature.TestKeyringPairAlice,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				API: tt.fields.API,
				URL: tt.fields.URL,
				Key: tt.fields.Key,
			}
			got, err := c.isWatcher()
			if (err != nil) != tt.wantErr {
				t.Errorf("isWatcher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isWatcher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnection_JudgeAndRegisterWatcher(t *testing.T) {
	Connect()
	type fields struct {
		API *gsrpc.SubstrateAPI
		Key *signature.KeyringPair
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "t-1",
			fields: fields{
				API: api,
				Key: &signature.TestKeyringPairAlice,
			},
			wantErr: false,
		},
		{
			name: "t-2",
			fields: fields{
				API: api,
				Key: params.Watcher,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connection{
				API: tt.fields.API,
				Key: tt.fields.Key,
			}
			if err := c.RegisterWatcher(); (err != nil) != tt.wantErr {
				t.Errorf("RegisterWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
