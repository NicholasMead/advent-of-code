package astar

import (
	"aoc/day12/queue"
	"errors"
)

type node[T comparable] interface {
	GetNeighbours() []T
	GetHuristic(T) int
	comparable
}

func FindRoute[Node node[Node]](start, end Node) ([]Node, error) {
	pathCosts := map[Node]int{start: 0}
	routes := map[Node]Node{}

	queue := queue.CreateCostOrdered[Node, int]()
	queue.Push(start, 0)

	for {
		current, _, next := queue.Pop()

		if !next {
			return nil, errors.New("No path found")
		} else if current == end {
			return reconstructPath(current, routes), nil
		}

		currentCost := pathCosts[current]
		neighbours := current.GetNeighbours()

		for _, neighbour := range neighbours {
			neighbourCost := currentCost + current.GetHuristic(neighbour)
			prevCost, prevVisited := pathCosts[neighbour]

			if !prevVisited || neighbourCost < prevCost {
				pathCosts[neighbour] = neighbourCost
				routes[neighbour] = current
				queue.Push(neighbour, neighbourCost+neighbour.GetHuristic(end))
			}
		}
	}
}

func reconstructPath[Node node[Node]](node Node, routes map[Node]Node) []Node {
	path := []Node{node}

	for {
		if prev, found := routes[path[0]]; found {
			path = append([]Node{prev}, path...)
		} else {
			return path
		}
	}
}
