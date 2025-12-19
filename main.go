package main

import "fmt"
import "os"
import "log"
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=5432 sslmode=disable TimeZone=Europe/Paris"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Impossible de se connecter à la DB : %v", err)
	}

	fmt.Println("Connexion à PostgreSQL réussie avec GORM !")
}