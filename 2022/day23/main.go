package main

import (
	"aco/common"
	e "aoc/day23/elfs"
	"fmt"
	"math"
)

func main() {
	common.TimedMs(func() {
		input := parseInput()

		p1, p2 := part1(input), part2(input)

		fmt.Println("part1:", p1)
		fmt.Println("part2:", p2)
	})
}

func part1(input []e.Coord) int {
	elfs, _ := e.Run(input, 10)

	minX, maxX, minY, maxY := e.Bounds(elfs)
	area := (1 + maxX - minX) * (1 + maxY - minY)
	return area - len(input)
}

func part2(input []e.Coord) int {
	_, rounds := e.Run(input, math.MaxInt)

	return rounds
}

func parseInput() []e.Coord {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	all := []e.Coord{}
	y := 0
	for line := range input {
		for x, char := range line {
			if char == '#' {
				all = append(all, e.Coord{X: x, Y: y})
			}
		}
		y--
	}

	return all
}
