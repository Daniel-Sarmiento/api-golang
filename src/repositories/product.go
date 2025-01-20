package repositories

import (
	"api/src/models"
	"errors"
	"fmt"
)

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: []models.Product{},
	}
}

func (r *ProductRepository) Save(product models.Product) error {
	r.products = append(r.products, product)
	fmt.Println("Product saved:", product)
	return nil
}

func (r ProductRepository) GetAll() ([]models.Product, error) {
	return r.products, nil
}

func (r ProductRepository) GetByID(id int32) (*models.Product, error) {
	for _, myProduct := range r.products {
		if myProduct.GetId() == id {
			return &myProduct, nil
		}
	}

	return nil, errors.New("no existe el producto ")
}

func (r *ProductRepository) Update(id int32, updatedProduct models.Product) error {
	for i, myProduct := range r.products {
		if myProduct.GetId() == id {
			r.products[i] = updatedProduct
			fmt.Println("Product: ", updatedProduct)
			return nil
		}
	}

	return errors.New("producto no encontrado")
}