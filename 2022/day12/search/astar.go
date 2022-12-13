package search

import (
	"aoc/day12/collections"
	"aoc/day12/queue"
	"errors"
)

type AStar[Node node[Node]] interface {
	FindRoute(start, end Node) ([]Node, int, error)
}

func CreateAStar[Node node[Node]]() AStar[Node] {
	return &aStar[Node]{}
}

type node[T comparable] interface {
	GetNeighbours() []T
	GetNeighbourDist(T) int
	GetHuristic(T) int
	comparable
	collections.Hashable
}

type aStar[Node node[Node]] struct {
}

func (*aStar[Node]) FindRoute(start, end Node) ([]Node, int, error) {
	comparisons := 0
	pathCosts := collections.CreateHashmap[Node, int]()
	routes := collections.CreateHashmap[Node, Node]()
	queue := queue.CreateCostOrdered[Node, int]()
	pathCosts.Store(start, 0)
	queue.Push(start, 0)

	for {
		current, _, next := queue.Pop()

		if !next {
			return nil, comparisons, errors.New("No path found")
		} else if current == end {
			return reconstructPath(current, routes), comparisons, nil
		} else {
			comparisons++
		}

		currentCost, _ := pathCosts.Get(current)
		neighbours := current.GetNeighbours()

		for _, neighbour := range neighbours {
			neighbourCost := currentCost + current.GetNeighbourDist(neighbour)
			prevCost, visited := pathCosts.Get(neighbour)

			if !visited || neighbourCost < prevCost {
				pathCosts.Store(neighbour, neighbourCost)
				routes.Store(neighbour, current)
				queue.Push(neighbour, neighbourCost+neighbour.GetHuristic(end))
			}
		}
	}
}

func reconstructPath[Node node[Node]](node Node, routes collections.Map[Node, Node]) []Node {
	path := []Node{node}

	for {
		if prev, found := routes.Get(path[0]); found {
			path = append([]Node{prev}, path...)
		} else {
			return path
		}
	}
}
