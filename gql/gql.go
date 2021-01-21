package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Root struct {
	Query    *graphql.Object
	Mutation *graphql.Object
}

func NewRoot(db *gorm.DB) *Root {
	resolver := NewResolver(db)

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"artistList": &graphql.Field{
						Type:    graphql.NewList(ArtistType),
						Resolve: resolver.GetArtistList,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
					},
					"artistByName": &graphql.Field{
						Type:    ArtistType,
						Resolve: resolver.GetArtistByName,
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
					},
					"artistArtworks": &graphql.Field{
						Type:    graphql.NewList(ArtworkType),
						Resolve: resolver.GetArtistArtworks,
						Args: graphql.FieldConfigArgument{
							"artistId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
					},
				},
			},
		),
		Mutation: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Mutation",
				Fields: graphql.Fields{
					"createArtist": &graphql.Field{
						Type:        ArtistType,
						Description: "Create new artist",
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.CreateArtist,
					},
					"updateArtist": &graphql.Field{
						Type:        ArtistType,
						Description: "Update artist",
						Args: graphql.FieldConfigArgument{
							"artistId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"name": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
						},
						Resolve: resolver.UpdateArtist,
					},
					"deleteArtist": &graphql.Field{
						Type:        ArtistType,
						Description: "Delete artist",
						Args: graphql.FieldConfigArgument{
							"artistId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteArtist,
					},
					"createArtwork": &graphql.Field{
						Type:        ArtworkType,
						Description: "Create new artwork",
						Args: graphql.FieldConfigArgument{
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"artistId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.CreateArtwork,
					},
					"updateArtwork": &graphql.Field{
						Type:        ArtworkType,
						Description: "Update artwork",
						Args: graphql.FieldConfigArgument{
							"artworkId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
							"title": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							"artistId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.UpdateArtwork,
					},
					"deleteArtwork": &graphql.Field{
						Type:        ArtworkType,
						Description: "Delete artwork",
						Args: graphql.FieldConfigArgument{
							"artworkId": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.Int),
							},
						},
						Resolve: resolver.DeleteArtwork,
					},
				},
			},
		),
	}
	return &root
}

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}
