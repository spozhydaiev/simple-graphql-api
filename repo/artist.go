package repo

import (
	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"gorm.io/gorm"
)

type ArtistRepo interface {
	Create(artist *models.Artist) error
	Update(artist *models.Artist) error
	Delete(artist *models.Artist) error
	GetArtistList(searchString string) ([]*models.Artist, error)
	GetArtistByName(name string) (*models.Artist, error)
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
	err := repo.db.Model(&models.Artist{}).Where("id", artist.ID).Update("name", artist.Name).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artistRepo) Delete(artist *models.Artist) error {
	err := repo.db.Model(&models.Artist{}).Where("id", artist.ID).Delete(artist).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artistRepo) GetArtistList(searchString string) ([]*models.Artist, error) {
	var result []*models.Artist

	chain := repo.db.Model(&models.Artist{})

	if searchString != "" {
		chain = chain.Where("LOWER(name) LIKE LOWER(?)", "%"+searchString+"%")
	}

	err := chain.Preload("Artworks").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *artistRepo) GetArtistByName(name string) (*models.Artist, error) {
	var result *models.Artist

	err := repo.db.Model(&models.Artist{}).Where("name", name).Preload("Artworks").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
