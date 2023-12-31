// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/pkg/infra/repository"
	"github.com/tf63/code-api/pkg/resolver"
)

// Injectors from wire.go:

func InitializeResolver(db *sql.DB, rdb *redis.Client) resolver.Resolver {
	frameworkCodeRepository := repository.NewFrameworkCodeRepository(db)
	patternCodeRepository := repository.NewPatternCodeRepository(db)
	algorithmCodeRepository := repository.NewAlgorithmCodeRepository(db)
	languageCodeRepository := repository.NewLanguageCodeRepository(db)
	languageRepository := repository.NewLanguageRepository(db, rdb)
	frameworkRepository := repository.NewFrameworkRepository(db, rdb)
	algorithmRepository := repository.NewAlgorithmRepository(db, rdb)
	patternRepository := repository.NewPatternRepository(db, rdb)
	resolverResolver := resolver.NewResolver(frameworkCodeRepository, patternCodeRepository, algorithmCodeRepository, languageCodeRepository, languageRepository, frameworkRepository, algorithmRepository, patternRepository)
	return resolverResolver
}
