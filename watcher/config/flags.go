package config

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

var (
	VerbosityFlag = &cli.StringFlag{
		Name:  "verbosity",
		Usage: "Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail",
		Value: log.LvlInfo.String(),
	}

	ConfigFileFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "JSON configuration file",
	}

	StakeInfoFileFlag = &cli.StringFlag{
		Name:  "file",
		Usage: "Store last stake info file",
		Value: DefaultStakeInfoFile(),
	}

	BlockStoreFileFlag = &cli.StringFlag{
		Name:  "blockstore",
		Usage: "Store last block umber file",
		Value: DefaultLatestBlockFile(),
	}
)
