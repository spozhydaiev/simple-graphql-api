package repo

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func New(conn string) (*DB, error) {
	dsn := "host=localhost user=postgres password=abkjcjabz24 dbname=gallery port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
