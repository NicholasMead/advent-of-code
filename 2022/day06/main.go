package main

import (
	"aco/common"
	"aoc/day6/signal"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawInput := <-common.ReadInput(os.Args[1])
	input := make(chan byte)
	go func(raw string) {
		for _, r := range raw {
			input <- byte(r)
		}
	}(rawInput)

	markerSize, e := strconv.Atoi(os.Args[2])
	reader := signal.CreateReader(markerSize)
	start, e := reader.FindStart(input)

	if e != nil {
		panic(e)
	} else {
		fmt.Println("Start:", start)
	}
}
