# NuLink-watcher

<p align="center">
  <a href="https://www.nulink.org/"><img src="https://github.com/NuLink-network/nulink/blob/94c5538a5fdc25e7d4391f4f2e4af60b3c480fc1/logo/nulink-bg-1.png" width=40%  /></a>
</p>

<p align="center">
  <a href="https://github.com/NuLink-network"><img src="https://img.shields.io/badge/Playground-NuLink_Network-brightgreen?logo=Parity%20Substrate" /></a>
  <a href="http://nulink.org/"><img src="https://img.shields.io/badge/made%20by-NuLink%20Foundation-blue.svg?style=flat-square" /></a>
  <a href="https://github.com/NuLink-network/nulink-watcher"><img src="https://img.shields.io/badge/project-Nulink_Watcher-yellow.svg?style=flat-square" /></a>
</p>

This is the Repo for NuLink watcher nodes. The watcher nodes would use the ChainBridge(an ETH-Polkadot bridge) to retrieve the information of stakers and workers from the Nucypher contract in Ethereum. All active watcher nodes constuct the watcher network which provide the necessary data to the mirror pallet in Polkadot.

## Getting Started
### Edit configuration file 
The configuration file path can be specified on the command line or the default path can be used.

The structure of the configuration file is as:
```json
{
  "ethereumConfig": {
    "url": "https://mainnet.infura.io/v3/8cce6b470ad44fb5a3621aa34243647f",
    "http": true,
    "depositContractAddr": "0xbbD3C0C794F40c4f993B03F65343aCC6fcfCb2e2"
  },
  "platonConfig": {
    "url": "http://35.247.155.162:6789",
    "http": true,
    "privateKey": "814a3ba398642139a6436632f46921d9ab62e05c365598f4949651af40e537ac",
    "stakeContractAddr": "0x70F70012d1B9F36A3B8db77A58f62547FB1bA19a",
    "policyContractAddr": "0x4d2e3B25b2982F15022BE1a43757DBed3F529fe4"
  },
  "nuLinkChainConfig": {
    "url": "ws://127.0.0.1:9944"
  }
}
```
The default configuration file paths for different systems are as follows:

| OS| File Path |
| -------- | ------  | 
| darwin  | $HOME/Library/NuLinkWatcher/config.json | 
| windows | $HOME/AppData/Roaming/NuLinkWatcher/config.json  | 
| other   | $HOME/NuLinkWatcher/config.json | 

### Build
```shell
cd cmd/watcher
go build -o watcher main.go
```

## Run
Going watcher -h can get help infos.

### Run the watcher with default parameters
```shell
./watcher
```

### Run the watcher with the specified parameters
```shell
./watcher --verbosity debug --config $HOME/config.json --file $HOME/stake_info.json
```

### Command parameters
You can use the default configuration or specify related configurations. The parameters you can specify are mainly the following three.

`verbosity`: Specify the output level of the log, the default is the info level, the optional levels include trace, debug, info, warn, error, crit.

`config`: This flag can be used to specify a json configuration file.

`file`: This flag can be used to specify a last stake info file. The contents of this file do not require user editing. If no file is specified, the program will automatically create this file
