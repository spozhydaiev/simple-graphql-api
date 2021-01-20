package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

func main() {
	schema, err := graphql.NewSchema(defineSchema()) // определение схемы рассмотрим  чуть позже
	if err != nil {
		log.Panic("Error when creating the graphQL schema", err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	}) // Здесь же есть интересный параметр FormatErrorFn - функция для форматирования ошибок

	http.Handle("/graphql", h) // путь для доступа к интерфейсу playground и для отправки запросов
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("Error when starting the http server", err)
	}
}
