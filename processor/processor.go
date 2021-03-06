package processor

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const PortRange = 5000

var Scans int64

func getProtocols() []string {
	return []string{"tcp", "udp"}
}

func Process(urls <-chan string) <-chan Result {
	ports := make([]int, PortRange)
	result := make(chan Result)

	for i := 0; i < PortRange; i++ {
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
	conn, err := net.DialTimeout(protocol, address, 100*time.Millisecond)
	atomic.AddInt64(&Scans, 1)
	if err != nil {
		wg.Done()
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	result <- createResult(port, address, protocol)
	wg.Done()
}

func parseURLWithPort(urlRaw string, port int) string {
	return fmt.Sprintf("%s:%d", urlRaw, port)
}
