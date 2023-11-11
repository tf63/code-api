package resolver

import (
	"fmt"
	"strconv"

	"github.com/tf63/code-api/api"
	"github.com/tf63/code-api/internal/model"
)

// ----------------------------------------------------------------
// Master
// ----------------------------------------------------------------
func NewLanguageDto(m model.Language) api.Language {
	return api.Language{
		LanguageID: fmt.Sprintf("%d", m.LanguageId),
		Name:       m.Name,
	}
}

func NewFrameworkDto(m model.Framework) api.Framework {
	return api.Framework{
		FrameworkID: fmt.Sprintf("%d", m.FrameworkId),
		Name:        m.Name,
	}
}

func NewAlgorithmDto(m model.Algorithm) api.Algorithm {
	return api.Algorithm{
		AlgorithmID: fmt.Sprintf("%d", m.AlgorithmId),
		Name:        m.Name,
	}
}

func NewPatternDto(m model.Pattern) api.Pattern {
	return api.Pattern{
		PatternID: fmt.Sprintf("%d", m.PatternId),
		Name:      m.Name,
	}
}

// ----------------------------------------------------------------
// FrameworkCode
// ----------------------------------------------------------------
func NewFrameworkCodeDto(m model.FrameworkCode) api.FrameworkCode {
	return api.FrameworkCode{
		FrameworkCodeID: fmt.Sprintf("%d", m.FrameworkCodeId),
		Content:         m.Content,
		Nrow:            m.Nrow,
		CreatedAt:       m.CreatedAt,
		ToolID:          fmt.Sprintf("%d", m.ToolId),
	}
}

func NewFindFrameworkCodeFromDto(input api.FindFrameworkCode) (*model.FindFrameworkCode, error) {
	findFrameworkCode := model.FindFrameworkCode{}

	if input.Offset != nil {
		findFrameworkCode.Offset = *input.Offset
	}

	if input.Limit != nil {
		findFrameworkCode.Limit = *input.Limit
	}

	if input.StartRow != nil {
		findFrameworkCode.StartRow = *input.StartRow
	}

	if input.EndRow != nil {
		findFrameworkCode.EndRow = *input.EndRow
	}

	toolId, err := strconv.Atoi(input.ToolID)
	if err != nil {
		return nil, err
	}
	findFrameworkCode.ToolId = toolId

	return &findFrameworkCode, nil
}

// ----------------------------------------------------------------
// PatternCode
// ----------------------------------------------------------------
func NewPatternCodeDto(m model.PatternCode) api.PatternCode {
	return api.PatternCode{
		PatternCodeID: fmt.Sprintf("%d", m.PatternCodeId),
		Content:       m.Content,
		Nrow:          m.Nrow,
		CreatedAt:     m.CreatedAt,
		LanguageID:    fmt.Sprintf("%d", m.LanguageId),
		PatternID:     fmt.Sprintf("%d", m.PatternId),
	}
}

func NewFindPatternCodeFromDto(input api.FindPatternCode) (*model.FindPatternCode, error) {
	findPatternCode := model.FindPatternCode{}

	if input.Offset != nil {
		findPatternCode.Offset = *input.Offset
	}

	if input.Limit != nil {
		findPatternCode.Limit = *input.Limit
	}

	if input.StartRow != nil {
		findPatternCode.StartRow = *input.StartRow
	}

	if input.EndRow != nil {
		findPatternCode.EndRow = *input.EndRow
	}

	languageId, err := strconv.Atoi(input.LanguageID)
	if err != nil {
		return nil, err
	}
	findPatternCode.LanguageId = languageId

	patternId, err := strconv.Atoi(input.PatternID)
	if err != nil {
		return nil, err
	}
	findPatternCode.PatternId = patternId

	return &findPatternCode, nil
}

// ----------------------------------------------------------------
// AlgorithmCode
// ----------------------------------------------------------------
func NewAlgorithmCodeDto(m model.AlgorithmCode) api.AlgorithmCode {
	return api.AlgorithmCode{
		AlgorithmCodeID: fmt.Sprintf("%d", m.AlgorithmCodeId),
		Content:         m.Content,
		Nrow:            m.Nrow,
		CreatedAt:       m.CreatedAt,
		LanguageID:      fmt.Sprintf("%d", m.LanguageId),
		AlgorithmID:     fmt.Sprintf("%d", m.AlgorithmId),
	}
}

func NewFindAlgorithmCodeFromDto(input api.FindAlgorithmCode) (*model.FindAlgorithmCode, error) {
	findAlgorithmCode := model.FindAlgorithmCode{}

	if input.Offset != nil {
		findAlgorithmCode.Offset = *input.Offset
	}

	if input.Limit != nil {
		findAlgorithmCode.Limit = *input.Limit
	}

	if input.StartRow != nil {
		findAlgorithmCode.StartRow = *input.StartRow
	}

	if input.EndRow != nil {
		findAlgorithmCode.EndRow = *input.EndRow
	}

	languageId, err := strconv.Atoi(input.LanguageID)
	if err != nil {
		return nil, err
	}
	findAlgorithmCode.LanguageId = languageId

	algorithmId, err := strconv.Atoi(input.AlgorithmID)
	if err != nil {
		return nil, err
	}
	findAlgorithmCode.AlgorithmId = algorithmId

	return &findAlgorithmCode, nil
}

// ----------------------------------------------------------------
// LanguageCode
// ----------------------------------------------------------------
func NewLanguageCodeDto(m model.LanguageCode) api.LanguageCode {
	return api.LanguageCode{
		LanguageCodeID: fmt.Sprintf("%d", m.LanguageCodeId),
		Content:        m.Content,
		Nrow:           m.Nrow,
		CreatedAt:      m.CreatedAt,
		LanguageID:     fmt.Sprintf("%d", m.LanguageId),
	}
}

func NewFindLanguageCodeFromDto(input api.FindLanguageCode) (*model.FindLanguageCode, error) {
	findLanguageCode := model.FindLanguageCode{}

	if input.Offset != nil {
		findLanguageCode.Offset = *input.Offset
	}

	if input.Limit != nil {
		findLanguageCode.Limit = *input.Limit
	}

	if input.StartRow != nil {
		findLanguageCode.StartRow = *input.StartRow
	}

	if input.EndRow != nil {
		findLanguageCode.EndRow = *input.EndRow
	}

	languageId, err := strconv.Atoi(input.LanguageID)
	if err != nil {
		return nil, err
	}
	findLanguageCode.LanguageId = languageId

	return &findLanguageCode, nil
}
