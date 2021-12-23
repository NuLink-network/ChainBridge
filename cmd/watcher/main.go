// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only
/*
Provides the command-line interface for the watcher application.

For configuration and CLI commands see the README: https://github.com/ChainSafe/ChainBridge.
*/
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NuLink-network/watcher/watcher/chains/ethereum"
	"github.com/NuLink-network/watcher/watcher/chains/substrate"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/ethereum/go-ethereum/log"
)

//var app = cli.NewApp()
//
//var cliFlags = []cli.Flag{
//	config.ConfigFileFlag,
//	config.VerbosityFlag,
//	config.KeystorePathFlag,
//	config.BlockstorePathFlag,
//	config.FreshStartFlag,
//	config.LatestBlockFlag,
//	config.MetricsFlag,
//	config.MetricsPort,
//}
//
//var generateFlags = []cli.Flag{
//	config.PasswordFlag,
//	config.Sr25519Flag,
//	config.Secp256k1Flag,
//	config.SubkeyNetworkFlag,
//}
//
//var devFlags = []cli.Flag{
//	config.TestKeyFlag,
//}
//
//var importFlags = []cli.Flag{
//	config.EthereumImportFlag,
//	config.PrivateKeyFlag,
//	config.Sr25519Flag,
//	config.Secp256k1Flag,
//	config.PasswordFlag,
//	config.SubkeyNetworkFlag,
//}
//
//var accountCommand = cli.Command{
//	Name:  "accounts",
//	Usage: "manage bridge keystore",
//	Description: "The accounts command is used to manage the bridge keystore.\n" +
//		"\tTo generate a new account (key type generated is determined on the flag passed in): watcher accounts generate\n" +
//		"\tTo import a keystore file: watcher accounts import path/to/file\n" +
//		"\tTo import a geth keystore file: watcher accounts import --ethereum path/to/file\n" +
//		"\tTo import a private key file: watcher accounts import --privateKey private_key\n" +
//		"\tTo list keys: watcher accounts list",
//	Subcommands: []*cli.Command{
//		{
//			Action: wrapHandler(handleGenerateCmd),
//			Name:   "generate",
//			Usage:  "generate bridge keystore, key type determined by flag",
//			Flags:  generateFlags,
//			Description: "The generate subcommand is used to generate the bridge keystore.\n" +
//				"\tIf no options are specified, a secp256k1 key will be made.",
//		},
//		{
//			Action: wrapHandler(handleImportCmd),
//			Name:   "import",
//			Usage:  "import bridge keystore",
//			Flags:  importFlags,
//			Description: "The import subcommand is used to import a keystore for the bridge.\n" +
//				"\tA path to the keystore must be provided\n" +
//				"\tUse --ethereum to import an ethereum keystore from external sources such as geth\n" +
//				"\tUse --privateKey to create a keystore from a provided private key.",
//		},
//		{
//			Action:      wrapHandler(handleListCmd),
//			Name:        "list",
//			Usage:       "list bridge keystore",
//			Description: "The list subcommand is used to list all of the bridge keystores.\n",
//		},
//	},
//}
//
//var (
//	Version = "0.0.1"
//)
//
//// init initializes CLI
//func init() {
//	app.Action = run
//	app.Copyright = "Copyright 2019 Watcher Systems Authors"
//	app.Name = "watcher"
//	app.Usage = "Watcher"
//	app.Authors = []*cli.Author{{Name: "Watcher Systems 2019"}}
//	app.Version = Version
//	app.EnableBashCompletion = true
//	app.Commands = []*cli.Command{
//		&accountCommand,
//	}
//
//	app.Flags = append(app.Flags, cliFlags...)
//	app.Flags = append(app.Flags, devFlags...)
//}
//
//func main() {
//	if err := app.Run(os.Args); err != nil {
//		log.Error(err.Error())
//		os.Exit(1)
//	}
//}
//
//func startLogger(ctx *cli.Context) error {
//	logger := log.Root()
//	handler := logger.GetHandler()
//	var lvl log.Lvl
//
//	if lvlToInt, err := strconv.Atoi(ctx.String(config.VerbosityFlag.Name)); err == nil {
//		lvl = log.Lvl(lvlToInt)
//	} else if lvl, err = log.LvlFromString(ctx.String(config.VerbosityFlag.Name)); err != nil {
//		return err
//	}
//	log.Root().SetHandler(log.LvlFilterHandler(lvl, handler))
//
//	return nil
//}
//
//func run(ctx *cli.Context) error {
//	err := startLogger(ctx)
//	if err != nil {
//		return err
//	}
//
//	log.Info("Starting ChainBridge...")
//
//	cfg, err := config.GetConfig(ctx)
//	if err != nil {
//		return err
//	}
//
//	log.Debug("Config on initialization...", "config", *cfg)
//
//	// Check for test key flag
//	var ks string
//	var insecure bool
//	if key := ctx.String(config.TestKeyFlag.Name); key != "" {
//		ks = key
//		insecure = true
//	} else {
//		ks = cfg.KeystorePath
//	}
//
//	// Used to signal core shutdown due to fatal error
//	sysErr := make(chan error)
//	c := core.NewCore(sysErr)
//
//	for _, chain := range cfg.Chains {
//		chainId, errr := strconv.Atoi(chain.Id)
//		if errr != nil {
//			return errr
//		}
//		chainConfig := &core.ChainConfig{
//			Name:           chain.Name,
//			Id:             msg.ChainId(chainId),
//			URL:       chain.URL,
//			From:           chain.From,
//			KeystorePath:   ks,
//			Insecure:       insecure,
//			BlockstorePath: ctx.String(config.BlockstorePathFlag.Name),
//			FreshStart:     ctx.Bool(config.FreshStartFlag.Name),
//			LatestBlock:    ctx.Bool(config.LatestBlockFlag.Name),
//			Opts:           chain.Opts,
//		}
//		var newChain core.Chain
//		var m *metrics.ChainMetrics
//
//		logger := log.Root().New("chain", chainConfig.Name)
//
//		if ctx.Bool(config.MetricsFlag.Name) {
//			m = metrics.NewChainMetrics(chain.Name)
//		}
//
//		if chain.Type == "ethereum" {
//			newChain, err = ethereum.InitializeChain(chainConfig, logger, sysErr, m)
//		} else if chain.Type == "substrate" {
//			newChain, err = substrate.InitializeChain(chainConfig, logger, sysErr, m)
//		} else {
//			return errors.New("unrecognized Chain Type")
//		}
//
//		if err != nil {
//			return err
//		}
//		c.AddChain(newChain)
//
//	}
//
//	// Start prometheus and health server
//	if ctx.Bool(config.MetricsFlag.Name) {
//		port := ctx.Int(config.MetricsPort.Name)
//		blockTimeoutStr := os.Getenv(config.HealthBlockTimeout)
//		blockTimeout := config.DefaultBlockTimeout
//		if blockTimeoutStr != "" {
//			blockTimeout, err = strconv.ParseInt(blockTimeoutStr, 10, 0)
//			if err != nil {
//				return err
//			}
//		}
//		h := health.NewHealthServer(port, c.Registry, int(blockTimeout))
//
//		go func() {
//			http.Handle("/metrics", promhttp.Handler())
//			http.HandleFunc("/health", h.HealthStatus)
//			err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
//			if errors.Is(err, http.ErrServerClosed) {
//				log.Info("Health status server is shutting down", err)
//			} else {
//				log.Error("Error serving metrics", "err", err)
//			}
//		}()
//	}
//
//	c.Start()
//
//	return nil
//}

