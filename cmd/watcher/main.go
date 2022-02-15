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
	"strconv"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"

	"github.com/NuLink-network/watcher/watcher/chains/ethereum"
	"github.com/NuLink-network/watcher/watcher/chains/substrate"
	"github.com/NuLink-network/watcher/watcher/config"
	"github.com/NuLink-network/watcher/watcher/params"
)

var (
	app     = cli.NewApp()
	Version = "0.1.0"
)

var cliFlags = []cli.Flag{
	config.VerbosityFlag,
	config.ConfigFileFlag,
	config.StakeInfoFileFlag,
}

func init() {
	app.Action = run
	app.Version = Version
	app.Copyright = "Copyright 2021 Watcher Systems Authors"
	app.Name = "watcher"

	app.Flags = append(app.Flags, cliFlags...)

	//app.Before = func(ctx *cli.Context) error {
	//	return setup(ctx)
	//}
	//app.After = func(ctx *cli.Context) error {
	//	return exit(ctx)
	//}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Error("run failed", "error", err.Error())
		os.Exit(1)
	}
}

func InitializeChain(cfg *config.Config) (*ethereum.Listener, error) {
	stop := make(chan struct{}, 1)
	ethconn := ethereum.NewConnection(cfg.EthereumConfig.URL, cfg.EthereumConfig.Http, stop)
	if err := ethconn.Connect(); err != nil {
		return nil, err
	}

	//kp, err := signature.KeyringPairFromSecret(cfg.NuLinkChainConfig.Seed, cfg.NuLinkChainConfig.Network)
	//if err != nil {
	//	return nil, err
	//}

	subconn := substrate.NewConnection(cfg.NuLinkChainConfig.URL, params.Watcher, stop)
	if err := subconn.Connect(); err != nil {
		return nil, err
	}

	return &ethereum.Listener{
		Config:  cfg.EthereumConfig,
		Ethconn: ethconn,
		Subconn: subconn,
		Stop:    stop,
	}, nil
}

var listener *ethereum.Listener

func run(ctx *cli.Context) error {
	if err := setup(ctx); err != nil {
		return err
	}

	cfg, err := config.GetConfig(ctx)
	if err != nil {
		return err
	}

	//lp := ctx.String(config.BlockStoreFileFlag.Name)
	//number, err := ethereum.ReadLatestBlock(lp)
	//if err != nil {
	//	return err
	//}
	//if number.Cmp(cfg.EthereumConfig.StartBlock) == 1 {
	//	cfg.EthereumConfig.StartBlock = number
	//}

	listener, err = InitializeChain(cfg)
	if err != nil {
		log.Error("failed to initialize chain", "error", err)
		return err
	}
	//listener.LatestBlockPath = lp
	listener.LastStakeInfoPath = ctx.String(config.StakeInfoFileFlag.Name)

	if err := listener.Subconn.SubmitTx(substrate.RegisterWatcher); err != nil {
		log.Error("failed to register watcher", "error", err)
		return err
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
	case <-listener.Stop:
		log.Info("listener stop...")
	case <-sigs:
		log.Info("received the exit signal, ready to exit...")
	}

	_ = exit(ctx)
	return nil
}

func setup(ctx *cli.Context) error {
	if err := startLogger(ctx); err != nil {
		return err
	}
	log.Info("setup watcher...")
	//return ethereum.ReadStakeInfoFromFile(ctx.String(config.StakeInfoFileFlag.Name))
	return nil
}

func exit(ctx *cli.Context) error {
	log.Info("exit watcher...")
	listener.Ethconn.Close()
	//return ethereum.WriteStakeInfoToFile(ctx.String(config.StakeInfoFileFlag.Name))
	return nil
}

func startLogger(ctx *cli.Context) error {
	var lvl log.Lvl
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))

	if lvlToInt, err := strconv.Atoi(ctx.String(config.VerbosityFlag.Name)); err == nil {
		lvl = log.Lvl(lvlToInt)
	} else if lvl, err = log.LvlFromString(ctx.String(config.VerbosityFlag.Name)); err != nil {
		return err
	}
	glogger.Verbosity(lvl)
	log.Root().SetHandler(glogger)

	return nil
}
