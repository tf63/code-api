package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// LanguageCodeRepository
type LanguageCodeRepository interface {
	ReadLanguageCode(findLanguageCode entity.FindLanguageCode) (languageCodes []entity.LanguageCode, err error) // LanguageCodeの取得
}
