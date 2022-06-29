package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Bar with a label")

	// create the bar
	loadingBar := bar.New()

	// set the label
	loadingBar.SetLabel("Demoing bar...")

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		if i == 100 {
			loadingBar.SetLabel("Demo complete.")
		}
		loadingBar.SetCurrent(i)
	}
}
