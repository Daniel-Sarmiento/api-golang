package main

import (
	"api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	myGin := gin.Default()
	
	routes.AddProductRoutes(myGin)

	myGin.Run()
}
