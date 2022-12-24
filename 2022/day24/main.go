package main

import (
	"aco/common"
	"aoc/day24/mountain"
	mtn "aoc/day24/mountain"
	"fmt"
)

func main() {
	common.TimedMs(func() {
		input := parseInput()

		p1, p2 := part1(input), part2(input)

		fmt.Println("part1:", p1)
		fmt.Println("part2:", p2)
	})
}

func part1(m mtn.Mountain) int {
	var (
		start = mtn.Coord{0, -1}
		end   = mtn.Coord{m.Width - 1, m.Height}
	)
	return mtn.FindShortestPath(m, start, end, 0)
}

func part2(m mtn.Mountain) int {
	var (
		start = mtn.Coord{0, -1}
		end   = mtn.Coord{m.Width - 1, m.Height}
	)
	leg1 := mtn.FindShortestPath(m, start, end, 0)
	leg2 := mtn.FindShortestPath(m, end, start, leg1)
	leg3 := mtn.FindShortestPath(m, start, end, leg2)
	return leg3
}

func parseInput() mountain.Mountain {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	raw := [][]rune{}
	for line := range input {
		if line[2] == '#' {
			continue //Throw Away the first and last line
		} else {
			var rawLine = []rune{}
			for _, tile := range line[1 : len(line)-1] {
				rawLine = append(rawLine, tile)
			}
			raw = append(raw, rawLine)
		}
	}

	return mtn.ParseMountain(raw)
}
