package main

import (
	"aco/common"
	"fmt"
)

func main() {
	common.TimedMs(func() {
		parseInput()

		p1, p2 := part1(), part2()

		fmt.Println("part1:", p1)
		fmt.Println("part2:", p2)
	})
}

func part1() int {
	return 0
}

func part2() int {
	return 0
}

func parseInput() any {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	return input
}
