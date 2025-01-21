package adapters

import (
	"api/src/domain"
	"errors"
	"fmt"
)

type MySQLProductRepository struct {
	products []domain.Product
}

func NewMySQLProductRepository() *MySQLProductRepository {
	return &MySQLProductRepository{
		products: []domain.Product{},
	}
}

func (r *MySQLProductRepository) Save(product domain.Product) error {
	r.products = append(r.products, product)
	return nil
} 

func (r MySQLProductRepository) GetAll() ([]domain.Product, error) {
	return r.products, nil
}

func (r MySQLProductRepository) GetByID(id int) (*domain.Product, error) {
	for _, myProduct := range r.products {
		if myProduct.GetId() == id {
			return &myProduct, nil
		}
	}
	
	return nil, errors.New("no existe el producto")
}

func (r *MySQLProductRepository) Update(id int, updatedProduct domain.Product) error {
	for i, myProduct := range r.products {
		if myProduct.GetId() == id {
			r.products[i] = updatedProduct
			fmt.Println("Product: ", updatedProduct)
			return nil
		}
	}

	return errors.New("producto no encontrado")
}

func (r *MySQLProductRepository) Delete(id int) error {
	for i, myProduct := range r.products {
		if myProduct.GetId() == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			fmt.Println("Producto eliminado: ", id)
			return nil
		}
	}
	
	return errors.New("producto no encontrado")
}