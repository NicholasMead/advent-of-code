package grid

func Cube(s Surface, dim int) Grid {
	return cube{s, dim}
}

type cube struct {
	Surface
	dim int
}

// Move implements Grid
func (c cube) Move(from Vector, dir Direction) (Vector, Direction) {
	target := from.Move(dir)
	targetDir := dir

	if !c.inBounds(target) {
		/*
			layout:
			_ A B
			_ C _
			D E _
			F _ _
		*/
		face := c.faceOfPosition(from)
		switch dir {
		case Right:
			switch face[1] {
			case 0: //B->E
				target[0] = c.dim*2 - 1
				target[1] = c.dim*2 + c.relRev(from[1])
				targetDir = Left
			case 1: //C->B
				target[0] = c.dim*2 + c.rel(from[1])
				target[1] = c.dim*1 - 1
				targetDir = Up
			case 2: //E->B
				target[0] = c.dim*3 - 1
				target[1] = c.relRev(from[1])
				targetDir = Left
			case 3: //F-> E
				target[0] = c.dim*1 + c.rel(from[1])
				target[1] = c.dim*3 - 1
				targetDir = Up
			}
		case Down:
			switch face[0] {
			case 0: //F-B
				target[0] = c.dim*2 + c.rel(from[0])
				target[1] = c.dim * 0
				targetDir = Down
			case 1: //E->F
				target[0] = c.dim*1 - 1
				target[1] = c.dim*3 + c.rel(from[0])
				targetDir = Left
			case 2: //B->C
				target[0] = c.dim*2 - 1
				target[1] = c.dim*1 + c.rel(from[0])
				targetDir = Left
			}
		case Left:
			switch face[1] {
			case 0: //A->D
				target[0] = c.dim * 0
				target[1] = c.dim*2 + c.relRev(from[1])
				targetDir = Right
			case 1: //C->D
				target[0] = c.dim*0 + c.rel(from[1])
				target[1] = c.dim * 2
				targetDir = Down
			case 2: //D->A
				target[0] = c.dim * 1
				target[1] = c.dim*0 + c.relRev(from[1])
				targetDir = Right
			case 3: //F->A
				target[0] = c.dim*1 + c.rel(from[1])
				target[1] = c.dim * 0
				targetDir = Down
			}
		case Up:
			switch face[0] {
			case 0: // D->C
				target[0] = c.dim * 1
				target[1] = c.dim + c.rel(from[0])
				targetDir = Right
			case 1: // A->F
				target[0] = c.dim * 0
				target[1] = c.dim*3 + c.rel(from[0])
				targetDir = Right
			case 2: // B->F
				target[0] = c.dim*0 + c.rel(from[0])
				target[1] = c.dim*4 - 1
				targetDir = Up
			}
		}
	}

	tile := c.Surface[target]
	if tile == Empty {
		return target, targetDir
	}

	return from, dir
}

func (c cube) faceOfPosition(v Vector) Vector {
	return Vector{v[0] / c.dim, v[1] / c.dim}
}

func (c cube) rel(i int) int {
	return i % c.dim
}

func (c cube) relRev(i int) int {
	return c.dim - 1 - c.rel(i)
}
