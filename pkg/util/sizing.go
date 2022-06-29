package util

import "golang.org/x/crypto/ssh/terminal"

func TermWidth() int {
	termWidth, _, err := terminal.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	return termWidth
}
