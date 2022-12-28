package grid

type Grid interface {
	Move(from Vector, dir Direction) (Vector, Direction)
}
