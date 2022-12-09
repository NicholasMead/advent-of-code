package rope

import "oac/day9/coord"

type Rope interface {
	MoveDirection(coord.Direction)
	Head() coord.Coord
	Tail() coord.Coord
}

type rope struct {
	segments []coord.Coord
}

// Head implements Rope
func (r *rope) Head() coord.Coord {
	return r.segments[0]
}

// MoveDirection implements Rope
func (r *rope) MoveDirection(dir coord.Direction) {
	r.segments[0] = r.segments[0].StepDirection(dir)

	for i := 1; i < len(r.segments); i++ {
		r.segments[i] = r.segments[i].StepTowards(r.segments[i-1])
	}
}

// Tail implements Rope
func (r *rope) Tail() coord.Coord {
	return r.segments[len(r.segments)-1]
}

func CreateRope(length int) Rope {
	return &rope{make([]coord.Coord, length)}
}
