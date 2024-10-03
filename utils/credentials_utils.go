package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func getSign(hexString string) string {
	messageHash := common.Hex2Bytes(hexString)
	// 加上以太坊签名前缀
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

	return ""
}

func CleanHexPrefix(hexString string) string {
	if strings.HasPrefix(hexString, "0x") {
		return hexString[2:] // 返回去掉前缀后的字符串
	}
	return hexString // 如果没有前缀，则返回原始字符串
}
