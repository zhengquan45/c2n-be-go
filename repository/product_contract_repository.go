package repository

import (
	"c2n/api/request"
	"c2n/model"

	"gorm.io/gorm"
)

type ProductContractRepository struct {
	DB *gorm.DB
}

func NewProductContractRepository(db *gorm.DB) *ProductContractRepository {
	return &ProductContractRepository{DB: db}
}

func (r *ProductContractRepository) GetById(productContractId string) (*model.ProductContract, error) {
	var productContract model.ProductContract

	if err := r.DB.Where("id = ?", productContractId).First(&productContract).Error; err != nil {
		return nil, err
	}
	return &productContract, nil
}

func (r *ProductContractRepository) List() ([]model.ProductContract, error) {
	var productContracts []model.ProductContract

	if err := r.DB.Find(&productContracts).Error; err != nil {
		return nil, err
	}
	return productContracts, nil
}

func (r *ProductContractRepository) Update(productContractUpdateRequest *request.ProductContractUpdateRequest) error {
	if err := r.DB.Model(&model.ProductContract{}).Where("id = ?", productContractUpdateRequest.ID).Select(
		"SaleContractAddress",
		"TokenAddress",
		"TokenPriceInPT",
		"TotalTokensSold",
		"SaleEnd",
		"UnlockTime",
		"RegistrationTimeStarts",
		"RegistrationTimeEnds",
		"SaleStart",
	).Updates(model.ProductContract{
		SaleContractAddress:    productContractUpdateRequest.SaleAddress,
		TokenAddress:           productContractUpdateRequest.SaleToken,
		TokenPriceInPT:         productContractUpdateRequest.TokenPriceInPT,
		TotalTokensSold:        productContractUpdateRequest.TotalToken,
		SaleEnd:                productContractUpdateRequest.SaleEndTime,
		UnlockTime:             productContractUpdateRequest.TokensUnlockTime,
		RegistrationTimeStarts: productContractUpdateRequest.RegistrationStart,
		RegistrationTimeEnds:   productContractUpdateRequest.RegistrationEnd,
		SaleStart:              productContractUpdateRequest.SaleStartTime,
	}).Error; err != nil {
		return err
	}
	return nil
}
