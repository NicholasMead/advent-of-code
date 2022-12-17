package shapes

import "math"

type Rock interface {
	GetStones() []Coord
	Move(x, y int) Rock
	GetBounds() (left, right, top, bottom int)
}

func CreateRock(x, y int, patten Patten) Rock {
	return rock{
		Coord{x, y},
		patten,
	}
}

type rock struct {
	anchor Coord
	patten Patten
}

// GetBounds implements Shape
func (r rock) GetBounds() (left int, right int, top int, bottom int) {
	left = math.MaxInt
	right = -math.MaxInt
	bottom = math.MaxInt
	top = -math.MaxInt

	for _, p := range r.GetStones() {
		if p.X > right {
			right = p.X
		}
		if p.X < left {
			left = p.X
		}
		if p.Y > top {
			top = p.Y
		}
		if p.Y < bottom {
			bottom = p.Y
		}
	}
	return left, right, top, bottom
}

// Move implements Shape
func (r rock) Move(x, y int) Rock {
	return rock{
		r.anchor.Move(x, y),
		r.patten,
	}
}

// GetPositions implements Shape
func (r rock) GetStones() []Coord {
	stones := make([]Coord, 0, len(r.patten))
	for _, p := range r.patten {
		stones = append(stones, r.anchor.Add(p))
	}
	return stones
}
