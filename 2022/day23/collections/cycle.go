package collections

type Cycles[T any] interface {
	Next() T
	Reset(int)
}

func Cycle[T any](items ...T) Cycles[T] {
	return &cycle[T]{items, 0}
}

type cycle[T any] struct {
	items []T
	pos   int
}

func (c *cycle[T]) Next() T {
	next := c.items[c.pos]
	c.pos = (c.pos + 1) % len(c.items)
	return next
}

func (c *cycle[T]) Reset(i int) {
	c.pos = i % len(c.items)
}
