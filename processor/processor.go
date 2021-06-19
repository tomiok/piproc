package processor

import (
	"fmt"
	"net"
	"sync"
)

const portRange = 500

func Process(urls <-chan string) <-chan Result {
	ports := make([]int, portRange)
	result := make(chan Result)

	for i := 0; i < portRange; i++ {
		ports[i] = i + 1
	}
	go func() {
		wg := &sync.WaitGroup{}
		for urlRaw := range urls {
			for _, port := range ports {
				wg.Add(1)
				go process(port, result, urlRaw, wg)
			}
		}
		wg.Wait()
		close(result)
	}()
	return result
}

func process(port int, result chan<- Result, urlRaw string, wg *sync.WaitGroup) {
	defer wg.Done()
	network := parseURLWithPort(urlRaw, port)
	if network == "" {
		return
	}
	_, err := net.Dial("tcp", network)
	if err == nil {
		result <- createResult(port, urlRaw)
	}
}

func parseURLWithPort(urlRaw string, port int) string {
	return fmt.Sprintf("%s:%d", urlRaw, port)
}
