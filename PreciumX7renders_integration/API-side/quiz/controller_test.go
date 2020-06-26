package quiz

import (
	"log"
	"testing"
)

func Test_loadController(t *testing.T) {
	NewController()
	if PrintLog {
		t.Fail()
	}
}
func Test_GenerateWallet(t *testing.T) {
	controller := NewController()
	txHash, newAddress, err := controller.GenerateWallet()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("txHash :", txHash)
	log.Println("address :", newAddress)
	if PrintLog {
		t.Fail()
	}
}
