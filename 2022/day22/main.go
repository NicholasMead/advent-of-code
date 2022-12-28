package main

import (
	"aco/common"
	"aoc/day22/grid"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	common.TimedMs(func() {
		var (
			surface, instructions = parseInput()
			flat                  = grid.Flat(surface)
			cube                  = grid.Cube(surface, 50)
			startX, _             = surface.RowBounds(0)
			start                 = grid.Vector{startX, 0}
			p1, d1                = start, grid.Right
			p2, d2                = start, grid.Right
		)

		fmt.Printf("0/%v", len(instructions))

		for c, i := range instructions {
			if c%100 == 0 {
				fmt.Printf("\r%v/%v", c, len(instructions))
			}
			count, err := strconv.Atoi(i)

			if err == nil {
				for n := 0; n < count; n++ {
					p1, d1 = flat.Move(p1, d1)
					p2, d2 = cube.Move(p2, d2)
				}
			} else {
				r := grid.Rotation(i)
				d1 = d1.Rotate(r)
				d2 = d2.Rotate(r)
			}
		}
		fmt.Printf("\rCompleted: %v\n", len(instructions))

		fmt.Printf("part1: %v (%v)\n", grid.Score(p1, d1), p1)
		fmt.Printf("part1: %v (%v)\n", grid.Score(p2, d2), p2)
	})
}

func parseInput() (grid.Surface, []string) {
	flags := common.GetFlags()
	input := common.ReadFullInputPath(*flags.File)

	var (
		surface      = grid.Surface{}
		instructions []string
	)

	for y, line := range input[0 : len(input)-1] {
		for x, c := range line {
			if c != ' ' {
				surface[grid.Vector{x, y}] = grid.Tile(c)
			}
		}
	}

	instExp := regexp.MustCompile(`\d+|[RL]`)
	instructions = instExp.FindAllString(input[len(input)-1], -1)

	return surface, instructions
}
