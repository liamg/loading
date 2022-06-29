package main

import (
	"fmt"
	"time"

	"github.com/liamg/loading/pkg/bar"
)

func main() {

	fmt.Println("Demo: Nested bars with logging")

	// create the bar
	parent := bar.New(
		bar.OptionWithRenderFunc(
			bar.RenderColoured(255, 102, 0),
		),
	)

	parent.SetLabel("Scanning AWS account...")

	services := []string{
		"cloudfront",
		"ec2",
		"ecs",
		"eks",
		"rds",
		"s3",
	}
	parent.SetTotal(len(services))

	for _, service := range services {
		parent.SetLabel(fmt.Sprintf("Scanning %s...", service))
		child := bar.New(
			bar.OptionHideOnFinish(true),
			bar.OptionWithPadding(2),
		).SetTotal(100)
		for i := 0; i <= 100; i++ {
			time.Sleep(time.Millisecond * 30)
			child.SetLabel(fmt.Sprintf("Scanning resource %02d...", i))
			child.Increment()
		}
		parent.Log("Scanned %s", service)
		parent.Increment()
	}

}
