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
	Watcher = &signature.KeyringPair{
		URI:       "0xe1d5a01954b8320d8c5ceb88199487b5a3821bbc4b520286360a71a946f22c33",
		Address:   "5EZQrJ2XBSNnzQ78jEBiRZj5S4SMEczZJMc8ANL54AWdtz9n",
		PublicKey: []byte{110, 93, 85, 181, 154, 147, 45, 198, 166, 76, 54, 68, 31, 165, 117, 6, 165, 42, 163, 142, 162, 20, 255, 118, 230, 14, 155, 9, 163, 214, 222, 121},
	}
)
