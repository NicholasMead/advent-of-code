package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	distance, depth, aim int
}

type Instruction struct {
	direction Direction
	magnitude int
}

type Direction string

const (
	Forward Direction = "forward"
	Up      Direction = "up"
	Down    Direction = "down"
)

func main() {
	lines := streamFileByLine(os.Args[1])
	simpPos, advPos := Position{}, Position{}

	for line := range lines {
		inst := ParseInstruction(line)
		simpPos = simpPos.ApplyInstructionSimp(inst)
		advPos = advPos.ApplyInstructionAdv(inst)
	}

	fmt.Printf("Part 1: (dist: %d, depth: %d) => %d\n", simpPos.distance, simpPos.depth, simpPos.depth*simpPos.distance)
	fmt.Printf("Part 2: (dist: %d, depth: %d, aim: %d) => %d\n", advPos.distance, advPos.depth, advPos.aim, advPos.depth*advPos.distance)
}

func streamFileByLine(path string) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			output <- scanner.Text()
		}
	}()

	return output
}

func ParseInstruction(text string) Instruction {
	part := strings.Split(text, " ")
	dir := Direction(part[0])
	mag, _ := strconv.Atoi(part[1])

	return Instruction{dir, mag}
}

func (p Position) ApplyInstructionSimp(instruction Instruction) Position {
	switch instruction.direction {
	case Forward:
		p.distance += instruction.magnitude
	case Down:
		p.depth += instruction.magnitude
	case Up:
		p.depth -= instruction.magnitude
	}
	return p
}

func (p Position) ApplyInstructionAdv(instruction Instruction) Position {
	switch instruction.direction {
	case Forward:
		p.distance += instruction.magnitude
		p.depth += p.aim * instruction.magnitude
	case Down:
		p.aim += instruction.magnitude
	case Up:
		p.aim -= instruction.magnitude
	}
	return p
}
