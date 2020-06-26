package tx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"time"

	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EtherGateway : Infura network gateway. Please see details `infura.io`
type EtherGateway struct {
	NetworkType     string
	ProjectID       string
	ProjectSecret   string
	EndPoint        string
	Host            string
	GasLimit        uint64
	CommandContract string
}

// Body :  Body for post request
type RequestBody struct {
	JSONRPC string   `json:"jsonrpc,omitempty"`
	Method  string   `json:"method,omitempty"`
	Params  []string `json:"params,omitempty"`
	ID      int      `json:"id,omitempty"`
}

const latestBlock string = "latest"

type RequestBodyWithObject struct {
	JSONRPC string        `json:"jsonrpc,omitempty"`
	Method  string        `json:"method,omitempty"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id,omitempty"`
}

type ResponseBody struct {
	JSONRPC string `json:"jsonrpc,omitempty"`
	ID      int    `json:"id,omitempty"`
	// Result  interface{} `json:"result,omitempty"`
	Result string `json:"result,omitempty"`
	Error  Error  `json:"error,omitempty"`
}
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
type CallParams struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Gas      string `json:"gas,omitempty"`
	GasPrice string `json:"gas_price,omitempty"`
	Value    string `json:"value,omitempty"`
	Data     string `json:"data,omitempty"`
}

const (
	MethodNetworkType         string = "net_version"
	MethodGasPrice            string = "eth_gasPrice"
	MethodGetBalance          string = "eth_getBalance"
	MethodGetTransactionCount string = "eth_getTransactionCount"
	MethodGetHashRate         string = "eth_hashrate"
	MethodSendRawTransaction  string = "eth_sendRawTransaction"
	MethodEthCall             string = "eth_call"
)

