package handler

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"c2n/model"
	"c2n/service"
	"c2n/utils"
)

type EncodeHandler struct {
	EncodeService *service.EncodeService
}

func NewEncodeHandler(encodeService *service.EncodeService) *EncodeHandler {
	return &EncodeHandler{
		EncodeService: service.NewEncodeService(),
	}
}

func (h *EncodeHandler) SignRegistration(c *gin.Context) {
	userAddress := c.PostForm("userAddress")
	contractAddress := c.PostForm("contractAddress")
	log.Printf("userAddress: %s, contractAddress: %s", userAddress, contractAddress)
	if userAddress == "" || contractAddress == "" {
		panic("userAddress or contractAddress is empty")
	}
	contractAddr := utils.CleanHexPrefix(contractAddress)
	userAddr := utils.CleanHexPrefix(userAddress)
	concat := strings.ToLower(userAddr + contractAddr)
	hex := common.HexToHash(concat).Hex()
	sign := h.EncodeService.Sign(hex)
	c.JSON(http.StatusOK, model.OkWithData[string](sign))
}
