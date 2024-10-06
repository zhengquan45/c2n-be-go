package repository

import (
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
