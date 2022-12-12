package main

import (
	"aco/common"
	"aoc/day12/astar"
	"aoc/day12/grid"
	"flag"
	"fmt"
)

func main() {
	file := flag.String("f", "-", "Input file")
	flag.Parse()

	input := common.ReadInput(*file)
	lines := []string{}
	for i := range input {
		lines = append(lines, i)
	}

	g := grid.Create(len(lines), len(lines[0]))
	var start, end grid.Node
	for y, row := range lines {
		for x, h := range row {
			switch h {
			case 'S':
				g.AddNode(x, y, 0)
				start, _ = g.GetNode(x, y)
			case 'E':
				g.AddNode(x, y, 25)
				end, _ = g.GetNode(x, y)
			default:
				g.AddNode(x, y, int(h-'a'))
			}
		}
	}

	fmt.Println(g)
	fmt.Println("***Starting part 1***")
	fmt.Println("Start:", start)
	fmt.Println("End:", end)

	route, err := astar.FindRoute(start, end)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("***Found***")
		fmt.Printf("Length:%d\nSteps:%d\n", len(route), len(route)-1)
	}

	fmt.Println("***Starting part 2***")
	starts := g.GetNodesAtElevation(0)
	fmt.Println(len(starts)+1, "starting points")
	best := len(route) - 1
	for _, start = range starts {
		route, err = astar.FindRoute(start, end)
		if err == nil && len(route)-1 < best {
			best = len(route) - 1
		}
	}
	fmt.Println("Part 2:", best)
}
