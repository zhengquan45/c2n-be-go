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

// @Summary Product Base Info
// @Description Get product base info
// @Tags product
// @Accept json
// @Produce json
// @Param productContractId query string true "Product Contract Id"
// @Success 200 {object} model.Result[response.ProductContractResponse]
// @Router /product/base_info [get]
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

// @Summary Product List
// @Description Get product list
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} model.Result[response.ProductContractResponse]
// @Router /product/list [get]
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

// @Summary Update Product Contract
// @Description Update product contract
// @Tags product
// @Accept json
// @Produce json
// @Param productContractId formData string true "Product Contract Id"
// @Param name formData string true "Name"
// @Param description formData string true "Description"
// @Param price formData string true "Price"
// @Param status formData string true "Status"
// @Success 200 {object} model.Result[string]
// @Router /update [post]
func (h *ProductContractHandler) Update(c *gin.Context) {
	var productContractUpdateRequest request.ProductContractUpdateRequest
	if err := c.ShouldBindJSON(&productContractUpdateRequest); err != nil {
		panic(&middleware.BusinessError{ReCode: enums.INVALID_PARAMETERS})
	}
	h.ProductContractService.Update(&productContractUpdateRequest)
	c.JSON(http.StatusOK, model.Ok[string]())
}
