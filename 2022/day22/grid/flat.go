package grid

import "fmt"

func Flat(surface Surface) Grid {
	return flat{surface}
}

type flat struct {
	Surface
}

// Move implements Grid
func (f flat) Move(from Vector, dir Direction) (Vector, Direction) {
	if !f.inBounds(from) {
		err := fmt.Errorf("Cannot start out of bounds: %v", from)
		panic(err)
	}

	target := from.Move(dir)

	if !f.inBounds(target) {
		switch dir {
		case Up:
			_, target[1] = f.ColumnBounds(target[0])
		case Down:
			target[1], _ = f.ColumnBounds(target[0])

		case Left:
			_, target[0] = f.RowBounds(target[1])
		case Right:
			target[0], _ = f.RowBounds(target[1])
		}
	}

	tile := f.Surface[target]
	if tile == Empty {
		return target, dir
	}

	return from, dir
}
