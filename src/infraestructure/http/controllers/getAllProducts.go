package controllers

import (
	"api/src/application"
	"api/src/infraestructure/http/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	getAllProductsUseCase *application.GetAllProductsUseCase
}

func NewGetAllProductsController(getAllProductsUseCase *application.GetAllProductsUseCase) *GetAllProductsController{
	return &GetAllProductsController{
		getAllProductsUseCase: getAllProductsUseCase,
	}
}

func (ctrl GetAllProductsController) Run(c *gin.Context) {
	products, err := ctrl.getAllProductsUseCase.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.GeneralResponse{
			Success: false,
			Message: "ocurrió un error al obtener los productos",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, responses.GeneralResponse{
		Success: true,
		Message: "se ubtuvieron los productos exitosamente",
		Data:    products,
	})
}