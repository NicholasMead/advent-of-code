package main

import (
	"aco/common"
	co "aoc/day15/coord"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"time"
)

type Target co.Coord
type Sensor co.Coord
type Beacon co.Coord
type ScanMap map[Sensor]Beacon

func main() {
	start := time.Now()
	flags := common.GetFlags()

	rowNum, _ := strconv.Atoi(flags.Args[0])
	max, _ := strconv.Atoi(flags.Args[1])
	input := common.ReadInputPath(*flags.File)

	scans := ScanMap{}

	for line := range input {
		exp := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
		match := exp.FindAllStringSubmatch(line, -1)

		if len(match) >= 1 {
			sx, _ := strconv.Atoi(match[0][1])
			sy, _ := strconv.Atoi(match[0][2])
			bx, _ := strconv.Atoi(match[0][3])
			by, _ := strconv.Atoi(match[0][4])
			sensor := Sensor{sx, sy}
			beacon := Beacon{bx, by}
			scans[sensor] = beacon
		}
	}

	fmt.Println("part1:", countEmptyInRow(scans, rowNum))

	min, max := 0, max
	targets := make(chan Target, 100000)
	answer := make(chan Target, 16)

	var workerWaitGroup sync.WaitGroup
	search := func(i int) {
		defer workerWaitGroup.Done()
		fmt.Println(i, "started")
		for target := range targets {
			if !inScanRange(target, scans) {
				fmt.Println(i, "found!")
				answer <- target
				return
			}
		}
		fmt.Println(i, "finished")
	}
	for x := 0; x < 16; x++ {
		workerWaitGroup.Add(1)
		go search(x)
	}

	queueTargers := func(sensor Sensor, beacon Beacon, wg *sync.WaitGroup) {
		dist := co.Dist(co.Coord(sensor), co.Coord(beacon))

		for _, t := range co.Coord(sensor).GetTrace(dist) {
			if t.X >= min && t.X <= max && t.Y >= min && t.Y <= max {
				targets <- Target(t)
			}
		}
		wg.Done()
	}

	var targetWaitGroup sync.WaitGroup
	for s, b := range scans {
		targetWaitGroup.Add(1)
		go queueTargers(s, b, &targetWaitGroup)
	}

	go func() {
		targetWaitGroup.Wait()
		close(targets)
	}()

	go func() {
		workerWaitGroup.Wait()
		close(answer)
	}()

	for found := range answer {
		freq := found.X*4000000 + found.Y
		fmt.Printf("part2: (%v,%v) => %v\n", found.X, found.Y, freq)
		break
	}
	end := time.Now()
	fmt.Printf("Execution time: %vms\n", end.UnixMilli()-start.UnixMilli())
}

func countEmptyInRow(scans ScanMap, rowNum int) int {
	count := 0
	left, right := getXBounds(scans)

	for x := left; x <= right; x++ {
		target := Target{x, rowNum}

		for sen, bec := range scans {
			if co.Coord(sen) == co.Coord(target) || co.Coord(bec) == co.Coord(target) {
				break
			}

			if inRange(target, sen, bec) {
				count++
				break
			}
		}
	}

	for x := left - 1; true; x-- {
		target := Target{x, rowNum}
		found := false

		for sen, bec := range scans {
			if inRange(target, sen, bec) {
				count++
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	for x := right + 1; true; x++ {
		target := Target{x, rowNum}
		found := false

		for sen, bec := range scans {
			if inRange(target, sen, bec) {
				count++
				found = true
				break
			}
		}

		if !found {
			break
		}
	}

	return count
}

func getXBounds(scans ScanMap) (left, right int) {
	var lSensor, rSensor Sensor
	var first bool = true

	for sensor := range scans {
		if first {
			lSensor, rSensor = sensor, sensor
			first = false
			continue
		}

		if sensor.X < lSensor.X {
			lSensor = sensor
		}
		if sensor.X > rSensor.X {
			rSensor = sensor
		}
	}

	return lSensor.X, rSensor.X
}

func inScanRange(target Target, scans ScanMap) bool {
	for sen, bec := range scans {
		if inRange(target, sen, bec) {
			return true
		}
	}

	return false
}

func inRange(target Target, sensor Sensor, beacon Beacon) bool {
	sensorRange := co.Dist(co.Coord(sensor), co.Coord(beacon))
	targetRange := co.Dist(co.Coord(sensor), co.Coord(target))

	return targetRange <= sensorRange
}
