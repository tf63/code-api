package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// FrameworkRepository
type FrameworkRepository interface {
	ReadFrameworks() (frameworks []entity.Framework, err error) // Frameworkの取得
}
