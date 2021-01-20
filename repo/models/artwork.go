package models

import "gorm.io/gorm"

type Artwork struct {
	gorm.Model
	Title    string `json:"title"`
	ArtistID uint   `json:"artist"`
}
