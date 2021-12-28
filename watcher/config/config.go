// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

const configFile = "/config.json"

func DefaultConfigFile() string {
	return DefaultDir() + configFile
}

type Config struct {
}

func (c *Config) validate() error {
	// todo
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
		log.Warn("err loading json file", "err", err.Error())
		return &cfg, err
	}
	//if ksPath := ctx.String(KeystorePathFlag.Name); ksPath != "" {
	//	fig.KeystorePath = ksPath
	//}

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
