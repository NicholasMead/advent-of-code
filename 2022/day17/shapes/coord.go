package shapes

type Coord struct{ X, Y int }

func (c Coord) Add(a Coord) Coord {
	c.X += a.X
	c.Y += a.Y
	return c
}

func (c Coord) Move(x, y int) Coord {
	c.X += x
	c.Y += y
	return c
}
