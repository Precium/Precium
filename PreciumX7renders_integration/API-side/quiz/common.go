package quiz

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

const PrintLog bool = true

const (
	SubConfigController     string = "controller"
	SubConfigWalletContract string = "userWallet"
	SubconfigPrecium        string = "precium"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func loadSubConfig(subName string) (*viper.Viper, string, error) {
	file, err := os.Open(basePath[:len(basePath)-4] + "envs.yaml")
	workdir, _ := os.Getwd()
	file2, err2 := os.Open(workdir + "/envs.yaml")
	if err != nil && err2 != nil {
		fmt.Println(err)
		return nil, "", err
	} else if err != nil && err2 == nil {
		file = file2
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}
	viper.SetConfigType("YAML")
	var subConfig *viper.Viper
	viper.ReadConfig(bytes.NewBuffer(buffer))
	currentEnvName := viper.GetString("envs")
	fmt.Println("Loaded CURRENT ENV PARAM ", currentEnvName)
	if strings.Compare("", currentEnvName) == 0 {
		currentEnvName = "dev"
	}
	fmt.Println("Working ENV param : ", currentEnvName)

	subConfig = viper.Sub(currentEnvName + "." + subName)
	return subConfig, currentEnvName, nil

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
