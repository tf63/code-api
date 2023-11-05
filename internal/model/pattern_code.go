package model

import "time"

type PatternCode struct {
	PatternCodeId int
	Content       string
	Nrow          int
	PatternId     int
	LanguageId    int
	CreatedAt     time.Time
}

type FindPatternCode struct {
	PatternId  int
	LanguageId int
	StartRow   int
	EndRow     int
	Offset     int
	Limit      int
}
