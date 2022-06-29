package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Custom characters")

	// create the bar
	loadingBar := bar.New(
		bar.OptionWithRenderFunc(
			bar.RenderCustomCharacters('=', '_'),
		),
	)

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		loadingBar.SetCurrent(i)
	}

}
