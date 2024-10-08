package handler

import (
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"c2n/enums"
	"c2n/middleware"
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

// @Summary Sign Registration
// @Description Sign Registration
// @Tags encode
// @Accept json
// @Produce json
// @Param userAddress formData string true "User Address"
// @Param contractAddress formData string true "Contract Address"
// @Success 200 {string} string "Sign"
// @Router /encode/sign_registration [post]
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

// @Summary Sign Participation
// @Description Sign Participation
// @Tags encode
// @Accept json
// @Produce json
// @Param userAddress formData string true "User Address"
// @Param contractAddress formData string true "Contract Address"
// @Param amount formData string true "Amount"
// @Success 200 {string} string "Sign"
// @Router /encode/sign_participation [post]
func (h *EncodeHandler) SignParticipation(c *gin.Context) {
	userAddress := c.PostForm("userAddress")
	contractAddress := c.PostForm("contractAddress")
	amount := c.PostForm("amount")
	log.Printf("userAddress: %s, contractAddress: %s amount: %s", userAddress, contractAddress, amount)
	if userAddress == "" || contractAddress == "" || amount == "" {
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}
	contractAddr := utils.CleanHexPrefix(contractAddress)
	userAddr := utils.CleanHexPrefix(userAddress)
	//TODO Amount should be converted to hex toHexStringNoPrefixZeroPadded
	concat := strings.ToLower(userAddr + amount + contractAddr)
	hex := common.HexToHash(concat).Hex()
	sign := h.EncodeService.Sign(hex)
	c.JSON(http.StatusOK, model.OkWithData[string](sign))
}
