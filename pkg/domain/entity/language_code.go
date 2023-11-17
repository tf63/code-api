package entity

import "time"

type LanguageCode struct {
	LanguageCodeId int
	Content        string
	Nrow           int
	LanguageId     int
	CreatedAt      time.Time
}

type FindLanguageCode struct {
	LanguageId int
	StartRow   int
	EndRow     int
	Offset     int
	Limit      int
}
