# NuLink-watcher

<p align="center">
  <a href="https://www.nulink.org/"><img src="https://github.com/NuLink-network/nulink/blob/94c5538a5fdc25e7d4391f4f2e4af60b3c480fc1/logo/nulink-bg-1.png" width=40%  /></a>
</p>

<p align="center">
  <a href="https://github.com/NuLink-network"><img src="https://img.shields.io/badge/Playground-NuLink_Network-brightgreen?logo=Parity%20Substrate" /></a>
  <a href="http://nulink.org/"><img src="https://img.shields.io/badge/made%20by-NuLink%20Foundation-blue.svg?style=flat-square" /></a>
  <a href="https://github.com/NuLink-network/nulink-watcher"><img src="https://img.shields.io/badge/project-Nulink_Watcher-yellow.svg?style=flat-square" /></a>
</p>
## Introduction
This is the Repo for NuLink watcher nodes. The watcher nodes would use the ChainBridge(an ETH-Polkadot bridge) to retrieve the information of stakers and workers from the Nucypher contract in Ethereum.  

## Work flow
Once the watcher nodes started, it would first try to register in [NuProxy pallet](https://github.com/NuLink-network/nulink-chain/blob/main/pallets/nuproxy/src/lib.rs). It will fall into the following three situation:
1. If the node is illegal for registration, it will report a warning and exit.
2. If the node is legal for registration and already registered,  it will jump to the next step. 
3. Otherwise it would successfully register into the NuProxy pallet.

After the initial step, the watcher nodes would monitor the Ethereum network. If the staker information stored in NuCypher contract changes,  it would synchronize the updated status by sending the [updating extrinsic](https://github.com/NuLink-network/nulink-chain/blob/main/pallets/nuproxy/src/lib.rs#L169) in Polkadto Parachain periodically. 

## Getting Started
### Edit configuration file 
You can edit the default configuration file [here ](https://github.com/NuLink-network/nulink-watcher/blob/main/config.json) accordingly.

The structure of the configuration file is as:
```json5
{
  // stake info sync frequency, 100 means sync every 100 blocks
  "epochSize": 100,
  "ethereumConfig": {
    // the url of the ethereum RPC node
    "url": "https://mainnet.infura.io/v3/your_project_id",
    // whether the url of the ethereum RPC node is http protocol, 
    // currently only the http protocol is supported, so this parameter must be true
    "http": true,
    // the address of the nucypher deposit contract
    "depositContractAddr": "0xbbD3C0C794F40c4f993B03F65343aCC6fcfCb2e2"
  },
  "nuLinkChainConfig": {
    // the url of the NuLink RPC node
    "url": "ws://127.0.0.1:9944"
  }
}
```

### Build
```shell
cd cmd/watcher
go build -o watcher main.go
```

## Run
Going watcher -h can get help infos.

### Run the watcher with the specified parameters
```shell
./watcher --config ../../config.json  --mock
```

### Command parameters
You can use the default configuration or specify related configurations. The parameters you can specify are mainly the following.

`config`: This flag can be used to specify a json configuration file.

`mock`: Start the project in mock mode.
