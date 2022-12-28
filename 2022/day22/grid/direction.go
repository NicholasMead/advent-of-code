package grid

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

type Rotation string

const (
	clockwise     Rotation = "R"
	antiClockwise Rotation = "L"
)

func (d Direction) Rotate(r Rotation) Direction {
	transform := map[Rotation]int{
		clockwise:     1,
		antiClockwise: 3,
	}
	return Direction((int(d) + transform[r]) % 4)
}
