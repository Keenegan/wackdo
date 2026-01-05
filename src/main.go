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

	r.GET("/status", controllers.Status)

	r.POST("/menu", controllers_menu.PostMenu)
	r.GET("/menu", controllers_menu.GetMenu)
	r.DELETE("/menu", controllers_menu.DeleteMenu)
	r.PATCH("/menu", controllers_menu.UpdateMenu)

	r.POST("/product", controllers_products.PostProduct)
	r.GET("/products", controllers_products.GetProducts)
	r.DELETE("/product", controllers_products.DeleteProduct)
	r.PATCH("/product", controllers_products.UpdateProduct)

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
