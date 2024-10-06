package handler

import (
	"c2n/api/request"
	"c2n/api/response"
	"c2n/enums"
	"c2n/middleware"
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
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}

	productContract, err := h.ProductContractService.GetById(productId)
	if err != nil {
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}
	productContractResponse := response.ProductContractToProductContractResponse(*productContract)
	c.JSON(http.StatusOK, model.OkWithData(productContractResponse))
}

func (h *ProductContractHandler) List(c *gin.Context) {
	productContracts, err := h.ProductContractService.List()
	if err != nil {
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}
	var productContractResponses []response.ProductContractResponse
	for _, productContract := range productContracts {
		productContractResponses = append(productContractResponses, response.ProductContractToProductContractResponse(productContract))
	}

	c.JSON(http.StatusOK, model.OkWithData(productContractResponses))
}

func (h *ProductContractHandler) Update(c *gin.Context) {
	var productContractUpdateRequest request.ProductContractUpdateRequest
	if err := c.ShouldBindJSON(&productContractUpdateRequest); err != nil {
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}
	h.ProductContractService.Update(&productContractUpdateRequest)
	c.JSON(http.StatusOK, model.Ok[string]())
}
