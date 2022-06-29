package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Rainbow bar")

	// create the bar with the rainbow renderer
	loadingBar := bar.New(
		bar.OptionWithRenderFunc(bar.RenderRainbow),
	)

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		loadingBar.SetCurrent(i)
	}

}
