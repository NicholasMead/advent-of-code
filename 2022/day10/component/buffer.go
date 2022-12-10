package component

import "time"

type Buffer[T any] interface {
	Component

	Peak() []T
}

type memBuffer[T any] struct {
	internalClock
	log []T
	in  <-chan T
}

func (l *memBuffer[T]) Run(clk <-chan time.Time) (stop func(), err error) {
	action := func(_ time.Time) {
		i := <-l.in
		l.log = append(l.log, i)
	}

	return l.onTick(clk, action)
}

func (l *memBuffer[T]) Peak() []T {
	output := make([]T, len(l.log))
	copy(output, l.log)
	return output
}

func MemBuffer[T any](in <-chan T) Buffer[T] {
	return &memBuffer[T]{
		log: []T{},
		in:  in,
	}
}
