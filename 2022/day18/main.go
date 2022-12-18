package main

import (
	"aco/common"
	"aoc/day18/lava"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	cubes := getInputCubes()
	droplet := lava.FormDroplet(cubes)
	tot, ext := droplet.SurfaceArea()
	fmt.Println("Part1:", tot)
	fmt.Println("Part2:", ext)

	end := time.Now()
	fmt.Println("Elamsde (ms):", end.UnixMilli()-start.UnixMilli())
}

func getInputCubes() []lava.Cube {
	cubes := []lava.Cube{}

	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	for line := range input {
		i := strings.Split(line, ",")
		cube := lava.Cube{
			X: asNum(i[0]),
			Y: asNum(i[1]),
			Z: asNum(i[2]),
		}
		cubes = append(cubes, cube)
	}

	return cubes
}

func asNum(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
