// Declaration of the main package
package main

// Importing packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type assignment struct {
	lower int
	upper int
}

func (a assignment) FullyOverlaps(b assignment) bool {
	return (a.lower <= b.lower && a.upper >= b.upper)
}

func (a assignment) PartiallyOverlaps(b assignment) bool {
	lower := a.lower <= b.lower && a.upper >= b.lower
	upper := a.upper >= b.upper && a.lower <= b.upper
	return lower || upper
}

// Main function
func main() {
	input := ReadInput(os.Args[1])

	assignments := Split(GetPairAssignment(input), 2)

	//part1
	fullOverlaps := FindFullOverlaps(assignments[0])
	part1 := Count(fullOverlaps)

	//part2
	partiallyOverlaps := FindPartialOverlaps(assignments[1])
	part2 := Count(partiallyOverlaps)

	fmt.Println("Part 1:", <-part1)
	fmt.Println("Part 2:", <-part2)
}

// Reads input file
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

func GetPairAssignment(input <-chan string) <-chan [2]assignment {
	output := make(chan [2]assignment)

	go func() {
		for line := range input {
			assignments := [2]assignment{}
			segments := strings.Split(line, ",")

			for i := 0; i < 2; i++ {
				bounds := strings.Split(segments[i], "-")
				assignments[i].lower, _ = strconv.Atoi(bounds[0])
				assignments[i].upper, _ = strconv.Atoi(bounds[1])
			}

			output <- assignments
		}
		close(output)
	}()

	return output
}

func FindFullOverlaps(pairs <-chan [2]assignment) <-chan [2]assignment {
	output := make(chan [2]assignment)

	go func() {
		for pair := range pairs {
			if pair[0].FullyOverlaps(pair[1]) || pair[1].FullyOverlaps(pair[0]) {
				output <- pair
			}
		}

		close(output)
	}()

	return output
}

func FindPartialOverlaps(pairs <-chan [2]assignment) <-chan [2]assignment {
	output := make(chan [2]assignment)

	go func() {
		for pair := range pairs {
			//is A partially overlaps B, B must partially overlap A 🤯
			if pair[0].PartiallyOverlaps(pair[1]) {
				output <- pair
			}
		}

		close(output)
	}()

	return output
}

func Count[T any](channel <-chan T) <-chan int {
	output := make(chan int)

	go func() {
		count := 0
		for {
			_, open := <-channel
			if open {
				count++
			} else {
				output <- count
				close(output)
				return
			}
		}
	}()

	return output
}

func Split[T any](input <-chan T, count int) []chan T {
	outputs := []chan T{}

	for i := 0; i < count; i++ {
		outputs = append(outputs, make(chan T))
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
