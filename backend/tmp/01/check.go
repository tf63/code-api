package main

import (
	"os"

	"github.com/tf63/code-api/tmp/utils"
)

func main() {
	title := "todo 1"
	done := "false"
	ticketIdx := 1

	if err := utils.Check(title, done, ticketIdx); err != nil {
		println("invalid value")
		os.Exit(1)
	}

	println("success")
}
