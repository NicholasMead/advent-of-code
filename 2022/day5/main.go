package main

// Importing packages
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Crate rune
type Stack []Crate
type Cargo []Stack
type Instruction struct {
	count int
	from  int
	to    int
}

// Main function
func main() {
	{
		cargoInput, instructionsInput := ReadInput(os.Args[1])
		instructions := ParseInstructions(instructionsInput)
		cargo := ParseCargo(cargoInput)
		for instruction := range instructions {
			fmt.Println(CargoToString(cargo))
			cargo = ApplyInstructionSimple(cargo, instruction)
		}
		fmt.Println(CargoToString(cargo))
		fmt.Println("Message1: ", ReadMessage(cargo))
	}

	{
		cargoInput, instructionsInput := ReadInput(os.Args[1])
		instructions := ParseInstructions(instructionsInput)
		cargo := ParseCargo(cargoInput)
		for instruction := range instructions {
			fmt.Println(CargoToString(cargo))
			cargo = ApplyInstruction9001(cargo, instruction)
		}
		fmt.Println(CargoToString(cargo))
		fmt.Println("Message2: ", ReadMessage(cargo))
	}

}

// Reads input file
func ReadInput(path string) (<-chan string, <-chan string) {
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
	cargo := make(chan string)
	instructions := make(chan string)
	go func() {
		file := []string{}
		for {
			line, err := reader.ReadString('\n')
			file = append(file, strings.Trim(line, "\n"))

			if err != nil {
				if err == io.EOF {
					break
				} else {
					panic(err)
				}
			}
		}

		cargoLines := []string{}
		instructionsLines := []string{}

		parsingCargo := true
		for _, line := range file {
			switch {
			case strings.TrimSpace(line) == "":
				parsingCargo = false
			case parsingCargo:
				cargoLines = append(cargoLines, line)
			case !parsingCargo:
				instructionsLines = append(instructionsLines, line)
			}
		}
		cargoLines = cargoLines[:len(cargoLines)-1] //no need for stack identifiers

		go func() {
			for _, line := range cargoLines {
				cargo <- line
			}
			close(cargo)
		}()
		go func() {
			for _, line := range instructionsLines {
				instructions <- line
			}
			close(instructions)
		}()
	}()

	return cargo, instructions
}

func ParseCargo(manifest <-chan string) Cargo {
	stacks := []Stack{}

	for row := range manifest {
		r, _ := regexp.Compile("([\\w])|(\\s{3}\\s?)")
		items := r.FindAllString(row, -10)

		for i := len(stacks); i < len(items); i++ {
			stacks = append(stacks, Stack{})
		}

		for i, item := range items {
			if len(item) == 1 {
				stacks[i] = append(stacks[i], Crate(item[0]))
			}
		}
	}

	return stacks
}

func ParseInstructions(instructions <-chan string) <-chan Instruction {
	output := make(chan Instruction)

	go func() {
		defer close(output)
		for row := range instructions {
			r := regexp.MustCompile("move \\d+ from \\d+ to \\d+")
			if r.MatchString(row) {
				r := regexp.MustCompile("\\d+")
				items := r.FindAllString(row, 3)

				count, _ := strconv.Atoi(items[0])
				from, _ := strconv.Atoi(items[1])
				to, _ := strconv.Atoi(items[2])

				output <- Instruction{
					count: count,
					from:  from - 1,
					to:    to - 1,
				}
			}
		}
	}()

	return output
}

func ApplyInstructionSimple(cargo Cargo, instruction Instruction) Cargo {
	from, to, count := instruction.from, instruction.to, instruction.count

	if count <= 0 {
		return cargo
	}

	cargo = ApplyInstruction9001(cargo, Instruction{
		from:  from,
		to:    to,
		count: 1,
	})

	return ApplyInstructionSimple(cargo, Instruction{
		from:  from,
		to:    to,
		count: count - 1,
	})
}

func ApplyInstruction9001(cargo Cargo, instruction Instruction) Cargo {
	from, to, count := instruction.from, instruction.to, instruction.count

	if len(cargo[from]) < count {
		fmt.Println(cargo, instruction, "Stack too small")
		panic(fmt.Sprintf("Stack %d is too small to move %d crates.", from, count))
	}

	moving, remainder := DivideStack(cargo[from], count)
	cargo[from] = remainder
	cargo[to] = append(moving, cargo[to]...)

	return cargo
}

func ReadMessage(cargo Cargo) string {
	message := ""
	for _, stack := range cargo {
		if len(stack) > 0 {
			message += string(stack[0])
		}
	}
	return message
}

func CargoToString(cargo Cargo) string {
	output := "["
	for _, stack := range cargo {
		output += StackToString(stack)
	}
	output += "]"
	return output
}

func StackToString(stack Stack) string {
	output := "["
	for _, crate := range stack {
		output += string(crate)
	}
	output += "]"
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

func DivideStack(stack Stack, position int) (Stack, Stack) {
	a, b := append(stack[:position]), append(stack[position:])
	return Copy(a), Copy(b)
}

func Copy[T any](a []T) []T {
	b := make([]T, len(a))

	for i, v := range a {
		b[i] = v
	}

	return b
}
