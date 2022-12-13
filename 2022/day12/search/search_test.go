package search

import (
	"aco/common"
	"aoc/day12/grid"
	"testing"
)

func BenchmarkSearch(b *testing.B) {
	input := common.ReadInputPath("./benchmark.txt")
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

	const MaxInt = int(^uint(0) >> 1)

	b.Run("AStar", func(b *testing.B) {
		astar := CreateAStar[grid.Node]()
		best := MaxInt
		for n := 0; n < b.N; n++ {
			route, comp, err := astar.FindRoute(start, end)
			if len(route) != 413 || err != nil {
				b.Error(len(route), err)
			}
			if comp < best {
				best = comp
			}
		}
		b.Log(best)
	})

	b.Run("Dijkstra", func(b *testing.B) {
		dijkstra := CreateDijkstra[grid.Node]()
		best := MaxInt
		for n := 0; n < b.N; n++ {
			route, count, err := dijkstra.FindRoute(start, func(n grid.Node) bool { return n == end })
			if len(route) != 413 || err != nil {
				b.Error(len(route), err)
			}
			if count < best {
				best = count
			}
		}
		b.Log(best)
	})
}
