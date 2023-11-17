package resolver

import (
	"github.com/tf63/code-api/pkg/domain/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Fwcr repository.FrameworkCodeRepository
	Lgcr repository.LanguageCodeRepository
	Ptcr repository.PatternCodeRepository
	Arcr repository.AlgorithmCodeRepository
	Lgr  repository.LanguageRepository
	Arr  repository.AlgorithmRepository
	Fwr  repository.FrameworkRepository
	Ptr  repository.PatternRepository
}

func NewResolver(fwcr repository.FrameworkCodeRepository,
	ptcr repository.PatternCodeRepository,
	arcr repository.AlgorithmCodeRepository,
	lgcr repository.LanguageCodeRepository,
	lgr repository.LanguageRepository,
	fwr repository.FrameworkRepository,
	arr repository.AlgorithmRepository,
	ptr repository.PatternRepository) Resolver {
	return Resolver{Fwcr: fwcr, Ptcr: ptcr, Arcr: arcr, Lgcr: lgcr,
		Lgr: lgr, Fwr: fwr, Arr: arr, Ptr: ptr}
}
