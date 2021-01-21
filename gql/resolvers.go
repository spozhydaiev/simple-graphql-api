package gql

import (
	"errors"
	"log"

	"github.com/SergioBravo/simple-graphql-api/repo"
	"github.com/SergioBravo/simple-graphql-api/repo/models"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type resolver struct {
	db *gorm.DB
}

func NewResolver(db *gorm.DB) *resolver {
	return &resolver{
		db: db,
	}
}

func (r *resolver) CreateArtist(p graphql.ResolveParams) (interface{}, error) {
	artist := &models.Artist{
		Name: p.Args["name"].(string),
	}
	db := repo.NewArtistRepo(r.db)
	err := db.Create(artist)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (r *resolver) UpdateArtist(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["artistId"].(int)
	if ok {
		artist := &models.Artist{
			ID:   uint(id),
			Name: p.Args["name"].(string),
		}
		db := repo.NewArtistRepo(r.db)
		err := db.Update(artist)
		if err != nil {
			return nil, err
		}
		return artist, nil
	}
	return nil, errors.New("unable to get artist id from arguments")
}

func (r *resolver) DeleteArtist(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["artistId"].(int)
	if ok {
		artist := &models.Artist{
			ID: uint(id),
		}
		db := repo.NewArtistRepo(r.db)
		err := db.Delete(artist)
		if err != nil {
			return nil, err
		}
		return artist, nil
	}
	return nil, errors.New("unable to get artist id from arguments")
}

func (r *resolver) GetArtistList(p graphql.ResolveParams) (interface{}, error) {
	db := repo.NewArtistRepo(r.db)
	name, ok := p.Args["name"].(string)
	if ok {
		artists, err := db.GetArtistList(name)
		if err != nil {
			return nil, err
		}

		return artists, nil
	}
	return nil, errors.New("unable to get artist id from arguments")
}

func (r *resolver) GetArtistByName(p graphql.ResolveParams) (interface{}, error) {
	db := repo.NewArtistRepo(r.db)
	name, ok := p.Args["name"].(string)
	if ok {
		artists, err := db.GetArtistByName(name)
		if err != nil {
			return nil, err
		}

		return artists, nil
	}
	return nil, errors.New("unable to get artist id from arguments")
}

func (r *resolver) GetArtistArtworks(p graphql.ResolveParams) (interface{}, error) {
	db := repo.NewArtworkRepo(r.db)
	id, ok := p.Args["artistId"].(int)
	log.Println(ok)

	if ok {
		artists, err := db.GetArtworks(id)
		if err != nil {
			return nil, err
		}

		return artists, nil
	}

	return nil, errors.New("unable to get artist id from arguments")

}

func (r *resolver) CreateArtwork(p graphql.ResolveParams) (interface{}, error) {
	artistId, ok := p.Args["artistId"].(int)
	if ok {
		artwork := &models.Artwork{
			ArtistID: uint(artistId),
			Title:    p.Args["title"].(string),
		}
		db := repo.NewArtworkRepo(r.db)
		err := db.Create(artwork)
		if err != nil {
			return nil, err
		}
		return artwork, nil
	}
	return nil, errors.New("unable to get artist id from arguments")
}

func (r *resolver) UpdateArtwork(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["artworkId"].(int)
	if ok {
		artistId, ok := p.Args["artistId"].(int)
		if ok {
			artwork := &models.Artwork{
				ID:       uint(id),
				ArtistID: uint(artistId),
				Title:    p.Args["title"].(string),
			}
			db := repo.NewArtworkRepo(r.db)
			err := db.Update(artwork)
			if err != nil {
				return nil, err
			}
			return artwork, nil
		}
		return nil, errors.New("unable to get artist id from arguments")
	}
	return nil, errors.New("unable to get artwork id from arguments")
}

func (r *resolver) DeleteArtwork(p graphql.ResolveParams) (interface{}, error) {
	artworkId, ok := p.Args["artworkId"].(int)
	if ok {
		artwork := &models.Artwork{
			ID: uint(artworkId),
		}
		db := repo.NewArtworkRepo(r.db)
		err := db.Delete(artwork)
		if err != nil {
			return nil, err
		}
		return artwork, nil
	}
	return nil, errors.New("unable to get artwork id from arguments")
}
