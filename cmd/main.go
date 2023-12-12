package main

import (
	"app/cmd/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	productGroup := server.Group("/products")

	router := handler.NewProductHandler(productGroup)

	router.ProductRoutes()

	server.Run(":8080")

}
