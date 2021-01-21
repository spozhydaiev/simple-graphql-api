package gql

import (
	"github.com/graphql-go/graphql"
)

var ArtistType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Artist",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"artworks": &graphql.Field{
				Type: graphql.NewList(ArtworkType),
			},
		},
	},
)

var ArtworkType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Artwork",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"artist": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
