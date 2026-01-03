package main

import (
	"wackdo/src/controllers"
	controllers_menu "wackdo/src/controllers/menu"
	controllers_products "wackdo/src/controllers/product"
	controllers_user "wackdo/src/controllers/user"
	"wackdo/src/initializers"
	"wackdo/src/models"
	"wackdo/src/service/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(middleware.ErrorMiddleware())

	r.Use(cors.Default())

	r.GET("/status", func(c *gin.Context) {
		controllers.Status(c)
	})

	r.POST("/menu", func(c *gin.Context) {
		controllers_menu.PostMenu(c)
	})

	r.GET("/menu", func(c *gin.Context) {
		controllers_menu.GetMenu(c)
	})

	r.DELETE("/menu", func(c *gin.Context) {
		controllers_menu.DeleteMenu(c)
	})

	r.PATCH("/menu", func(c *gin.Context) {
		controllers_menu.UpdateMenu(c)
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

	r.POST("/register", controllers_user.Register)
	r.POST("/login", controllers_user.Login)

	r.PATCH(
		"/user/:id",
		middleware.AuthMiddleware(
			models.RoleAdmin,
			models.RoleManager,
		),
		controllers_user.UpdateUser,
	)

	r.Run()
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}
