package entity

import "time"

type AlgorithmCode struct {
	AlgorithmCodeId int
	Content         string
	Nrow            int
	AlgorithmId     int
	LanguageId      int
	CreatedAt       time.Time
}

type FindAlgorithmCode struct {
	AlgorithmId int
	LanguageId  int
	StartRow    int
	EndRow      int
	Offset      int
	Limit       int
}
