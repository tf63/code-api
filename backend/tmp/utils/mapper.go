package utils

import (
	"errors"
)

func Check(title string, done string, ticketIdx int) error {

	titleList := []string{"todo 1", "todo 2", "todo 3", "todo 4", "todo 5", "todo 6", "todo 7", "todo 8", "todo 9", "todo 10"}
	doneList := []string{"false", "true", "true", "true", "false", "true", "true", "false", "true", "false"}

	idx := ticketIdx - 1

	if title == titleList[idx] && done == doneList[idx] {
		return nil
	} else {
		return errors.New("error")
	}
}
