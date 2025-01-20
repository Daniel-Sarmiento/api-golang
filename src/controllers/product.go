package controllers

import (
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productRepository *repositories.ProductRepository
}

func NewProductController(productRepository *repositories.ProductRepository) *ProductController {
	return &ProductController{
		productRepository: productRepository,
	}
}

func (ctrl ProductController) GetAll(c *gin.Context) {
	products, err := ctrl.productRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "ocurrió un error al obtener los productos",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "se ubtuvieron los productos exitosamente",
		Data:    products,
	})
}

func (ctrl ProductController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "id inválido",
			Error:   err.Error(),
		})
		return
	}

	product, err := ctrl.productRepository.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "ocurrió un error al obtener el productos",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "se ubtuvo el producto exitosamente",
		Data:    product,
	})
}

func (ctrl ProductController) Create(c *gin.Context) {
	var newProduct models.Product // creo una variable para el nuevo producto

	// intento leer el json del body y lo paso a la variable newProduct
	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "ocurrió un error al leer la petición",
			Error:   err.Error(),
		})
		return
	}

	log.Println(newProduct)

	err = ctrl.productRepository.Save(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "ocurrió un error al guardar el producto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "producto creado exitosamente",
		Data: newProduct,
	})
}

func (ctrl ProductController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "id inválido",
			Error:   err.Error(),
		})
		return
	}

	var updatedProduct models.Product // creo una variable para el producto actualizado

	// intento leer el json del body y lo paso a la variable updatedProduct
	err = c.ShouldBindJSON(&updatedProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "ocurrió un error al leer la petición",
			Error: err.Error(),
		})
		return
	}

	err = ctrl.productRepository.Update(id, updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "ocurrió un error al actualizar el producto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "producto actualizado exitosamente",
	})
}

func (ctrl ProductController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{
			Success: false,
			Message: "id inválido",
			Error:   err.Error(),
		})
		return
	}

	err = ctrl.productRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.Response{
			Success: false,
			Message: "no se pudo eliminar el producto",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.Response{
		Success: true,
		Message: "producto eliminado exitosamente",
	})
}
