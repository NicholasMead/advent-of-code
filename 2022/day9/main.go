package main

import (
	"aco/common"
	"fmt"
	"oac/day9/coord"
	"oac/day9/rope"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := common.Split(common.ReadInput(os.Args[1]), 2)

	p1, p2 := ParseRopeTail(input[0], 2), ParseRopeTail(input[1], 10)

	fmt.Println("Part1:", <-p1)
	fmt.Println("Part2:", <-p2)
}

func ParseRopeTail(input <-chan string, size int) <-chan int {
	output := make(chan int, 1)
	rope := rope.CreateRope(size)
	visited := map[coord.Coord]struct{}{}

	go func() {
		for input := range input {
			parsed := strings.Split(input, " ")
			dir := rune(parsed[0][0])
			count, _ := strconv.Atoi(parsed[1])

			for i := 0; i < count; i++ {
				rope.MoveDirection(dir)
				visited[rope.Tail()] = struct{}{}
			}
		}
		output <- len(visited)
		close(output)
	}()

	return output
}
