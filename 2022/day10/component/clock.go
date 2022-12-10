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

func (c *internalClock) onTick(clk <-chan time.Time, action func(time.Time)) (stop func(), err error) {
	if err := c.startClock(clk); err != nil {
		return nil, err
	}

	go func() {
		for t := range c.clk {
			action(t)
		}
	}()

	return c.stopClock, nil
}

func Ticker() (clk <-chan time.Time, tick func()) {
	clock := make(chan time.Time)
	tick = func() {
		clock <- time.Now()
	}
	return clock, tick
}
