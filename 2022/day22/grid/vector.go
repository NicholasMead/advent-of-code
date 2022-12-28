package grid

type Vector [2]int

func (v Vector) Move(d Direction) Vector {
	moves := map[Direction]Vector{
		Up:    {+0, -1},
		Down:  {+0, +1},
		Left:  {-1, +0},
		Right: {+1, +0},
	}
	move := moves[d]

	return add(v, move)
}

func add(a, b Vector) Vector {
	return Vector{
		a[0] + b[0],
		a[1] + b[1],
	}
}
