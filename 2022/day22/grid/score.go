package grid

func Score(pos Vector, dir Direction) int {
	x, y := pos[0]+1, pos[1]+1
	return y*1000 + x*4 + int(dir)
}
