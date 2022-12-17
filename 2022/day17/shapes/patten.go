package shapes

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
