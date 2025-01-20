package routes

import (
	"api/src/controllers"
	"api/src/repositories"

	"github.com/gin-gonic/gin"
)

func AddProductRoutes(engine *gin.Engine) {

	productRepository := repositories.NewProductRepository()
	productController := controllers.NewProductController(productRepository)
	
	productsRoutes := engine.Group("/products")
	{
		productsRoutes.GET("", productController.GetAll)
		productsRoutes.GET("/:id", productController.GetById)
		productsRoutes.POST("", productController.Create)
		productsRoutes.PUT("/:id", productController.Update)
		productsRoutes.DELETE("/:id", productController.Delete)
	}

}