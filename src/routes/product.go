package routes

import (
	"api/src/controllers"
	"api/src/repositories"

	"github.com/gin-gonic/gin"
)

func AddProductRoutes(engine *gin.Engine) {

	productRepository := repositories.NewProductRepository()
	productController := controllers.NewProductController(productRepository)
	
	productsRoutes := engine.Group("products")
	{
		productsRoutes.GET("", productController.GetAll)
	}

}