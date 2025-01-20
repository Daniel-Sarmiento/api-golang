package routes

import (
	"api/src/infraestructure/api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func AddProductRoutes(engine *gin.Engine, productController controllers.CreateProductController) {

	/* productRepository := repositories.NewProductRepository()
	productController := controllers.NewProductController(productRepository)

	productsRoutes := engine.Group("/products")
	{
		productsRoutes.GET("", productController.GetAll)
		productsRoutes.GET("/:id", productController.GetById)
		productsRoutes.POST("", productController.Create)
		productsRoutes.PUT("/:id", productController.Update)
		productsRoutes.DELETE("/:id", productController.Delete)
	} */

}