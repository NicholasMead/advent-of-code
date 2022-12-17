package shapes

import "math"

type Patten []Coord

var (
	HLine Patten = Patten{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}
	Cross = Patten{
		{1, 0},
		{1, 1},
		{1, 2},
		{0, 1},
		{2, 1},
	}
	Corner = Patten{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{2, 2},
	}
	VLine Patten = Patten{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}
	Box Patten = Patten{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
)

func (p Patten) AboveY(y int) Patten {
	above := Patten{}
	for _, s := range p {
		if s.Y >= y {
			s.Y -= y
			above = append(above, s)
		}
	}
	return above
}

func (p Patten) Match(m Patten) bool {
	if len(p) != len(m) {
		return false
	}

	for _, ps := range p {
		match := false
		for _, ms := range m {
			if ps == ms {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	return true
}

// GetBounds implements Shape
func (p Patten) GetBounds() (left int, right int, top int, bottom int) {
	left = math.MaxInt
	right = -math.MaxInt
	bottom = math.MaxInt
	top = -math.MaxInt

	for _, s := range p {
		if s.X > right {
			right = s.X
		}
		if s.X < left {
			left = s.X
		}
		if s.Y > top {
			top = s.Y
		}
		if s.Y < bottom {
			bottom = s.Y
		}
	}
	return left, right, top, bottom
}
