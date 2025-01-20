package infraestructure

import (
	"api/src/domain"
	"errors"
	"fmt"
)

type MysqlProductRepository struct {
	products []domain.Product
}

func NewMysqlProductRepository() *MysqlProductRepository {
	return &MysqlProductRepository{}
}

func (r *MysqlProductRepository) Save(product domain.Product) error {
	r.products = append(r.products, product)
	return nil
} 

func (r MysqlProductRepository) GetAll() ([]domain.Product, error) {
	return r.products, nil
}

func (r MysqlProductRepository) GetByID(id int) (*domain.Product, error) {
	for _, myProduct := range r.products {
		if myProduct.GetId() == id {
			return &myProduct, nil
		}
	}
	
	return nil, errors.New("no existe el producto")
}

func (r *MysqlProductRepository) Update(id int, updatedProduct domain.Product) error {
	for i, myProduct := range r.products {
		if myProduct.GetId() == id {
			r.products[i] = updatedProduct
			fmt.Println("Product: ", updatedProduct)
			return nil
		}
	}

	return errors.New("producto no encontrado")
}

func (r *MysqlProductRepository) Delete(id int) error {
	for i, myProduct := range r.products {
		if myProduct.GetId() == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			fmt.Println("Producto eliminado: ", id)
			return nil
		}
	}
	
	return errors.New("producto no encontrado")
}