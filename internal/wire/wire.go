package wire

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/tf63/code-api/internal/repository"
	"github.com/tf63/code-api/internal/resolver"
)

func InitializeResolver(
	db *sql.DB, rdb *redis.Client,
) resolver.Resolver {
	wire.Build(
		repository.NewAlgorithmCodeRepository,
		repository.NewAlgorithmRepository,
		repository.NewFrameworkCodeRepository,
		repository.NewFrameworkRepository,
		repository.NewLanguageCodeRepository,
		repository.NewLanguageRepository,
		repository.NewPatternCodeRepository,
		repository.NewPatternRepository,
		resolver.NewResolver,
	)

	return resolver.Resolver{}
}
