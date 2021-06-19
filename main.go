package main

import (
	"fmt"
	"github.com/tomiok/piproc/processor"
)

func main() {
	co := processor.Extract("urls.txt")
	results := processor.Process(co)


	for port := range results {
		fmt.Printf("%d open\n", port)
	}
}
