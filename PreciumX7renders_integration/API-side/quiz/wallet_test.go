package quiz

import (
	"testing"
)

func Test_loadWallet(t *testing.T) {
	NewWallet()
	t.Fail()
}
func Test_Sweep(t *testing.T) {
	w := NewWallet()
	w.Sweep("0x881d035cbe642dedb2c6ab959434f1a58c488f56", 9.8765)
	t.Fail()
}
func Test_SweepedQty(t *testing.T) {
	w := NewWallet()
	w.SweepedQty("0x881d035cbe642dedb2c6ab959434f1a58c488f56")
	t.Fail()
}
