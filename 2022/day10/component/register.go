package component

import (
	"fmt"
	"time"
)

type Command struct {
	Program Program
	Arg     int
}

type Program string

const (
	addx = "addx"
	noop = "noop"
)

type linearRegister struct {
	value  int
	input  <-chan Command
	output chan<- int

	internalClock
}

func (r *linearRegister) String() string {
	return fmt.Sprintf("{v: %d}", r.value)
}

func (r *linearRegister) Run(clk <-chan time.Time) (stop func(), err error) {
	if err := r.startClock(clk); err != nil {
		return nil, err
	}

	go func() {
		busy := 0
		cmd := Command{noop, 0}

		for range r.clk {
			r.output <- r.value

			if busy == 0 {
				cmd, busy = r.readCmd()
			} else {
				busy--
			}

			if busy == 0 {
				r.execCmd(cmd)
			}
		}
	}()

	return r.stopClock, nil
}

func (r *linearRegister) readCmd() (Command, int) {
	cmd, open := <-r.input

	switch {
	case !open:
		cmd.Program = noop
		return cmd, 0
	case cmd.Program == addx:
		return cmd, 1
	default:
		return cmd, 0
	}
}

func (r *linearRegister) execCmd(cmd Command) {
	switch cmd.Program {
	case addx:
		(*r).value += cmd.Arg
	}
}

func Register(input <-chan Command, output chan<- int) Component {
	return &linearRegister{
		input:  input,
		output: output,

		value: 1,
	}
}
