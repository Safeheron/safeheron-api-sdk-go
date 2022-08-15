package mpcdemo

import (
	"testing"
)

func TestSendEther(t *testing.T) {
	// Replace with your address
	SendEther("0x53B1****321789", 0.01)
}

func TestSendERC20Token(t *testing.T) {
	// Replace with your address
	SendERC20Token("0x53B1****321789", 10)
}
