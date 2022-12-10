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
	log []T
	in  <-chan T
	req chan action
}

type action struct {
	enact func()
	done  chan<- interface{}
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
					i := <-l.in
					l.log = append(l.log, i)
				} else {
					l.clk = nil
				}
			case req := <-l.req:
				req.enact()
				req.done <- new(interface{})
				close(req.done)
			}
		}
	}()

	return l.stopClock, nil
}

func (l *memBuffer[T]) Peak() []T {
	output := make([]T, len(l.log))
	done := make(chan interface{})
	act := func() {
		copy(output, l.log)
	}

	l.req <- action{act, done}
	<-done

	return output
}

func MemBuffer[T any](in <-chan T) Buffer[T] {
	return &memBuffer[T]{
		log: []T{},
		in:  in,
		req: make(chan action),
	}
}
