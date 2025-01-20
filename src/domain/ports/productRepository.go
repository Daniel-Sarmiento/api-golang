package ports

import (
	"api/src/domain"
)

type IProductRepository interface {
	Save(product domain.Product) error
	GetAll() ([]domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Update(id int, product domain.Product) error
	Delete(id int) error
}