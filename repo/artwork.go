package repo

import (
	"context"

	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"gorm.io/gorm"
)

type ArtworkRepo interface {
	Create(ctx context.Context, artwork *models.Artwork) error
	Update(ctx context.Context, artwork *models.Artwork) error
	Delete(ctx context.Context, artwork *models.Artwork) error
}

type artworkRepo struct {
	db *gorm.DB
}

func NewArtworkRepo(db *gorm.DB) ArtworkRepo {
	return &artworkRepo{
		db: db,
	}
}

func (repo *artworkRepo) Create(ctx context.Context, artwork *models.Artwork) error {
	err := repo.db.Model(models.Artwork{}).Create(artwork).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artworkRepo) Update(ctx context.Context, artwork *models.Artwork) error {
	err := repo.db.Model(models.Artwork{}).Updates(artwork).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *artworkRepo) Delete(ctx context.Context, artwork *models.Artwork) error {
	err := repo.db.Model(models.Artwork{}).Delete(artwork).Error
	if err != nil {
		return err
	}
	return nil
}
