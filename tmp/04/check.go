package main

import (
	"os"

	"github.com/tf63/code-api/tmp/utils"
)

func main() {
	title := "todo 4"
	done := "true"
	ticketIdx := 4

	if err := utils.Check(title, done, ticketIdx); err != nil {
		println("invalid value")
		os.Exit(1)
	}

	println("success")
}
