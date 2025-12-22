package main

import (
	"wackdo/src/controllers"
	controllers_products "wackdo/src/controllers/product"
	"wackdo/src/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/status", func(c *gin.Context) {
		controllers.Status(c)
	})

	r.POST("/employee", func(c *gin.Context) {
		controllers.PostEmployee(c)
	})

	r.POST("/product", func(c *gin.Context) {
		controllers_products.PostProduct(c)
	})

	r.GET("/products", func(c *gin.Context) {
		controllers_products.GetProducts(c)
	})

	r.DELETE("/product", func(c *gin.Context) {
		controllers_products.DeleteProduct(c)
	})

	r.PATCH("/product", func(c *gin.Context) {
		controllers_products.UpdateProduct(c)
	})

	r.Run()
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}
