package main

import (
	"wackdo/src/controllers"
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

	r.Run()
}

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}
