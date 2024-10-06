package service

import (
	"c2n/api/request"
	"c2n/model"
	"c2n/repository"
)

type ProductContractService struct {
	ProductContractRepository *repository.ProductContractRepository // Ensure this type is defined in the repository package
}

func NewProductContractService(productContractRepository *repository.ProductContractRepository) *ProductContractService {
	return &ProductContractService{
		ProductContractRepository: productContractRepository,
	}
}

func (s *ProductContractService) GetById(productContractId string) (*model.ProductContract, error) {
	return s.ProductContractRepository.GetById(productContractId)
}

func (s *ProductContractService) List() ([]model.ProductContract, error) {
	return s.ProductContractRepository.List()
}

func (s *ProductContractService) Update(productContractUpdateRequest *request.ProductContractUpdateRequest) error {
	return s.ProductContractRepository.Update(productContractUpdateRequest)
}
