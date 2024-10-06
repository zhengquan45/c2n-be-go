package service

import (
	utils "c2n/utils"
	"log"
)

type EncodeService struct {
}

func NewEncodeService() *EncodeService {
	return &EncodeService{}
}

func (s *EncodeService) Sign(hexString string) string {
	log.Printf("signing hex string: %s", hexString)
	return utils.GetSign(hexString)
}
