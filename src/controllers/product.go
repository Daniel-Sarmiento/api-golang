package controllers

import (
	"api/src/repositories"
	"api/src/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository *repositories.ProductRepository
}

func NewProductController(productRepository *repositories.ProductRepository) *ProductController{
	return &ProductController{
		productRepository: productRepository,
	}
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	products, err := ctrl.productRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "ocurrió un erro al obtener los productos",
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "se ubtuvieron los productos exitosamente",
		Data: products,
	})
}