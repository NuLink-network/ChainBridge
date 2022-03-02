// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

const configFile = "/config.json"

func DefaultConfigFile() string {
	return DefaultDir() + configFile
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) <= 0
}

type Config struct {
	EpochSize         uint64            `json:"epochSize"`
	EthereumConfig    EthereumConfig    `json:"ethereumConfig"`
	NuLinkChainConfig NuLinkChainConfig `json:"nuLinkChainConfig"`
}

type EthereumConfig struct {
	URL                 string `json:"url"`
	Http                bool   `json:"http"`
	DepositContractAddr string `json:"depositContractAddr"`
	//StartBlock          *big.Int `json:"startBlock"`
	//BlockConfirmations  *big.Int `json:"blockConfirmations"`
}

type NuLinkChainConfig struct {
	URL string `json:"url"`
	//Seed    string `json:"seed"`
	//Network uint8  `json:"network"`
}

func (c *Config) validate() error {
	if c.EpochSize == 0 {
		c.EpochSize = EpochSize
	}
	if IsEmpty(c.EthereumConfig.URL) {
		return fmt.Errorf("required field URL for ethereum")
	}
	if IsEmpty(c.EthereumConfig.DepositContractAddr) {
		return fmt.Errorf("required field DepositContractAddr for ethereum")
	}
	if IsEmpty(c.NuLinkChainConfig.URL) {
		return fmt.Errorf("required field URL for nuLinkChain")
	}
	//if IsEmpty(c.NuLinkChainConfig.Seed) {
	//	return fmt.Errorf("required field Seed for substrate")
	//}
	return nil
}

func GetConfig(ctx *cli.Context) (*Config, error) {
	var cfg Config
	path := DefaultConfigFile()
	if file := ctx.String(ConfigFileFlag.Name); file != "" {
		path = file
	}
	err := loadConfig(path, &cfg)
	if err != nil {
		log.Warn("failed to loading json file", "err", err.Error())
		return &cfg, err
	}

	log.Debug("Loaded config", "path", path)
	err = cfg.validate()
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(file string, config *Config) error {
	ext := filepath.Ext(file)
	fp, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	log.Debug("Loading configuration", "path", filepath.Clean(fp))

	f, err := os.Open(filepath.Clean(fp))
	if err != nil {
		return err
	}

	if ext == ".json" {
		if err = json.NewDecoder(f).Decode(&config); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unrecognized extention: %s", ext)
	}

	return nil
}
