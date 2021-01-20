package gql

import (
	"github.com/SergioBravo/simple-graphql-api/repo"
	"github.com/graphql-go/graphql"
)

type ArtistResolver interface {
	GetArtist(p graphql.ResolveParams) interface{}
}

type resolver struct {
	db *repo.ArtistRepo
}

func NewResolver(db *repo.ArtistRepo) ArtistResolver {
	return &resolver{
		db: db,
	}
}

func (r *resolver) GetArtist(p graphql.ResolveParams) interface{} {
	artists := r.db.GetArtist()
	return artists
}
