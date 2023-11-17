package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// LanguageRepository
type LanguageRepository interface {
	ReadLanguages() (languages []entity.Language, err error) // Languageの取得
}
