package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
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

	"github.com/NuLink-network/watcher/watcher/bindings/nucypher"
	"github.com/NuLink-network/watcher/watcher/chains/substrate"
	"github.com/NuLink-network/watcher/watcher/config"
	"github.com/NuLink-network/watcher/watcher/params"
)

var first = true
var accountID types.AccountID
var stakeInfoList = make(substrate.StakeInfos, 0)

type Listener struct {
	Config  config.EthereumConfig
	Ethconn *Connection
	Subconn *substrate.Connection
	//LatestBlockPath   string
	LastStakeInfoPath string
	Stop              chan struct{}
}

func init() {
	bs, _ := hex.DecodeString("d0e38319ea54a336f829da6cb91e11806c3474c354a5c05476b76ea97d6db414")
	accountID = types.NewAccountID(bs)
}

// PollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `StartBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *Listener) PollBlocks() error {
	var (
		currentBlock = big.NewInt(1)
		retry        = params.BlockRetryLimit
	)

	log.Info("Polling Blocks...")

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
			if latestBlock.Cmp(currentBlock) != 1 {
				log.Debug("Block not ready, will retry", "target", latestBlock.Uint64()+1, "latest", latestBlock)
				time.Sleep(params.BlockRetryInterval)
				continue
			}
			log.Info("get latest block", "block", latestBlock)

			err = l.syncStakeInfos(latestBlock)
			if err != nil {
				l.Stop <- struct{}{}
				return err
			}

			//if err := WriteLatestBlock(l.LatestBlockPath, latestBlock); err != nil {
			//	log.Error("Failed to write latest block", "block", latestBlock, "err", err)
			//}

			// Goto next block and reset retry counter
			currentBlock = latestBlock
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

func (l *Listener) syncStakeInfos(latestBlock *big.Int) error {
	if first || latestBlock.Uint64()%uint64(params.EpochLength) == 0 {
		first = false
		log.Info("ready to update stake info to nulink", "block", latestBlock)

		stakeInfos, err := l.GetStakeInfo()
		if err != nil {
			return err
		}

		lastInfos, err := ReadStakeInfos(l.LastStakeInfoPath)
		if err != nil {
			return err
		}
		top20StakeInfos := AssignCoinbase(stakeInfos.LockedBalanceTop20(), lastInfos)
		if err := l.Subconn.SubmitTx(substrate.UpdateStakeInfo, top20StakeInfos); err != nil {
			log.Error("failed to update stake info to nulink", "count", len(top20StakeInfos), "error", err)
			return err
		}
		log.Info("succeeded to update stake info to nulink", "count", len(top20StakeInfos))

		if err := WriteStakeInfos(l.LastStakeInfoPath, top20StakeInfos); err != nil {
			return err
		}
	} else if latestBlock.Uint64()%10 == 0 {
		if err := l.Subconn.SubmitTx(substrate.UpdateStakeInfo, substrate.StakeInfos{}); err != nil {
			log.Error("failed to update empty stake info to nulink", "count", 0, "error", err)
			return err
		}
		log.Info("succeeded to update empty stake info to nulink", "count", 0)
	}
	return nil
}

func AssignCoinbase(top20StakeInfos substrate.StakeInfos, lastInfos map[string][32]byte) substrate.StakeInfos {
	newStakeIndex := make([]int, 0)
	accounts := make(map[types.AccountID]struct{}, len(params.AccountIDs))
	for k, v := range params.AccountIDs {
		accounts[k] = v
	}

	for i, info := range top20StakeInfos {
		cb, ok := lastInfos[ethcommon.Bytes2Hex(info.WorkBase)]
		if ok {
			top20StakeInfos[i].Coinbase = cb
			delete(accounts, cb)
			continue
		}
		newStakeIndex = append(newStakeIndex, i)
	}

	as := make([]types.AccountID, 0, len(accounts))
	for a := range accounts {
		as = append(as, a)
	}
	for i, s := range newStakeIndex {
		top20StakeInfos[s].Coinbase = as[i]
	}
	return top20StakeInfos
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

func (l *Listener) GetStakeInfo() (substrate.StakeInfos, error) {
	stakeInfos := make(substrate.StakeInfos, 0)
	nc, err := nucypher.NewNucypher(ethcommon.HexToAddress(l.Config.DepositContractAddr), l.Ethconn.Client)
	if err != nil {
		log.Error("failed to new nucypher", "error", err)
		return stakeInfos, nil
	}
	length, err := nc.GetStakersLength(nil)
	if err != nil {
		log.Error("failed to get stakes length", "error", err)
		return stakeInfos, nil
	}
	log.Info("succeeded to get stakes length", "length", length.Uint64())

	for i := int64(0); i < length.Int64(); i++ {
		staker, err := nc.Stakers(nil, big.NewInt(i))
		if err != nil {
			log.Error("failed to get stakes", "index", i, "error", err)
			continue
		}

		info, err := nc.StakerInfo(nil, staker)
		if err != nil {
			log.Error("failed to get stake info", "staker", staker, "error", err)
			continue
		}

		stakeInfos = append(stakeInfos, &substrate.StakeInfo{
			Coinbase:      types.NewAccountID(staker[:]),
			WorkBase:      staker[:],
			IsWork:        true,
			LockedBalance: types.NewU128(*info.Value),
			WorkCount:     0,
		})
		log.Debug("succeeded to import stake info", "staker", staker)
	}
	return stakeInfos, err
}

func ReadStakeInfos(file string) (map[string][32]byte, error) {
	//stakeInfoList := make(substrate.StakeInfos, 0)
	stakeInfoList := make(map[string][32]byte, 0)
	// If it exists, load and return
	exists, err := fileExists(file)
	if err != nil {
		return stakeInfoList, err
	}
	if !exists {
		log.Warn("stake info file does not exist")
		return stakeInfoList, nil
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("read stake info list from file filed", "error", err)
		return stakeInfoList, err
	}
	if len(data) == 0 {
		log.Warn("stake info file is empty")
		return stakeInfoList, nil
	}

	//var infos []*substrate.StakeInfo
	var infos map[string][32]byte
	if err := json.Unmarshal(data, &infos); err != nil {
		log.Error("json unmarshal stake info list failed", "error", err)
		return stakeInfoList, err
	}

	stakeInfoList = infos

	return stakeInfoList, err
}

func WriteStakeInfos(file string, infos substrate.StakeInfos) error {
	// Create dir if it does not exist
	if _, err := os.Stat(file); os.IsNotExist(err) {
		dir, _ := filepath.Split(file)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	stakeInfos := make(map[string][32]byte)
	for _, info := range infos {
		stakeInfos[ethcommon.Bytes2Hex(info.WorkBase)] = info.Coinbase
	}

	data, err := json.Marshal(stakeInfos)
	if err != nil {
		log.Error("json marshal stake info list failed", "error", err)
		return err
	}
	if err := ioutil.WriteFile(file, data, 0664); err != nil {
		log.Error("write stake info list to file filed", "error", err)
		return err
	}
	log.Info("write stake info list to file succeeded", "count", len(stakeInfos))
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
