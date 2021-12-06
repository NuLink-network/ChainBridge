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
)

var BlockRetryInterval = time.Second * 5

var (
	StartBlock         = big.NewInt(10000000)
	BlockConfirmations = big.NewInt(12)
)

var (
	DepositContractAddress = common.HexToAddress("")
)
