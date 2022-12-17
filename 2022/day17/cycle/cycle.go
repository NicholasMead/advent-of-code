package cycle

type Cycle[T any] interface {
	Next() (next T, index int)
	Pos() (index int)
	Len() int
	Reset()
}

func CreateCycle[T any](items ...T) Cycle[T] {
	return &cycle[T]{
		items,
		0,
	}
}

type cycle[T any] struct {
	patten []T
	index  int
}

// Next implements Cycle
func (c *cycle[T]) Next() (next T, index int) {
	defer c.inc()

	index = c.index
	next = c.patten[index]
	return next, index
}

func (c *cycle[T]) Len() int {
	return len(c.patten)
}

// Next implements Cycle
func (c *cycle[T]) Pos() (index int) {
	index = c.index
	return index
}

// Next implements Cycle
func (c *cycle[T]) Reset() {
	c.index = 0
}

func (c *cycle[T]) inc() {
	c.index = (c.index + 1) % len(c.patten)
}
