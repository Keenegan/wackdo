package initializers

import (
	"os"
	"wackdo/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	var dsn string

	// Use DATABASE_URL if available (Render provides this)
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		dsn = databaseURL
	} else {
		// Fallback to individual env vars for local development
		host := os.Getenv("POSTGRES_HOST")
		if host == "" {
			host = "localhost"
		}
		dsn = "host=" + host +
			" user=" + os.Getenv("POSTGRES_USER") +
			" password=" + os.Getenv("POSTGRES_PASSWORD") +
			" dbname=" + os.Getenv("POSTGRES_DB") +
			" port=" + os.Getenv("POSTGRES_PORT") +
			" sslmode=disable TimeZone=Europe/Paris"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database : " + err.Error())
	}

	DB.AutoMigrate(&models.Product{}, &models.Menu{}, &models.User{}, &models.Order{}, &models.OrderLine{})
}
