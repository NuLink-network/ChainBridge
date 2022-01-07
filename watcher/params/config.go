package params

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
)

var (
	WSEndpoint   = ""
	HTTPEndpoint = ""
)

var (
	BlockRetryLimit = 5
	EpochLength     = 1000
)

var BlockRetryInterval = time.Second * 5

var (
	StartBlock         = big.NewInt(10000000)
	BlockConfirmations = big.NewInt(12)
)

var (
	DepositContractAddress = common.HexToAddress("0xBaE99723a3F2374DB1D8287A8679f38aa9c0674d")
)

var (
	Galen = &signature.KeyringPair{
		URI:       "0xe1d5a01954b8320d8c5ceb88199487b5a3821bbc4b520286360a71a946f22c33",
		Address:   "5EZQrJ2XBSNnzQ78jEBiRZj5S4SMEczZJMc8ANL54AWdtz9n",
		PublicKey: []byte{110, 93, 85, 181, 154, 147, 45, 198, 166, 76, 54, 68, 31, 165, 117, 6, 165, 42, 163, 142, 162, 20, 255, 118, 230, 14, 155, 9, 163, 214, 222, 121},
	}
	Watcher = &signature.KeyringPair{
		URI:       "0xb233c8411f90fa880d3f23a90af5e89896b73bf8ce95217f310c5f497ab3aaf6",
		Address:   "5GVzGPGT2M1sU6eTPxcEkjv4TdHxSNzmvuq4v7xHFWeV2NgK",
		PublicKey: []byte{196, 57, 78, 98, 99, 75, 136, 144, 122, 244, 88, 154, 171, 188, 174, 43, 209, 67, 85, 59, 179, 24, 144, 175, 63, 121, 48, 194, 47, 132, 95, 65},
	}
)
