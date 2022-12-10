package grid

type Grid interface {
	addTree(x, y, h int)

	IsVisible(x, y int) bool
	SceanicScore(x, y int) int
	GetSize(x, y int) int
	Height() int
	Width() int
}

type grid struct {
	width, height int
	trees         [][]int
}

// SceanicScore implements Grid
func (g *grid) SceanicScore(x int, y int) int {
	h := g.trees[x][y]
	scores := [4]int{}
	views := [4][]int{
		g.lookLeftFrom(x, y),
		g.lookUpFrom(x, y),
		g.lookRightFrom(x, y),
		g.lookDownFrom(x, y),
	}

	for i, view := range views {
		for _, tree := range view {
			scores[i]++
			if tree >= h {
				break
			}
		}
	}

	score := 1
	for _, s := range scores {
		score *= s
	}

	return score
}

// Height implements Grid
func (g *grid) Height() int {
	return g.height
}

// Width implements Grid
func (g *grid) Width() int {
	return g.width
}

// GetHeight implements Grid
func (g *grid) GetSize(x int, y int) int {
	return g.trees[x][y]
}

// addTree implements Grid
func (g *grid) addTree(x int, y int, h int) {
	g.trees[x][y] = h
}

// isVisible implements Grid
func (g *grid) IsVisible(x int, y int) bool {
	h := g.trees[x][y]
	hidden := [4]bool{}
	views := [4][]int{
		g.lookLeftFrom(x, y),
		g.lookUpFrom(x, y),
		g.lookRightFrom(x, y),
		g.lookDownFrom(x, y),
	}

	for i, view := range views {
		for _, tree := range view {
			if tree >= h {
				hidden[i] = true
				break
			}
		}
	}

	for _, hidden := range hidden {
		if !hidden {
			return true
		}
	}
	return false
}

func (g *grid) lookUpFrom(x int, y int) []int {
	trees := []int{}
	for i := y - 1; i >= 0; i-- {
		trees = append(trees, g.trees[x][i])
	}
	return trees
}

func (g *grid) lookDownFrom(x int, y int) []int {
	trees := []int{}
	for i := y + 1; i < g.height; i++ {
		trees = append(trees, g.trees[x][i])
	}
	return trees
}

func (g *grid) lookRightFrom(x int, y int) []int {
	trees := []int{}
	for i := x + 1; i < g.width; i++ {
		trees = append(trees, g.trees[i][y])
	}
	return trees
}

func (g *grid) lookLeftFrom(x int, y int) []int {
	trees := []int{}
	for i := x - 1; i >= 0; i-- {
		trees = append(trees, g.trees[i][y])
	}
	return trees
}

func NewGrid(trees [][]int) Grid {
	height, width := len(trees), len(trees[0])

	t := make([][]int, width)
	for w := range t {
		t[w] = make([]int, height)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t[x][y] = trees[x][y]
		}
	}

	return &grid{width, height, trees}
}
