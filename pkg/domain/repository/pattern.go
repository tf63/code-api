package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// PatternRepository
type PatternRepository interface {
	ReadPatterns() (patterns []entity.Pattern, err error) // Patternの取得
}
