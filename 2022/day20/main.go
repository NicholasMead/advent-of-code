package main

import (
	"aco/common"
	"fmt"
	"math"
	"strconv"

	"golang.org/x/exp/slices"
)

const key = 811_589_153

func main() {
	common.TimedMs(func() {
		frame := parseInput()
		keyFrame := make([]int, len(frame))

		for i := range keyFrame {
			keyFrame[i] = frame[i] * key
		}

		p1 := mixFrame(frame, 1)
		p2 := mixFrame(keyFrame, 10)

		fmt.Println("part1:", p1)
		fmt.Println("part2:", p2)
	})
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func mixFrame(frame []int, count int) int {
	var zero *int
	length := len(frame)
	current := make([]*int, length)
	for i := range frame {
		current[i] = &frame[i]
		if frame[i] == 0 {
			zero = &frame[i]
		}
	}

	for i := 0; i < length*count; i++ {
		target := &frame[i%length]
		current = mixInFrame(current, target)
	}

	zeroIndex := slices.Index(current, zero)
	value := 0
	for i := 1; i <= 3; i++ {
		value += *current[(zeroIndex+i*1000)%length]
	}
	return value
}

func mixInFrame(frame []*int, target *int) []*int {
	startPos := slices.Index(frame, target)

	len := len(frame) - 1 //Cant swap with self!
	moves := *target % len
	endPos := startPos + moves

	if endPos <= 0 {
		endPos += len
	} else {
		endPos = endPos % len
	}

	next := slices.Delete(frame, startPos, startPos+1)
	return slices.Insert(next, endPos, target)
}

func deRef(list []*int) []int {
	act := make([]int, len(list))
	for i := range act {
		act[i] = *list[i]
	}
	return act
}

func parseInput() []int {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)
	ints := []int{}

	for line := range input {
		v, _ := strconv.Atoi(line)
		ints = append(ints, v)
	}
	return ints
}
