package grid

import "errors"

type Grid interface {
	GetNode(x, y int) (Node, error)
	AddNode(x, y, h int)
	GetNodesAtElevation(h int) []Node
}

func Create(height, width int) Grid {
	nodes := make([][]Node, width)
	for i := range nodes {
		nodes[i] = make([]Node, height)
	}
	return &grid{height, width, nodes}
}

type grid struct {
	height, width int
	nodes         [][]Node
}

// GetNodesAtElevation implements Grid
func (g *grid) GetNodesAtElevation(h int) []Node {
	nodes := []Node{}
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			node := g.nodes[x][y]
			if node.h == h {
				nodes = append(nodes, node)
			}
		}
	}
	return nodes
}

// addNode implements Grid
func (g *grid) AddNode(x, y int, h int) {
	g.nodes[x][y] = Node{
		x:    x,
		y:    y,
		h:    h,
		grid: g,
	}
}

// getNode implements Grid
func (g *grid) GetNode(x int, y int) (Node, error) {
	if x < 0 || y < 0 || x >= g.width || y >= g.height {
		return *new(Node), errors.New("Out of Bounds")
	} else {
		return g.nodes[x][y], nil
	}
}

func (g grid) String() string {
	s := ""
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			s += string(rune('a' + g.nodes[x][y].h))

		}
		s += "\n"
	}
	return s
}
