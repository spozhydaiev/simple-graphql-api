package repo

import (
	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"gorm.io/gorm"
)

type ArtistRepo interface {
	Create(artist *models.Artist) error
	Update(artist *models.Artist) error
	Delete(artist *models.Artist) error
	GetArtist() ([]*models.Artist, error)
}

type artistRepo struct {
	db *gorm.DB
}

func NewArtistRepo(db *gorm.DB) ArtistRepo {
	return &artistRepo{
		db: db,
	}
}

func (repo *artistRepo) Create(artist *models.Artist) error {
	err := repo.db.Create(artist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artistRepo) Update(artist *models.Artist) error {
	err := repo.db.Where("id =", artist.ID).Updates(artist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artistRepo) Delete(artist *models.Artist) error {
	err := repo.db.Where("id =", artist.ID).Delete(artist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artistRepo) GetArtist() ([]*models.Artist, error) {
	var result []*models.Artist

	err := repo.db.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
