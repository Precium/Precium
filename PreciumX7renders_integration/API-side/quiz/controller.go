package quiz

import (
	"log"

	"github.com/PreciumFoundation/7renders-erc-api/tx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Controller : Quiz system's manager.
type Controller struct {
	abi                    *abi.ABI
	gateway                *tx.EtherGateway
	CommandContractAddress string
	ControllerPrivateKey   string
}

//NewController : Make a controller's instance
func NewController() *Controller {

	v, curEnvs, err := loadSubConfig(SubConfigController)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	contract := tx.LoadContract(tx.ContractTypeController, curEnvs)
	g := tx.NewEtherGateway()
	c := &Controller{
		abi:     &contract,
		gateway: g,
		// Some source code for environment variables.
		// Caused by Precium x 7renders security problems, deleted.
	}
	if c.CommandContractAddress == "" {
		log.Fatal("There is no config value.", c)
		return nil
	}
	return c
}

//GenerateWallet : Make a new contract(user wallet). This function returns txhash, new address(chan), error.
func (c *Controller) GenerateWallet() (string, string, error) {
	data, err := c.abi.Pack("makeWallet")
	if err != nil {
		log.Fatal("makewallet data :", data)
		return "", "", err
	}

	txHash, err := c.gateway.Transaction(c.ControllerPrivateKey, 0, c.CommandContractAddress, data)
	if err != nil {
		log.Fatal("Error occurred", err)
		return "", "", err
	}
	log.Println("Success to generate user's wallet, txhash :", txHash)
	// Get an event log(included contract address), after the transaction has done.

	chanReceipt, err := c.gateway.GetReceiptByTxHash(txHash)

	r := <-chanReceipt
	hex := common.Bytes2Hex(r.Logs[0].Data)
	address := common.HexToAddress(hex)
	strAddress := address.String()

	return txHash, strAddress, nil
}
