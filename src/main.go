package main

import (
	"time"
	"wackdo/src/controllers"
	controllers_menu "wackdo/src/controllers/menu"
	controllers_order "wackdo/src/controllers/order"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Only for development purpose
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/status", controllers.Status)

	// Only managers can manage menu
	menuRoutes := r.Group("/menu")
	menuRoutes.Use(middleware.AuthMiddleware(models.RoleManager))
	{
		menuRoutes.POST("", controllers_menu.PostMenu)
		menuRoutes.GET("", controllers_menu.GetMenus)
		menuRoutes.GET("/:id", controllers_menu.GetMenuById)
		menuRoutes.GET("/search", controllers_menu.GetMenuByName)
		menuRoutes.DELETE("/:id", controllers_menu.DeleteMenu)
		menuRoutes.PATCH("", controllers_menu.UpdateMenu)
	}

	// Only managers can manage products
	productRoutes := r.Group("/product")
	productRoutes.Use(middleware.AuthMiddleware(models.RoleManager))
	{
		productRoutes.POST("", controllers_products.PostProduct)
		productRoutes.GET("", controllers_products.GetProducts)
		productRoutes.GET("/:id", controllers_products.GetProductById)
		productRoutes.GET("/search", controllers_products.GetProductByName)
		productRoutes.DELETE("/:id", controllers_products.DeleteProduct)
		productRoutes.PATCH("", controllers_products.UpdateProduct)
	}

	// Everyone can login & register (but only as an employee)
	r.POST("/register", controllers_user.Register)
	r.POST("/login", controllers_user.Login)

	// Only manager can update user role & email
	r.PATCH(
		"/user/:id",
		middleware.AuthMiddleware(
			models.RoleManager,
		),
		controllers_user.UpdateUser,
	)

	// Only manager can list all users
	r.GET("/users",
		middleware.AuthMiddleware(
			models.RoleManager,
		),
		controllers_user.GetUsers)

	// Only employees and managers can create orders
	r.POST("/order/",
		middleware.AuthMiddleware(
			models.RoleEmployee,
			models.RoleManager,
		),
		controllers_order.PostOrder)

	// Everyone can see orders
	r.GET("/orders", controllers_order.GetOrders)
	r.GET("/order/:id", controllers_order.GetOrder)

	// Only prep & managers can update order status
	r.PATCH("/order/:id",
		middleware.AuthMiddleware(
			models.RolePrep,
			models.RoleManager,
		),
		controllers_order.PatchOrder)

	r.Run()
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}
