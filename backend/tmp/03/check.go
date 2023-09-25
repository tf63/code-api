package main

import (
	"os"

	"github.com/tf63/code-api/tmp/utils"
)

func main() {
	title := ""
	done := ""
	ticketIdx := 3

	if err := utils.Check(title, done, ticketIdx); err != nil {
		println("invalid value")
		os.Exit(1)
	}

	println("success")
}
