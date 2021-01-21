package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergioBravo/simple-graphql-api/gql"
	"github.com/SergioBravo/simple-graphql-api/migration"
	"github.com/SergioBravo/simple-graphql-api/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize our api and return a pointer to our router for http.ListenAndServe
	// and a pointer to our db to defer its closing when main() is finished
	migration.Migrate()
	router, _ := initializeAPI()

	// Listen on port 4000 and if there's an error log it and exit
	log.Fatal(http.ListenAndServe(":4000", router))

}

func initializeAPI() (*chi.Mux, *gorm.DB) {
	// Create a new router
	router := chi.NewRouter()
	dsn := "host=localhost user=postgres password=abkjcjabz24 dbname=gallery port=5432 sslmode=disable"

	// Create a new connection to our pg database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := gql.NewRoot(db)

	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    rootQuery.Query,
			Mutation: rootQuery.Mutation,
		},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{GqlSchema: &sc}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,       // log api request calls
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}
