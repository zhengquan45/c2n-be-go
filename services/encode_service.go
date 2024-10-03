package services

import "log"

func Sign(hexString string) string {
	log.Printf("signing hex string: %s", hexString)
	return hexString
}
