package repository

import "github.com/tf63/code-api/pkg/domain/entity"

// AlgorithmCodeRepository
type AlgorithmCodeRepository interface {
	ReadAlgorithmCode(findAlgorithmCode entity.FindAlgorithmCode) (algorithmCodes []entity.AlgorithmCode, err error) // AlgorithmCodeの取得
}
