package models

import (
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	Artworks  []Artwork      `gorm:"foreignKey:ArtistID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"artworks"`
}
