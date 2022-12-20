package common

import (
	"fmt"
	"time"
)

func TimedMs(fn func()) {
	start := time.Now()
	fn()
	end := time.Now()
	fmt.Println("elapsed (ms):", end.UnixMilli()-start.UnixMilli())
}
