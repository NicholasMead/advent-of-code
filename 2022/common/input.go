package common

import (
	"bufio"
	"os"
	"strings"
)

func ReadInput(path string) <-chan string {
	var stream *os.File

	switch path {
	case "":
		panic("No input path")
	case "-":
		stream = os.Stdin
	default:
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		stream = file
	}

	reader := bufio.NewReader(stream)
	output := make(chan string)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if line != "" {
				output <- strings.TrimSpace(line)
			}

			if err != nil {
				close(output)
				return
			}
		}
	}()

	return output
}
