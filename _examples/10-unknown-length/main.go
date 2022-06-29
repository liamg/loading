package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Bar with unknown total")

	// create the bar
	loadingBar := bar.New()

	// set the total to an unknown amount
	loadingBar.SetTotal(0)

	// increment the bar to 100 over several seconds
	for i := 0; i <= 100; i++ {
		time.Sleep(time.Millisecond * 30)
		loadingBar.SetCurrent(i)
	}

	// we should finish the bar afterwards, as with no total, it cannot otherwise know when it is finished
	loadingBar.Finish()

}
