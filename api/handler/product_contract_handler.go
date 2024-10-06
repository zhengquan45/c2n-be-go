package handler

import (
	"c2n/model"
	"c2n/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductContractHandler struct {
	ProductContractService *service.ProductContractService
}

func NewProductContractHandler(productContractService *service.ProductContractService) *ProductContractHandler {
	return &ProductContractHandler{
		ProductContractService: productContractService,
	}
}

func (h *ProductContractHandler) BaseInfo(c *gin.Context) {
	productId := c.Query("productContractId")
	if productId == "" {
		panic("productContractId is empty")
	}

	productContract, err := h.ProductContractService.GetById(productId)
	if err != nil {

	}
	c.JSON(http.StatusOK, model.OkWithData(productContract))
}
