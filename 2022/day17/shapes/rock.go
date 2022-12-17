package shapes

type Rock interface {
	GetStones() []Coord
	Move(x, y int) Rock
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
