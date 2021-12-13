package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/NuLink-network/watcher/watcher/chains/substrate"
	"github.com/NuLink-network/watcher/watcher/params"
)

type listener struct {
	ethconn Connection
	subconn substrate.Connection
	stop    chan struct{}
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `StartBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	var (
		currentBlock = params.StartBlock
		retry        = params.BlockRetryLimit
	)

	log.Info("Polling Blocks...", "block", currentBlock)

	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				log.Error("Polling failed, retries exceeded")
				return nil
			}

			latestBlock, err := l.ethconn.LatestBlock()
			if err != nil {
				log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(params.BlockRetryInterval)
				continue
			}

			// Sleep if currentBlock is greater than latestBlock; currentBlock > latestBlock
			if currentBlock.Cmp(latestBlock) == 1 {
				time.Sleep(params.BlockRetryInterval)
				continue
			}

			// Sleep if the difference is less than BlockConfirmations; (latestBlock - currentBlock) < BlockConfirmations
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(params.BlockConfirmations) == -1 {
				log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(params.BlockRetryInterval)
				continue
			}

			// Parse out events
			err = l.getDepositEventsForBlock(currentBlock)
			if err != nil {
				log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = params.BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	log.Debug("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(params.DepositContractAddress, Deposited, latestBlock, latestBlock)

	// querying for logs
	logs, err := l.ethconn.client.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, lg := range logs {
		// 1. get data from Topics and Data
		staker := lg.Topics[0]
		value := ethcommon.BytesToHash(lg.Data[:32]).Big()
		periods := ethcommon.BytesToHash(lg.Data[32:]).Big()

		// 2. send tx to substrate
		// todo
		if err := l.subconn.SubmitTx(substrate.BaseMethod, staker, value, periods); err != nil {
			log.Error("failed to send tx to substrate", "staker", staker, "value", value, "periods", periods, "err", err)
		}
	}

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		Topics: [][]ethcommon.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}
