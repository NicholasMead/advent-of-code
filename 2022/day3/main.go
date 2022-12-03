// Declaration of the main package
package main

// Importing packages
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type rucksack struct {
	compartment1 string
	compartment2 string
}

// Main function
func main() {
	start := time.Now().UnixMicro()

	if len(os.Args) < 2 {
		panic("Input paramater required")
	}

	//Pre Work
	input := ReadInput(os.Args[1])

	//Split
	c0, c1, c2 := make(chan string, 1024), make(chan string, 1024), make(chan string, 1024)
	Split(input, c0, c1, c2)

	//Count
	rucksackCount := Count(c0)

	//Challange 1
	var challange1 <-chan int
	{
		rucksacks := SplitRucksackByCompartment(c1)
		duplicates := FindCommonItems(rucksacks)
		priorities := ConvertToPriority(duplicates)
		challange1 = Sum(priorities)
	}

	//Challange 2
	var challange2 <-chan int
	{
		groups := Group(c2, 3)
		groupItems := FindGroupItem(groups)
		priorities := ConvertToPriority(groupItems)
		challange2 = Sum(priorities)
	}

	fmt.Printf("\t  Rucksacks: %d\n\tChallange 1: %d\n\tChallange 2: %d\n", <-rucksackCount, <-challange1, <-challange2)
	end := time.Now().UnixMicro()
	fmt.Printf(" Elapased time (us): %d\n", end-start)
}

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

func SplitRucksackByCompartment(input <-chan string) <-chan rucksack {
	rucksacks := make(chan rucksack)

	go func() {
		for contence := range input {
			length := len(contence)
			if (length % 2) != 0 {
				panic(fmt.Sprint("Cannot handle rucksack with odd numbered contense:", contence))
			}

			rucksacks <- rucksack{
				compartment1: contence[:length/2],
				compartment2: contence[length/2:],
			}
		}
		close(rucksacks)
	}()

	return rucksacks
}

func FindCommonItems(rucksacks <-chan rucksack) <-chan rune {
	duplicates := make(chan rune)

	go func() {
		for rucksack := range rucksacks {
			items := map[rune]struct{}{}

			for _, item := range rucksack.compartment1 {
				if strings.ContainsRune(rucksack.compartment2, item) {
					items[item] = struct{}{}
				}
			}

			for item := range items {
				duplicates <- item
			}
		}
		close(duplicates)
	}()

	return duplicates
}

func FindGroupItem(groups <-chan []string) <-chan rune {
	groupItem := make(chan rune)

	go func() {
		for group := range groups {
			for _, item := range group[0] {
				found := true
				for _, rucksack := range group[1:] {
					if !strings.ContainsRune(rucksack, item) {
						found = false
					}
				}
				if found {
					groupItem <- item
					break
				}
			}
		}
		close(groupItem)
	}()

	return groupItem
}

func Group[T any](items <-chan T, size int) <-chan []T {
	groups := make(chan []T)

	go func() {
		for {
			group := []T{}
			for i := 0; i < size; i++ {
				item, next := <-items
				if !next {
					close(groups)
					return
				}
				group = append(group, item)
			}
			groups <- group
		}
	}()

	return groups
}

func ConvertToPriority(items <-chan rune) <-chan int {
	priority := make(chan int)

	go func() {
		for item := range items {
			switch {
			case item >= 'a' && item <= 'z':
				priority <- int(item-'a') + 1
			case item >= 'A' && item <= 'Z':
				priority <- int(item-'A') + 27
			default:
				panic(fmt.Sprint("Unknown item:", item))
			}
		}
		close(priority)
	}()

	return priority
}

func Sum(values <-chan int) <-chan int {
	output := make(chan int, 1)

	go func() {
		sum := 0

		for value := range values {
			sum += value
		}

		output <- sum
		close(output)
	}()

	return output
}

func Split[T any](input <-chan T, outputs ...chan<- T) {
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
}

func Terminate[T any](channel <-chan T) {
	go func() {
		for {
			_, open := <-channel
			if !open {
				return
			}
		}
	}()
}

func Count[T any](channel <-chan T) <-chan int {
	output := make(chan int, 1)

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
