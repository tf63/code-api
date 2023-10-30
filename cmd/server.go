package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tf63/code-api/api" // 修正
	"github.com/tf63/code-api/external"
	"github.com/tf63/code-api/internal/repository"
	"github.com/tf63/code-api/internal/resolver" // 修正
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

	pgcr := repository.NewProgramCodeRepository(db)
	ptcr := repository.NewPatternCodeRepository(db)
	arcr := repository.NewAlgorithmCodeRepository(db)
	lgr := repository.NewLanguageRepository(db, rdb)
	fwr := repository.NewFrameworkRepository(db, rdb)
	arr := repository.NewAlgorithmRepository(db, rdb)
	ptr := repository.NewPatternRepository(db, rdb)

	resolver := resolver.Resolver{Pgcr: pgcr, Ptcr: ptcr, Arcr: arcr,
		Lgr: lgr, Fwr: fwr, Arr: arr, Ptr: ptr}
	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
