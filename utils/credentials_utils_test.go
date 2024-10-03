package utils

import (
	"testing"
)

func TestCleanHexPrefix(t *testing.T) {
	hexString := "0x1234567890"
	result := CleanHexPrefix(hexString)
	if result != "1234567890" {
		t.Errorf("CleanHexPrefix failed, expected %s, got %s", "1234567890", result)
	}
}
