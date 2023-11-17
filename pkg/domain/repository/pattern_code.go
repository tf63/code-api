package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// PatternCodeRepository
type PatternCodeRepository interface {
	ReadPatternCode(findPatternCode entity.FindPatternCode) (patternCodes []entity.PatternCode, err error) // PatternCodeの取得
}
