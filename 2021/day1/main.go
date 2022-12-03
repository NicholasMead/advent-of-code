package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := SplitWithBuffer(ReadInput(), 2, 1024)

	part1 := CalculateIncreases(input[0])
	part2 := CalculateIncreases(MovingSum(input[1], 3))

	fmt.Println("Part 1:", <-part1)
	fmt.Println("Part 2:", <-part2)
}

func ReadInput() <-chan int {
	reader := bufio.NewReader(os.Stdin)
	output := make(chan int)

	go func() {
		for {
			line, end := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if line == "" {
				if end != nil {
					close(output)
					return
				}
				continue
			}

			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			output <- val
		}
	}()

	return output
}

func CalculateIncreases(input <-chan int) <-chan int {
	output := make(chan int)

	go func() {
		prev, increase := -1, 0

		for val := range input {
			switch {
			case prev == -1:
				prev = val
			case val > prev:
				increase++
			}
			prev = val
		}

		output <- increase
		close(output)
	}()

	return output
}

func MovingSum(input <-chan int, windowSize int) <-chan int {
	output := make(chan int)

	go func() {
		buffer := []int{0}

		for val := range input {
			for i := 0; i < len(buffer); i++ {
				buffer[i] += val
			}

			if len(buffer) == windowSize {
				output <- buffer[0]
				buffer = append(buffer[1:], 0)
			} else {
				buffer = append(buffer, 0)
			}
		}

		close(output)
	}()

	return output
}

func SplitWithBuffer[T any](input <-chan T, count int, buffer int) []chan T {
	outputs := []chan T{}

	for i := 0; i < count; i++ {
		outputs = append(outputs, make(chan T, buffer))
	}

	go func() {
		for item := range input {
			for _, output := range outputs {
				output <- item
			}
		}

		for _, output := range outputs {
			close(output)
		}
	}()

	return outputs
}
