package gql

import (
	"fmt"

	"github.com/SergioBravo/simple-graphql-api/repo"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *repo.DB) *Root {
	resolver := Resolver{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"artists": &graphql.Field{
						Type:    graphql.NewList(ArtistType),
						Resolve: resolver.ArtistsResolver,
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
