package processor

import (
	"fmt"
	"net"
	"sync"
)

func Process(urls <-chan string) <-chan int {
	ports := make([]int, 1024)
	result := make(chan int)

	for i := 0; i < 1024; i++ {
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

func process(port int, result chan<- int, urlRaw string, wg *sync.WaitGroup) {
	network := parseURLWithPort(urlRaw, port)

	_, err := net.Dial("tcp", network)
	if err == nil {
		result <- port
	}
	wg.Done()
}

func parseURLWithPort(urlRaw string, port int) string {
	return fmt.Sprintf("%s:%d", urlRaw, port)
}
