package migration

import (
	"log"

	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := "host=localhost user=postgres password=abkjcjabz24 dbname=gallery port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to database")
	}

	db.AutoMigrate(&models.Artist{}, &models.Artwork{})
}
