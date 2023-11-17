package resolver

import (
	"github.com/tf63/code-api/pkg/domain/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	fwcr repository.FrameworkCodeRepository
	lgcr repository.LanguageCodeRepository
	ptcr repository.PatternCodeRepository
	arcr repository.AlgorithmCodeRepository
	lgr  repository.LanguageRepository
	arr  repository.AlgorithmRepository
	fwr  repository.FrameworkRepository
	ptr  repository.PatternRepository
}

func NewResolver(fwcr repository.FrameworkCodeRepository,
	ptcr repository.PatternCodeRepository,
	arcr repository.AlgorithmCodeRepository,
	lgcr repository.LanguageCodeRepository,
	lgr repository.LanguageRepository,
	fwr repository.FrameworkRepository,
	arr repository.AlgorithmRepository,
	ptr repository.PatternRepository) Resolver {
	return Resolver{fwcr: fwcr, ptcr: ptcr, arcr: arcr, lgcr: lgcr,
		lgr: lgr, fwr: fwr, arr: arr, ptr: ptr}
}
