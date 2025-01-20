package application

import (
	"api/src/domain"
	"api/src/domain/ports"
)

type CreateProductUseCase struct {
	productRepository ports.IProductRepository
}

func NewCreateProductUseCase(productRepository ports.IProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (uc CreateProductUseCase) Run(product domain.Product) error {
	return uc.productRepository.Save(product)
}