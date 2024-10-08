package utils

import (
	"c2n/config"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanHexPrefix(t *testing.T) {
	hexString := "0x1234567890"
	result := CleanHexPrefix(hexString)
	if result != "1234567890" {
		t.Errorf("CleanHexPrefix failed, expected %s, got %s", "1234567890", result)
	}
}

func TestGetSign(t *testing.T) {
	// 加载配置文件
	config.LoadConfig()
	hexString := "0xDe03b04D52EbA5B27F4460d450Fae531A1A95cB0"
	result := GetSign(hexString)
	fmt.Println(result)
	assert.Equal(t, result, "a698ea12d20c92fa698bf306910ca718bcdcb1bc3f35e743e8b0d44d150132a176e1406e86e4f51e4629b149746a2fe4e2efb313bf1b7fa85ba6f573beb332b41c")
}
