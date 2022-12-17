package main

import (
	"aco/common"
	cyc "aoc/day17/cycle"
	shp "aoc/day17/shapes"
	sim "aoc/day17/simulation"
	tow "aoc/day17/towers"
	"fmt"
)

func main() {
	parts := map[string]uint64{
		"part1": 2022,
		"part2": 1_000_000_000_000,
	}

	for part, size := range parts {
		sim := sim.Simulation{
			Shapes: cyc.CreateCycle(shp.HLine, shp.Cross, shp.Corner, shp.VLine, shp.Box),
			Jets:   cyc.CreateCycle(parseJets()...),
			Tower:  tow.CreateTower(7),
		}

		answer := sim.Run(size)
		fmt.Printf("%v: %v\n", part, answer)
	}
}

func parseJets() []int {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)
	jets := []int{}

	for line := range input {
		for _, j := range line {
			switch j {
			case '<':
				jets = append(jets, -1)
			case '>':
				jets = append(jets, 1)
			default:
				panic(fmt.Sprintf("What is: %c?", j))
			}
		}
	}
	return jets
}
