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
		productsRoutes.GET("/products/:id", productController.GetById)
		productsRoutes.POST("", productController.Create)
		productsRoutes.PUT("/products/:id", productController.Update)
		productsRoutes.DELETE("/products/:id", productController.Delete)
	}

}