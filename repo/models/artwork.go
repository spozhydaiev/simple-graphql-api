package models

import (
	"time"

	"gorm.io/gorm"
)

type Artwork struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Title    string `json:"title"`
	ArtistID uint   `json:"artist"`
}
