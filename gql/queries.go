package gql

import "github.com/graphql-go/graphql"

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"artist": &graphql.Field{
				Type:        ArtistType,
				Description: "Get artist by name",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, ok := p.Args["name"].(string)
					if ok {
						for _, artist := range artists {
							if string(artist.Name) == name {
								return artist, nil
							}
						}
					}
					return nil, nil
				},
			},
			"artist_list": &graphql.Field{
				Type:        graphql.NewList(ArtistType),
				Description: "Get list of artists",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return artists, nil
				},
			},
		},
	},
)
