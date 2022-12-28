package main

import (
	"aco/common"
	"aoc/day25/snafu"
	"fmt"
)

func main() {
	common.TimedMs(func() {
		input := parseInput()

		p1, p2 := part1(input), part2()

		fmt.Println("part1:", p1)
		fmt.Println("part2:", p2)
	})
}

func part1(input []string) string {
	ans := 0
	for _, i := range input {
		ans += snafu.StoI(i)
	}
	return snafu.ItoS(ans)
}

func part2() int {
	return 0
}

func parseInput() []string {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	values := []string{}
	for i := range input {
		values = append(values, i)
	}

	return values
}
