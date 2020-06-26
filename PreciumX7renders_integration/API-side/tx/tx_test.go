package tx

import (
	"testing"
)

//var gateway *EtherGateway

func makeGateway() *EtherGateway {
	return NewEtherGateway()
}
func Test_loadConfigs(t *testing.T) {
	makeGateway()
}
func Test_NetworkType(t *testing.T) {
	g := makeGateway()
	if networkID := g.GetCurrentNetworkType(); networkID == "" {
		t.Errorf("Empty network id!!")
	}
}

func Test_GasPrice(t *testing.T) {
	g := makeGateway()
	if gasPrice := g.GetCurrentGasPrice(); gasPrice == 0 {
		t.Errorf("Invalid GasPrice")
	}
}
func Test_TxEther(t *testing.T) {
	g := makeGateway()
	//g.txEther(0.004, "0x7aaf4fcB8AB215f719E6DaBb7f0a192b7024dD21")
}