func getRequestBody(method string, params []string) *RequestBody {
	return &RequestBody{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
}

func getRequestBodyWithObject(method string, params []interface{}) *RequestBodyWithObject {
	return &RequestBodyWithObject{
		JSONRPC: "2.0",
		ID:      1,
		Method:  method,
		Params:  params,
	}
}
func (e *EtherGateway) loadConfigs() error {
	// Some source code for environment variables.
	// Caused by Precium x 7renders security problems, deleted.
	file, err := os.Open(basePath[:len(basePath)-4] + "envs.yaml")
	// Too many code was exsited here...😉
	if e.NetworkType == "" || e.ProjectID == "" || e.ProjectSecret == "" ||
		e.EndPoint == "" || e.Host == "" || e.GasLimit == 0 {
		return fmt.Errorf("Wrong config had set! \n%+v", e)
	}
	return nil
}

// NewEtherGateway : Create EtherGateway instance
func NewEtherGateway() *EtherGateway {
	e := EtherGateway{}
	if err := e.loadConfigs(); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return &e
}

// SendRequest : Request to infura node.
func (e *EtherGateway) SendRequest(rBody RequestBody) (ResponseBody, error) {
	requestBody, err := json.Marshal(rBody)
	log.Println(string(requestBody))
	if err != nil {

		log.Fatal(err)
		return ResponseBody{}, err
	}
	resp, err := http.Post(
		e.EndPoint,
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
		return ResponseBody{}, err
	}
	defer resp.Body.Close()

	var response ResponseBody
	json.NewDecoder(resp.Body).Decode(&response)
	return response, nil
}

// SendRequest : Request to infura node.
func (e *EtherGateway) SendRequestWithObject(rBody RequestBodyWithObject) (ResponseBody, error) {
	requestBody, err := json.Marshal(rBody)
	log.Println(string(requestBody))
	if err != nil {
		log.Fatal(err)
		return ResponseBody{}, err
	}
	resp, err := http.Post(
		e.EndPoint,
		"application/json",
		bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
		return ResponseBody{}, err
	}
	defer resp.Body.Close()

	var response ResponseBody
	json.NewDecoder(resp.Body).Decode(&response)
	return response, nil
}

//GetCurrentNetworkType : Function for recognize network type
func (e *EtherGateway) GetCurrentNetworkType() string {
	b := getRequestBody(MethodNetworkType, nil)
	respBody, err := e.SendRequest(*b)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return respBody.Result
}

//GetCurrentGasPrice : GetCurrentGasPrice
func (e *EtherGateway) GetCurrentGasPrice() uint64 {
	b := getRequestBody(MethodGasPrice, nil)
	respBody, err := e.SendRequest(*b)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return uint64(hex2int(respBody.Result))
}

func (e *EtherGateway) makeRawTransaction(fromPrivKey string, amount float64, to string, data []byte) (string, error) {
	//make client object
	client, err := ethclient.Dial(e.Host)
	if err != nil {
		log.Fatal(err)
	}

	//Unlock wallet
	privateKey, err := crypto.HexToECDSA(fromPrivKey)
	if err != nil {
		log.Fatal("Invalid wallet information(pKey)")
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Invalid wallet information(invalid ecdsa instance)")
		return "", fmt.Errorf("Invalid ecdsa key pair.")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(int64(amount * math.Pow10(18)))
	gasLimit := e.GasLimit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress(to)
	var additionalData []byte
	if data != nil {
		additionalData = data
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, additionalData)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signedTx}
	rawTx := hex.EncodeToString(ts.GetRlp(0))
	return fmt.Sprintf("0x%s", rawTx), nil
}

// txEther : Could be transferred only pure ethereum, excluded data area. for network testing.
func (e *EtherGateway) txEther(fromPrivKey string, amount float64, to string) (string, error) {
	rawTx, err := e.makeRawTransaction(fromPrivKey, amount, to, nil)
	if err != nil {
		log.Fatal(err)
	}
	params := make([]string, 1)
	params[0] = rawTx
	b := getRequestBody(MethodSendRawTransaction, params)
	respBody, err := e.SendRequest(*b)
	return respBody.Result, nil
}

// Transaction :  It could be transferred with additional data.
func (e *EtherGateway) Call(from string, amount float64, to string, data []interface{}) ([]byte, error) {

	b := getRequestBodyWithObject(MethodEthCall, data)

	respBody, err := e.SendRequestWithObject(*b)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if respBody.Error.Code != 0 {
		log.Fatal(respBody.Error)
		return nil, fmt.Errorf("%s", respBody.Error.Message)
	}
	v := common.Hex2Bytes(respBody.Result[2:])
	return v, nil
}

// Transaction :  It could be transferred with additional data.
func (e *EtherGateway) Transaction(fromPrivKey string, amount float64, to string, data []byte) (string, error) {
	rawTx, err := e.makeRawTransaction(fromPrivKey, amount, to, data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	params := make([]string, 1)
	params[0] = rawTx
	b := getRequestBody(MethodSendRawTransaction, params)
	respBody, err := e.SendRequest(*b)
	return respBody.Result, nil
}

// GetRecieptByTxHash : Return receipt searched by txHash(included prefix `0x`).
func (e *EtherGateway) GetReceiptByTxHash(txhash string) (<-chan *types.Receipt, error) {
	client, err := ethclient.Dial(e.Host)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan *types.Receipt)
	maxWaitingCount := 100
	isTimeout := false
	go func() {
		for count := 0; count < maxWaitingCount; count++ {
			if count+1 >= maxWaitingCount {
				isTimeout = true
			}
			if count%10 == 0 {
				log.Println("[", goid(), "]waiting count ...(", count, ")")
			}
			receipt, err := client.TransactionReceipt(context.Background(),
				common.HexToHash(txhash))
			if err != nil {
				sleepTime := time.Second * 5
				log.Printf("[", goid(), "]Waiting for transaction has over(during %dms)", sleepTime)
				time.Sleep(sleepTime)
				continue
			}
			done <- receipt
			break
		}
		if isTimeout {
			log.Println("[", goid(), "]Event log gather pool is over.(timeout)")
			done <- nil
		}
	}()
	return done, nil
}
func (e *EtherGateway) GetClient() *ethclient.Client {
	client, err := ethclient.Dial(e.Host)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
