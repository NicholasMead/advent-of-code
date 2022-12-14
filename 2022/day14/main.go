package main

import (
	"aco/common"
	c "aoc/day14/cave"
	"aoc/day14/coord"
	"flag"
	"fmt"
	"strings"
)

func main() {
	file := flag.String("f", "-", "input file")
	flag.Parse()

	input := common.ReadInputPath(*file)
	cave := c.Create()

	for line := range input {
		parts := strings.Split(line, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			left, right := coord.FromString(parts[i]), coord.FromString(parts[i+1])
			cave.AddRange(left, right, c.Rock)
		}
	}

	fmt.Println(cave)
	fmt.Println("Starting...")

	hitFloor := false
	count := 0
	for {
		count++
		drop := coord.Coord{X: 500, Y: 0}
		pos, floor := cave.Drop(drop, c.Sand)

		if floor && !hitFloor {
			hitFloor = true
			fmt.Println(cave)
			fmt.Println("part1:", cave.Count(c.Sand)-1)
		} else if pos == drop {
			fmt.Println(cave)
			fmt.Println("part2:", cave.Count(c.Sand))
			break
		} else {
			if count > 0 && count%10000 == 0 {
				fmt.Println("count:", count, "pos:", pos)
			}
		}
	}
}
