// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"context"
	"math/big"
	"testing"

	"github.com/ChainSafe/chainbridge-utils/crypto/secp256k1"

	utils "github.com/NuLink-network/watcher/shared/ethereum"
)

func NewClient(t *testing.T, endpoint string, kp *secp256k1.Keypair) *utils.Client {
	client, err := utils.NewClient(endpoint, kp)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func GetLatestBlock(t *testing.T, client *utils.Client) *big.Int {
	block, err := client.Client.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	return block.Number()
}

func LockNonceAndUpdate(t *testing.T, client *utils.Client) {
	err := client.LockNonceAndUpdate()
	if err != nil {
		t.Fatal(err)
	}
}
