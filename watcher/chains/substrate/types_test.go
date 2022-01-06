package substrate

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	b               = &big.Int{}
	wantStakeInfo12 = StakeInfos{}
	wantStakeInfo19 = StakeInfos{}
	wantStakeInfo20 = StakeInfos{}
	address1        = common.HexToAddress("0xa7f6c9a5052a08a14ff0e3349094b6efbc591ea4")
)

func Init() {
	var ok bool
	// 1e28
	b, ok = new(big.Int).SetString("10000000000000000000000000000", 10)
	if !ok {
		panic("big.Int SetString failed")
	}

	wantStakeInfo12 = StakeInfos{
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(30), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(29), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(28), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(27), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(26), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(25), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(24), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(23), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(22), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(21), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(20), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(19), b)),
			WorkCount:     1,
		},
	}
	wantStakeInfo19 = StakeInfos{
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(30), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(29), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(28), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(27), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(26), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(25), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(24), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(23), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(22), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(21), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(20), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(19), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(18), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(17), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(16), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(15), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(14), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(13), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(12), b)),
			WorkCount:     1,
		},
	}
	wantStakeInfo20 = StakeInfos{
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(30), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(29), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(28), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(27), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(26), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(25), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(24), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(23), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(22), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(21), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(20), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(19), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(18), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(17), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(16), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(15), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(14), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(13), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(12), b)),
			WorkCount:     1,
		},
		{
			Coinbase:      types.NewAccountID([]byte(address1.Hex())),
			WorkBase:      address1[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(11), b)),
			WorkCount:     1,
		},
	}
}

func TestStakeInfos_LockedBalanceTop20(t *testing.T) {
	Init()

	tests := []struct {
		name string
		s    StakeInfos
		want []*StakeInfo
	}{
		{
			name: "has-20-get-20",
			s:    wantStakeInfo20,
			want: wantStakeInfo20,
		},
		{
			name: "has-20-get-20-case-2",
			s: append(wantStakeInfo19, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
			}...),
			want: append(wantStakeInfo19, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
			}...),
		},
		{
			name: "has-21-get-20",
			s: append(wantStakeInfo20, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
			}...),
			want: wantStakeInfo20,
		},
		{
			name: "has-23-get-20",
			s: append(wantStakeInfo20, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(2), b)),
					WorkCount:     1,
				},
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(3), b)),
					WorkCount:     1,
				},
			}...),
			want: wantStakeInfo20,
		},
		{
			name: "has-13-get-13",
			s: append(wantStakeInfo12, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
			}...),
			want: append(wantStakeInfo12, []*StakeInfo{
				{
					Coinbase:      types.NewAccountID([]byte(address1.Hex())),
					WorkBase:      address1[:],
					IsWork:        true,
					LockedBalance: types.NewU128(*new(big.Int).Mul(big.NewInt(1), b)),
					WorkCount:     1,
				},
			}...),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.LockedBalanceTop20(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LockedBalanceTop20() = %v, want %v", got, tt.want)
			}
		})
	}
}
