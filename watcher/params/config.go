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
	EpochLength     = 10000
)

var BlockRetryInterval = time.Second * 5

var (
	StartBlock         = big.NewInt(10000000)
	BlockConfirmations = big.NewInt(12)
)

var (
	DepositContractAddress = common.HexToAddress("0x52FC9c85678Fd16F63eC4a5a5E18674CD4C3F83c")
)
