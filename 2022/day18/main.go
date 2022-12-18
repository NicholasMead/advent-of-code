package main

import (
	"aco/common"
	"aoc/day18/lava"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cubes := getInputCubes()
	droplet := lava.FormDroplet(cubes)
	fmt.Println("Part1:", droplet.TotalSurfaceArea())
	fmt.Println("Part2:", droplet.ExternalSurfaceArea())
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
