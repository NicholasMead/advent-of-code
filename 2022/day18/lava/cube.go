package lava

type Cube struct {
	X, Y, Z int
}

func (cube Cube) getAdjacent() []Cube {
	return []Cube{
		{cube.X + 1, cube.Y, cube.Z},
		{cube.X - 1, cube.Y, cube.Z},
		{cube.X, cube.Y + 1, cube.Z},
		{cube.X, cube.Y - 1, cube.Z},
		{cube.X, cube.Y, cube.Z + 1},
		{cube.X, cube.Y, cube.Z - 1},
	}
}
