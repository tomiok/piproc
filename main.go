package main

import (
	"fmt"
	"github.com/tomiok/piproc/processor"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	profiling := os.Getenv("PROFILING")

	b, _ := strconv.ParseBool(profiling)

	if b {
		profile()
	}

	results := processor.Process(processor.Extract("urls.txt"))

	for result := range results {
		fmt.Printf("port %d open in host: %s with protocol %s \n",
			result.Port,
			result.Host,
			result.Protocol,
		)
	}

	elapsed := time.Since(now)
	fmt.Printf("process took %f seconds", elapsed.Seconds())
}

func profile() {
	f, err := os.Create("test.prof")
	if err != nil {
		log.Fatal(err)
	}
	_ = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
}
