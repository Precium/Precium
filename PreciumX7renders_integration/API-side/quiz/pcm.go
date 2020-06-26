package quiz

import (
	"log"
	"math/big"

	"github.com/PreciumFoundation/7renders-erc-api/tx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	colorLog "github.com/fatih/color"
)

// PCM :  It contains PCM abi data.
type PCM struct {
	abi             abi.ABI
	gateway         *tx.EtherGateway
	contractAddress string
}

func NewPCM() *PCM {
	g := tx.NewEtherGateway()
	v, curEnvs, err := loadSubConfig(SubconfigPrecium)
	if err != nil {
		log.Fatal("There is no config column(precium")
		return nil
	}
	//ContractData := tx.LoadContractByString(tx.ContractTypePrecium , curEnvs)
	contract := tx.LoadContract(tx.ContractTypePrecium, curEnvs)

	p := &PCM{
		abi:             contract,
		gateway:         g,
		contractAddress: v.GetString("contractAddress"),
	}
	return p
}

func (p *PCM) BalanceOf(OwnerHexAddress string) (string, error) {

	address := common.HexToAddress(OwnerHexAddress)
	data, err := p.abi.Pack("balanceOf", address)
	c := tx.CallParams{
		To:   p.contractAddress,
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

	packedData, err := p.gateway.Call(OwnerHexAddress, 0, p.contractAddress, cList)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	log.Println(string(packedData))
	var v *big.Int
	err = p.abi.Unpack(&v, "balanceOf", packedData)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Println(string(packedData))
	log.Println("left balance is :", v.String())
	return v.String(), nil
}

// TransferPCMFromOwner : Transfer PCM from onwer(for test),
func (p *PCM) TransferPCMFromOwner(amount float64, to string) {

	FromAccountKey := "574056cccd8c58222feed069d26a196039bd4b8857c76b7f0a2f30e3fa14855f"
	p.TransferPCM(FromAccountKey, amount, to)
}

// TransferPCM : TransferPCM.
func (p *PCM) TransferPCM(from string, amount float64, to string) {
	colorLog.Cyan("Start TranferPCM!!")

	v := ConvFloatEtherToWei(amount)

	data, err := p.abi.Pack("transfer", common.HexToAddress(to), v)
	if err != nil {
		log.Fatal(err)
		colorLog.Red(err.Error())
	}
	log.Println("Data for transfer: ", data)
	txHash, err := p.gateway.Transaction(from, 0, p.contractAddress, data)
	if err != nil {
		log.Fatal(err)
	}
	done, err := p.gateway.GetReceiptByTxHash(txHash)

	receipt := <-done
	log.Printf("%+v\n", receipt)
	colorLog.Green("%+v\n", receipt)
	colorLog.Cyan("End of TranferPCM!!")
}
