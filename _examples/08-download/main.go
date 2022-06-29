package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Built-in download bar")

	// create a temp file to save our download to locally
	f, _ := ioutil.TempFile(os.TempDir(), "loading-example")

	// create the bar
	if err := bar.Download("http://speedtest.ftp.otenet.gr/files/test10Mb.db", f); err != nil {
		panic(err)
	}

	fmt.Println("Download complete!")
}
