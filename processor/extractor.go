package processor

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Extract(location string) <-chan string {
	co := make(chan string)

	go func() {
		f, err := os.Open(location)

		if err != nil {
			panic(err.Error())
		}
		defer f.Close()
		reader := bufio.NewReader(f)

		for {
			line, err := reader.ReadString('\n')

			if err == io.EOF {
				break
			}

			line = strings.Trim(line, "\n")
			line = strings.Trim(line, "\r")

			co <- line
		}
		close(co)
	}()

	return co
}
