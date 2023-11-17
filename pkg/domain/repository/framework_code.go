package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// FrameworkCodeRepository
type FrameworkCodeRepository interface {
	ReadFrameworkCode(findFrameworkCode entity.FindFrameworkCode) (frameworkCodes []entity.FrameworkCode, err error) // FrameworkCodeの取得
}
