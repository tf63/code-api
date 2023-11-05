package model

import "time"

type ProgramCode struct {
	ProgramCodeId int
	Content       string
	Nrow          int
	ToolId        int
	CreatedAt     time.Time
}

type FindProgramCode struct {
	ToolId   int
	StartRow int
	EndRow   int
	Offset   int
	Limit    int
}
