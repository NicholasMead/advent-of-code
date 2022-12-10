package component

import (
	"time"
)

type Buffer[T any] interface {
	Component

	Peak() []T
}

type memBuffer[T any] struct {
	internalClock
	log  []T
	in   <-chan T
	read chan readAction[T]
}

type readAction[T any] struct {
	done chan<- []T
}

func (l *memBuffer[T]) Run(clk <-chan time.Time) (stop func(), err error) {
	if err := l.startClock(clk); err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case _, open := <-l.clk:
				if open {
					l.tick()
				} else {
					l.clk = nil
				}
			case read := <-l.read:
				read.done <- append([]T{}, l.log...)
				close(read.done)
			}
		}
	}()

	return l.stopClock, nil
}

func (l *memBuffer[T]) tick() {
	i := <-l.in
	l.log = append(l.log, i)
}

func (l *memBuffer[T]) Peak() []T {
	done := make(chan []T)

	l.read <- readAction[T]{done}
	return <-done
}

func MemBuffer[T any](in <-chan T) Buffer[T] {
	return &memBuffer[T]{
		log:  []T{},
		in:   in,
		read: make(chan readAction[T]),
	}
}
