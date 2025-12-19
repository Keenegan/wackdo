package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"wackdo/src/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=disable TimeZone=Europe/Paris"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la DB : %v", err)
	}

	db.AutoMigrate()

	fmt.Println("Connexion à PostgreSQL réussie avec GORM !")

	r := gin.Default()

	r.Use(cors.Default())

	db.AutoMigrate(&models.Employee{})
	insertEmployee := &models.Employee{Name: "Michel", Roles: []string{"EMPLOYEE"}}
	db.Create(insertEmployee)

	fmt.Printf("insert ID: %d, Code: %s, Price: %d\n",
		insertEmployee.ID, insertEmployee.Name, insertEmployee.Roles)

	readEmployee := &models.Employee{}
	db.First(&readEmployee, "Name = ?", "Michel")

	fmt.Printf(readEmployee.Name)

	r.GET("/status", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	r.Run()
}
