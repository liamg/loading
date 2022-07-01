package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Bar with a label")

	// create the bar
	loadingBar := bar.New(
		bar.OptionWithAutoComplete(false),
	)

	// set the label
	loadingBar.SetLabel("Demoing bar...")

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 20)
		loadingBar.SetCurrent(i)
	}
	for i := 100; i >= 50; i-- {
		time.Sleep(time.Millisecond * 20)
		loadingBar.SetCurrent(i)
	}
	for i := 50; i <= 100; i++ {
		time.Sleep(time.Millisecond * 20)
		loadingBar.SetCurrent(i)
	}

	loadingBar.SetLabel("All done.")
	loadingBar.Finish()
}
