package utils

import (
	"fmt"
	"testing"
)

func TestCleanHexPrefix(t *testing.T) {
	hexString := "0x1234567890"
	result := CleanHexPrefix(hexString)
	if result != "1234567890" {
		t.Errorf("CleanHexPrefix failed, expected %s, got %s", "1234567890", result)
	}
}

func TestGetSign(t *testing.T) {
	hexString := "0xDe03b04D52EbA5B27F4460d450Fae531A1A95cB0"
	result := GetSign(hexString)
	fmt.Println(result)
}
