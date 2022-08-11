package bar

import (
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	runeFullBox  = '█'
	runeEmptyBox = '▒'
)

// RenderFunc is a function that renders the graphical portion of the progress bar
type RenderFunc func(w io.Writer, current, total int)

// RenderSimple renders a simple rectangular progress bar using the default terminal colours
func RenderSimple(w io.Writer, current, total int) {
	RenderCustomCharacters(runeFullBox, runeEmptyBox)(w, current, total)
}

// RenderColored is an alias for RenderColoured
var RenderColored = RenderColoured

// RenderColoured renders a progress bar in the given rgb colour
func RenderColoured(r, g, b int) RenderFunc {
	return func(w io.Writer, current, total int) {
		_, _ = w.Write([]byte(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)))
		RenderSimple(w, current, total)
	}
}

// RenderCustomCharacters renders a simple rectangular progress bar using the supplied characters
func RenderCustomCharacters(complete rune, incomplete rune) RenderFunc {
	return func(w io.Writer, current, total int) {
		_, _ = w.Write([]byte(strings.Repeat(string(complete), current)))
		if total-current > 0 {
			_, _ = w.Write([]byte(strings.Repeat(string(incomplete), total-current)))
		}
	}
}

// RenderRainbow renders a rectangular progress bar that cycles through colours as it progresses across the terminal
func RenderRainbow(w io.Writer, current, total int) {
	for i := 0; i < current; i++ {
		fraction := float64(i) / float64(total)
		red := int((1 - fraction) * 255)
		green := int((2 * (0.5 - math.Abs(0.5-fraction))) * 255)
		blue := int((fraction) * 255)
		_, _ = w.Write([]byte(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c", red, green, blue, runeFullBox)))
	}
	_, _ = w.Write([]byte(strings.Repeat(" ", total-current)))
}
