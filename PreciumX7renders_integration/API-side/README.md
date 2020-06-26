# 7renders-erc-api
API server using infura.io's api

### Managing env variables using go-viper
- [go-viper](https://github.com/spf13/viper)
* viper could be injected parameters as run-time
* Based on YAML format environment variables in this project

### When you want to check transaction contents, please click below links and type txhash...(etc)
- [using ropsten](https://ropsten.etherscan.io/) 
- [using main-net](https://etherscan.io/)

###  using infura.io's nodes. and we created an account by "ino@k-cent.com" 
#### below link redirect to infura.io's page.
- [infura Dashboard](https://infura.io/stats/91783f0ea6314a0fb122ed25e1fbd0ff)
- [infura APIs](https://infura.io/docs/ethereum/json-rpc/)

##### getting start JSON-RPC 
- [Ethereum's JSON-RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC)
| A described page for ethereum JSON-RPC

### if you have any problems. please feel free to contact with the person in charge.
#### email : ino@k-cent.com / ino@precium.io 
----


## **Remeber**
### Have to make `envs.yaml` file before executing 
### Ethereum library needs gcc.
- [tdm64-gcc for windows](http://tdm-gcc.tdragon.net/download)


## Descript for envs.yaml's parameter 
#### `ControllerKey` : Pre-authorized wallet. you can be used when you want to more user wallet(=user contract).
#### `CommandContract` : Control tower for quiz system. This contract has to contain the function `makeWallet` and so on.
