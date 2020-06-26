package tx

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Have to match each other `contractFiles` and const value.
const (
	// for precium quiz(mini project)
	ContractTypeQuizWallet = iota
	ContractTypeController
	ContractTypePrecium
	//ContractTypeQuizAuthrizedCaller

)

func hex2int(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, err := strconv.ParseUint(cleaned, 16, 64)
	if err != nil {
		return uint64(0)
	}
	return uint64(result)
}

func getContractNames(currentEnvType string) []string {
	if currentEnvType == "dev" || currentEnvType == "alpha" {
		return []string{"abi/quiz/ropsten/wallet.json", "abi/quiz/ropsten/controller.json", "abi/quiz/ropsten/precium.json"}
	} else if currentEnvType == "beta" || currentEnvType == "prod" {
		return []string{"abi/quiz/mainnet/wallet.json", "abi/quiz/mainnet/controller.json", "abi/quiz/mainnet/precium.json"}
	}
	return []string{"abi/quiz/ropsten/wallet.json", "abi/quiz/ropsten/controller.json", "abi/quiz/ropsten/precium.json"}
}
func LoadContract(ContractType int, currentEnvType string) abi.ABI {
	contractList := getContractNames(currentEnvType)
	if ContractType >= len(contractList) || ContractType <= -1 {
		log.Fatal("invalid ContractType")
		//return nil
	}
	return *newABI(getContractNames(currentEnvType)[ContractType])
}

func LoadContractByString(ContractType int, curEnvs string) string {
	contractList := getContractNames(curEnvs)
	if ContractType >= len(contractList) || ContractType <= -1 {
		log.Fatal("invalid ContractType")
		//return nil
	}
	return newABIByString(getContractNames(curEnvs)[ContractType])
}

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func newABIByString(fileName string) string {
	file, err := os.Open(basePath + "\\" + fileName)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return ""
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(buffer)
}
func newABI(fileName string) *abi.ABI {
	file, err := os.Open(basePath + "/" + fileName)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return nil
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	abi, err := abi.JSON(bytes.NewBuffer(buffer))
	log.Println(`requested json file : `, fileName)
	if err != nil {
		log.Fatal(err)
	}

	for name, expM := range abi.Methods {
		log.Println("name : ", name, " / expM : ", expM)
	}
	return &abi
}
func ConvFloatEtherToWei(ether float64) *big.Int {
	gAmount := big.NewInt(int64(ether * 1e9))
	v := new(big.Int).Mul(gAmount, big.NewInt(1e9))
	return v
}

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
