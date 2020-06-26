package quiz

import (
	"testing"
)

func Test_BalanceOf(t *testing.T) {
	p := NewPCM()
	p.BalanceOf("0x03CDD0B4878BA94BF8bFadD6c2C2241759Fd158A")
	if PrintLog {
		t.Fail()
	}
}
