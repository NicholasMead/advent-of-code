package common

import (
	"bufio"
	"io/fs"
	"os"
	"strings"
)

func ReadInputPath(path string) <-chan string {
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
	return ReadInputBuffer(reader)
}

func ReadFullInputPath(path string) []string {
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

	reader := bufio.NewScanner(stream)
	input := []string{}

	for reader.Scan() {
		line := reader.Text()
		input = append(input, line)
	}

	return input
}

func ReadInputEmbed(file fs.File) <-chan string {
	reader := bufio.NewReader(file)
	return ReadInputBuffer(reader)
}

func ReadInputBuffer(reader *bufio.Reader) <-chan string {
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
