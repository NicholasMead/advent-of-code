package component

import (
	"errors"
	"time"
)

type internalClock struct {
	clk <-chan time.Time
}

func (c *internalClock) startClock(clk <-chan time.Time) error {
	if c.clk != nil {
		return errors.New("Already Running")
	} else {
		c.clk = clk
		return nil
	}
}

func (c *internalClock) stopClock() {
	c.clk = nil
}

func Ticker() (clk <-chan time.Time, tick func()) {
	clock := make(chan time.Time)
	tick = func() {
		clock <- time.Now()
	}
	return clock, tick
}
