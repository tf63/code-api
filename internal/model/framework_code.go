package model

import "time"

type FrameworkCode struct {
	FrameworkCodeId int
	Content         string
	Nrow            int
	ToolId          int
	CreatedAt       time.Time
}

type FindFrameworkCode struct {
	ToolId   int
	StartRow int
	EndRow   int
	Offset   int
	Limit    int
}
