package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Minimal bar")

	// create a bar without a stats function
	loadingBar := bar.New(bar.OptionWithoutStatsFuncs())

	// set the total to 100
	loadingBar.SetTotal(100)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		loadingBar.SetCurrent(i)
	}

}
