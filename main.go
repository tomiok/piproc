package main

import (
	"fmt"
	"github.com/tomiok/piproc/processor"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	results := processor.Process(processor.Extract("urls.txt"))

	for result := range results {
		fmt.Printf("%d open in host: %s \n", result.Port, result.Host)
	}
}

func profile() {
	f, err := os.Create("test.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
