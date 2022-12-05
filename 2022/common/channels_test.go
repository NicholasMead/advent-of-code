package common_test

import (
	"aco/common"
	"testing"
)

func TestToSlice(t *testing.T) {
	//assign
	base := 10
	channel := make(chan int, base)
	for i := 0; i < base; i++ {
		channel <- i
	}
	close(channel)

	//act
	output := <-common.ToSlice(channel)

	//assert
	if len(output) != base {
		t.Errorf("output len of %d does not match input len of %d.", len(output), base)
	}
	for i, v := range output {
		if i != v {
			t.Errorf("value %v in position %d is incorrect and should be %d", v, i, i)
		}
	}
}
