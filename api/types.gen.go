// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"time"
)

type Algorithm struct {
	AlgorithmID string `json:"algorithmId"`
	Name        string `json:"name"`
}

type AlgorithmCode struct {
	AlgorithmCodeID string    `json:"algorithmCodeId"`
	Content         string    `json:"content"`
	Nrow            int       `json:"nrow"`
	CreatedAt       time.Time `json:"createdAt"`
	LanguageID      string    `json:"languageId"`
	AlgorithmID     string    `json:"algorithmId"`
}

type FindAlgorithmCode struct {
	LanguageID  string `json:"languageId"`
	AlgorithmID string `json:"algorithmId"`
	StartRow    *int   `json:"startRow,omitempty"`
	EndRow      *int   `json:"endRow,omitempty"`
	Offset      *int   `json:"offset,omitempty"`
	Limit       *int   `json:"limit,omitempty"`
}

type FindFrameworkCode struct {
	ToolID   string `json:"toolId"`
	StartRow *int   `json:"startRow,omitempty"`
	EndRow   *int   `json:"endRow,omitempty"`
	Offset   *int   `json:"offset,omitempty"`
	Limit    *int   `json:"limit,omitempty"`
}

type FindLanguageCode struct {
	LanguageID string `json:"languageId"`
	StartRow   *int   `json:"startRow,omitempty"`
	EndRow     *int   `json:"endRow,omitempty"`
	Offset     *int   `json:"offset,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
}

type FindPatternCode struct {
	LanguageID string `json:"languageId"`
	PatternID  string `json:"patternId"`
	StartRow   *int   `json:"startRow,omitempty"`
	EndRow     *int   `json:"endRow,omitempty"`
	Offset     *int   `json:"offset,omitempty"`
	Limit      *int   `json:"limit,omitempty"`
}

type Framework struct {
	FrameworkID string `json:"frameworkId"`
	Name        string `json:"name"`
}

type FrameworkCode struct {
	FrameworkCodeID string    `json:"frameworkCodeId"`
	Content         string    `json:"content"`
	Nrow            int       `json:"nrow"`
	CreatedAt       time.Time `json:"createdAt"`
	ToolID          string    `json:"toolId"`
}

type Language struct {
	LanguageID string `json:"languageId"`
	Name       string `json:"name"`
}

type LanguageCode struct {
	LanguageCodeID string    `json:"languageCodeId"`
	Content        string    `json:"content"`
	Nrow           int       `json:"nrow"`
	CreatedAt      time.Time `json:"createdAt"`
	LanguageID     string    `json:"languageId"`
}

type NewAlgorithmCode struct {
	LanguageID  string `json:"languageId"`
	AlgorithmID string `json:"algorithmId"`
	Content     string `json:"content"`
}

type NewFrameworkCode struct {
	ToolID  string `json:"toolId"`
	Content string `json:"content"`
}

type NewLanguageCode struct {
	LanguageID string `json:"languageId"`
	Content    string `json:"content"`
}

type NewPatternCode struct {
	LanguageID string `json:"languageId"`
	PatternID  string `json:"patternId"`
	Content    string `json:"content"`
}

type Pattern struct {
	PatternID string `json:"patternId"`
	Name      string `json:"name"`
}

type PatternCode struct {
	PatternCodeID string    `json:"patternCodeId"`
	Content       string    `json:"content"`
	Nrow          int       `json:"nrow"`
	CreatedAt     time.Time `json:"createdAt"`
	LanguageID    string    `json:"languageId"`
	PatternID     string    `json:"patternId"`
}
