package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=Europe/Paris"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la DB : %v", err)
	}

	fmt.Println("Connexion à PostgreSQL réussie avec GORM !")

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/status", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.Run()
}
