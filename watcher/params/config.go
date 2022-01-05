package params

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