var ethereumConnection = ethereum.Connection{
	URL:  "http://127.0.0.1:8545",
	Http: true,
}

var substrateConnection = substrate.Connection{
	URL: "ws://127.0.0.1:9944",
	Key: &signature.TestKeyringPairAlice,
}

func InitializeChain() (*ethereum.Listener, error) {
	stop := make(chan struct{}, 1)
	ethconn := ethereum.NewConnection(ethereumConnection.URL, ethereumConnection.Http, stop)
	if err := ethconn.Connect(); err != nil {
		return nil, err
	}

	subconn := substrate.NewConnection(substrateConnection.URL, substrateConnection.Key, stop)
	if err := subconn.Connect(); err != nil {
		return nil, err
	}

	return &ethereum.Listener{
		Ethconn: ethconn,
		Subconn: subconn,
		Stop:    stop,
	}, nil
}

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogger)
}

func main() {
	listener, err := InitializeChain()
	if err != nil {
		log.Crit("failed to initialize chain", "error", err)
	}

	if err := listener.Subconn.SubmitTx(substrate.RegisterWatcher); err != nil {
		log.Crit("failed to register watcher", "error", err)
	}

	go func() {
		if err := listener.PollBlocks(); err != nil {
			log.Error("polling blocks failed", "error", err)
			return
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-sigs:
		log.Info("received the exit signal, ready to exit...")
	}

	listener.Ethconn.Close()
}
