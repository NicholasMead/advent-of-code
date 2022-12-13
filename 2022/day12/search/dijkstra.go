package search

import (
	"aoc/day12/collections"
	"aoc/day12/queue"
	"errors"
)

type Dijkstra[Node node[Node]] interface {
	FindRoute(start Node, condition func(Node) bool) ([]Node, int, error)
}

func CreateDijkstra[Node node[Node]]() Dijkstra[Node] {
	return &dijkstra[Node]{}
}

type dijkstra[Node node[Node]] struct{}

// FindRoute implements Dijkstra
func (*dijkstra[Node]) FindRoute(start Node, condition func(Node) bool) ([]Node, int, error) {
	count := 0
	nodeCost := collections.CreateHashmap[Node, int]()
	routes := collections.CreateHashmap[Node, Node]()
	queue := queue.CreateCostOrdered[Node, int]()
	nodeCost.Store(start, 0)
	queue.Push(start, 0)

	for {
		current, cost, next := queue.Pop()

		if !next {
			return nil, count, errors.New("Route not found")
		} else if condition(current) {
			return reconstructPath(current, routes), count, nil
		} else {
			count++
		}

		for _, neighbor := range current.GetNeighbours() {
			nCost := cost + current.GetNeighbourDist(neighbor)
			prevCost, visited := nodeCost.Get(neighbor)

			if !visited || nCost < prevCost {
				nodeCost.Store(neighbor, nCost)
				routes.Store(neighbor, current)
				queue.Push(neighbor, nCost)
			}
		}
	}
}
