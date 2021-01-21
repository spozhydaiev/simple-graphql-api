package repo

import (
	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"gorm.io/gorm"
)

type ArtworkRepo interface {
	Create(artwork *models.Artwork) error
	Update(artwork *models.Artwork) error
	Delete(artwork *models.Artwork) error
	GetArtworks(artistID int) ([]*models.Artwork, error)
}

type artworkRepo struct {
	db *gorm.DB
}

func NewArtworkRepo(db *gorm.DB) ArtworkRepo {
	return &artworkRepo{
		db: db,
	}
}

func (repo *artworkRepo) Create(artwork *models.Artwork) error {
	err := repo.db.Model(models.Artwork{}).Create(artwork).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artworkRepo) Update(artwork *models.Artwork) error {
	err := repo.db.Model(&models.Artwork{}).Where("id", artwork.ID).Updates(artwork).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artworkRepo) Delete(artwork *models.Artwork) error {
	err := repo.db.Model(&models.Artwork{}).Where("id", artwork.ID).Delete(artwork).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artworkRepo) GetArtworks(artistID int) ([]*models.Artwork, error) {
	var result []*models.Artwork

	err := repo.db.Model(&models.Artwork{}).Where("artist_id", artistID).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
