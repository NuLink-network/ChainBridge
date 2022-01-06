package ethereum

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/NuLink-network/watcher/watcher/chains/substrate"
	"github.com/NuLink-network/watcher/watcher/config"
	"github.com/NuLink-network/watcher/watcher/params"
)

var stakeInfoList = make(substrate.StakeInfos, 0)

type Listener struct {
	Config          config.EthereumConfig
	Ethconn         *Connection
	Subconn         *substrate.Connection
	LatestBlockPath string
	Stop            chan struct{}
}

// PollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `StartBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *Listener) PollBlocks() error {
	var (
		currentBlock = l.Config.StartBlock
		retry        = params.BlockRetryLimit
	)

	log.Info("Polling Blocks...", "block", currentBlock)

	for {
		select {
		case <-l.Stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				log.Error("Polling failed, retries exceeded")
				l.Stop <- struct{}{}
				return nil
				// Goto next block and reset retry counter
				//currentBlock.Add(currentBlock, big.NewInt(1))
				//retry = params.BlockRetryLimit
				//continue
			}

			latestBlock, err := l.Ethconn.LatestBlock()
			if err != nil {
				log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(params.BlockRetryInterval)
				continue
			}

			// Sleep if the difference is less than BlockConfirmations; (latestBlock - currentBlock) < BlockConfirmations
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.Config.BlockConfirmations) == -1 {
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
			if err := WriteLatestBlock(l.LatestBlockPath, latestBlock); err != nil {
				log.Error("Failed to write latest block", "block", latestBlock, "err", err)
			}

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = params.BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *Listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	log.Info("Querying block for deposit events", "block", latestBlock)
	query := buildQuery(ethcommon.HexToAddress(l.Config.DepositContractAddr), Deposited, latestBlock, latestBlock)

	// querying for logs
	logs, err := l.Ethconn.Client.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, lg := range logs {
		// 1. get data from Topics and Data
		staker := lg.Topics[1]
		value := ethcommon.BytesToHash(lg.Data[:32]).Big()
		periods := ethcommon.BytesToHash(lg.Data[32:]).Big()

		stakeInfoList = append(stakeInfoList, &substrate.StakeInfo{
			Coinbase:      types.NewAccountID(staker[:]),
			WorkBase:      staker[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*value),
			WorkCount:     0,
		})
		log.Info("find deposit event", "staker", staker, "value", value, "periods", periods)
	}
	if latestBlock.Uint64()%uint64(params.EpochLength) == 0 {
		if len(stakeInfoList) == 0 {
			return nil
		}

		if err := l.Subconn.SubmitTx(substrate.UpdateStakeInfo, stakeInfoList.LockedBalanceTop20()); err != nil {
			log.Error("failed to update stake info to nulink", "count", len(stakeInfoList), "error", err)
		} else {
			log.Error("succeeded to update stake info to nulink", "count", len(stakeInfoList))
		}

		stakeInfoList = make([]*substrate.StakeInfo, 0, 1000)
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

func fileExists(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func ReadStakeInfoFromFile(file string) error {
	// If it exists, load and return
	exists, err := fileExists(file)
	if err != nil {
		return err
	}
	if !exists {
		log.Warn("stake info file does not exist")
		return nil
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("read stake info list from file filed", "error", err)
		return err
	}
	if len(data) == 0 {
		log.Warn("stake info file is empty")
		return nil
	}

	var infos []*substrate.StakeInfo
	if err := rlp.DecodeBytes(data, &infos); err != nil {
		log.Error("rlp decode stake info list failed", "error", err)
		return err
	}

	if stakeInfoList == nil {
		stakeInfoList = make([]*substrate.StakeInfo, 0)
	}
	stakeInfoList = infos

	return nil
}

func WriteStakeInfoToFile(file string) error {
	// Create dir if it does not exist
	if _, err := os.Stat(file); os.IsNotExist(err) {
		dir, _ := filepath.Split(file)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	data, err := rlp.EncodeToBytes(stakeInfoList)
	if err != nil {
		log.Error("json marshal stake info list failed", "error", err)
		return err
	}
	if err := ioutil.WriteFile(file, data, 0664); err != nil {
		log.Error("write stake info list to file filed", "error", err)
		return err
	}
	log.Info("write stake info list to file success", "count", len(stakeInfoList))
	return nil
}

func WriteLatestBlock(file string, number *big.Int) error {
	// Create dir if it does not exist
	if _, err := os.Stat(file); os.IsNotExist(err) {
		dir, _ := filepath.Split(file)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Write bytes to file
	data := []byte(number.String())
	err := ioutil.WriteFile(file, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
func ReadLatestBlock(file string) (*big.Int, error) {
	// If it exists, load and return
	exists, err := fileExists(file)
	if err != nil {
		return nil, err
	}
	if exists {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		block, _ := new(big.Int).SetString(string(data), 10)
		return block, nil
	}
	// Otherwise just return 0
	return big.NewInt(0), nil
}
