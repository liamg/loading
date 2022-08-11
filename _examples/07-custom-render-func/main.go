package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("")
	fmt.Println("Demo: This bar has a custom render function")

	// create the bar
	loadingBar := bar.New(
		bar.OptionWithRenderFunc(func(w io.Writer, current, total int) {
			all := strings.Repeat("\x1b[93mThis loading bar is very customised. ", 10)
			_, _ = fmt.Fprint(w, all[:current])
			rem := total - current
			if rem < 0 {
				rem = 0
			}
			_, _ = fmt.Fprint(w, strings.Repeat(" ", rem))
		}),
	)

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		loadingBar.SetCurrent(i)
	}

}
