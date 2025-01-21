package application

import (
	"api/src/domain"
	"api/src/domain/ports"
)

type GetAllProductsUseCase struct {
	productRepository ports.IProductRepository
}

func NewGetAllProductsUseCase (productRepository ports.IProductRepository) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{
		productRepository: productRepository,
	}
}

func (uc GetAllProductsUseCase) Run() ([]domain.Product, error) {
	return uc.productRepository.GetAll()
}