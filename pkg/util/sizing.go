package util

import (
	"golang.org/x/term"
)

func TermWidth() int {
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	return termWidth
}
