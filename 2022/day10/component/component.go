package component

import (
	"time"
)

type Component interface {
	Run(clk <-chan time.Time) (stop func(), err error)
}


