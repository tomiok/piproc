package main

import (
	"fmt"
	"github.com/tomiok/piproc/processor"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	now := time.Now()

	results := processor.Process(processor.Extract("urls.txt"))
	for result := range results {
		fmt.Printf("port %d open in host: %s \n", result.Port, result.Host)
	}
	elapsed := time.Since(now)

	fmt.Printf("process took %f seconds", elapsed.Seconds())
}

func profile() {
	f, err := os.Create("test.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
