package quiz

import (
	"log"
	"math/big"

	"github.com/PreciumFoundation/7renders-erc-api/tx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	colorLog "github.com/fatih/color"
)

type Wallet struct {
	abi                  *abi.ABI
	gateway              *tx.EtherGateway
	ControllerPrivateKey string
}

//NewWallet : Make a new Wallet's instance
func NewWallet() *Wallet {
	g := tx.NewEtherGateway()
	v, curEnvs, err := loadSubConfig(SubConfigWalletContract)
	if err != nil {
		log.Fatal(err)
	}
	contract := tx.LoadContract(tx.ContractTypeQuizWallet, curEnvs)
	return &Wallet{
		abi:     &contract,
		gateway: g,
		// Some source code for environment variables.
		// Caused by Precium x 7renders security problems, deleted.
	}
}

//Sweep : To be move PCM from user wallet created by controller.
func (w *Wallet) Sweep(walletAddress string, amount float64) (string, error) {
	v := ConvFloatEtherToWei(amount)
	data, err := w.abi.Pack("sweep", v)
	if err != nil {
		log.Fatal(err)
		colorLog.Red(err.Error())
		return "", err
	}
	log.Println("Data for transfer: ", data)
	txHash, err := w.gateway.Transaction(w.ControllerPrivateKey, 0, walletAddress, data)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	done, err := w.gateway.GetReceiptByTxHash(txHash)
	receipt := <-done
	log.Println("txHash:", txHash)
	log.Printf("%+v\n", receipt)
	colorLog.Green("%+v\n", receipt)
	colorLog.Cyan("End of Sweep PCM!!")
	return txHash, nil
}

//SweepedQty : Confirm sweeped amount from target wallet account.
func (w *Wallet) SweepedQty(walletAddress string) (string, error) {

	data, err := w.abi.Pack("SweepedQty")
	c := tx.CallParams{
		To:   walletAddress,
		Data: "0x" + common.Bytes2Hex(data),
	}
	//cList := make([]tx.CallParams, 1)
	cList := make([]interface{}, 2)
	cList[0] = c
	cList[1] = "latest"
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	packedData, err := w.gateway.Call(walletAddress, 0, walletAddress, cList)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	log.Println(string(packedData))
	var v *big.Int
	err = w.abi.Unpack(&v, "SweepedQty", packedData)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println(string(packedData))
	log.Println("SweepedQty is :", v.String())
	return v.String(), nil
}
