package processor

import (
	"fmt"
	"net"
	"sync"
)

const portRange = 1024

func getProtocols() []string {
	return []string{"tcp", "udp"}
}

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
	address := parseURLWithPort(urlRaw, port)
	if address == "" {
		wg.Done()
		return
	}

	innerWG := &sync.WaitGroup{}
	for _, protocol := range getProtocols() {
		innerWG.Add(1)
		go worker(innerWG, result, address, protocol, port)
	}
	innerWG.Wait()

	wg.Done()
}

func worker(wg *sync.WaitGroup, result chan<- Result, address, protocol string, port int) {
	_, err := net.Dial("tcp", address)
	if err == nil {
		result <- createResult(port, address, protocol)
	}
	wg.Done()
}

func parseURLWithPort(urlRaw string, port int) string {
	return fmt.Sprintf("%s:%d", urlRaw, port)
}
