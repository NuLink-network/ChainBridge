package substrate

import (
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client"
	"github.com/centrifuge/go-substrate-rpc-client/config"
	"github.com/centrifuge/go-substrate-rpc-client/signature"
)

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
				api: api,
				key: &signature.TestKeyringPairAlice,
			}
			if err := c.RegisterWatcher(); (err != nil) != tt.wantErr {
				t.Errorf("RegisterWatcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
