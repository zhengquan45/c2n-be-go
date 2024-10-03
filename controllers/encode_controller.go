package controllers

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"c2n/models"
	"c2n/services"
	"c2n/utils"
)

func SignRegistration(c *gin.Context) {
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
	sign := services.Sign(hex)
	c.JSON(http.StatusOK, models.OkWithData[string](sign))
}
