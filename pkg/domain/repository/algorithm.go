package repository

import (
	"github.com/tf63/code-api/pkg/domain/entity"
)

// AlgorithmRepository
type AlgorithmRepository interface {
	ReadAlgorithms() (algorithms []entity.Algorithm, err error) // Algorithmの取得
}
