package routes

import (
	"api/src/infraestructure/http/controllers"

	"github.com/gin-gonic/gin"
)


func RegisterProductsRoutes(engine *gin.Engine, getAllProductsController *controllers.GetAllProductsController, createProductController *controllers.CreateProductController) {
	
	productsRoutes := engine.Group("/products") 
	{
		productsRoutes.GET("", getAllProductsController.Run)
		productsRoutes.POST("", createProductController.Run)
	}

}
