package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tf63/code-api/api"      // 修正
	"github.com/tf63/code-api/external" // 修正
	"github.com/tf63/code-api/internal/wire"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := external.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}

	rdb, err := external.ConnectRedis()
	if err != nil {
		log.Fatal("Failed to Connect Database")
	}

	resolver := wire.Wire(db, rdb)
	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
