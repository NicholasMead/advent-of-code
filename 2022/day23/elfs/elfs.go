package elfs

import (
	"aoc/day23/collections"
	"fmt"
	"math"
)

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

type Coord struct {
	X, Y int
}

func Bounds(elfs map[Coord]interface{}) (minX, maxX, minY, maxY int) {
	minX, minY = math.MaxInt, math.MaxInt
	maxX, maxY = math.MinInt, math.MinInt

	for elf := range elfs {
		minX = min(minX, elf.X)
		maxX = max(maxX, elf.X)
		minY = min(minY, elf.Y)
		maxY = max(maxY, elf.Y)
	}
	return minX, maxX, minY, maxY
}

func PrintElfs(elfs map[Coord]interface{}) string {
	minX, maxX, minY, maxY := Bounds(elfs)

	output := ""
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if _, elf := elfs[Coord{x, y}]; elf {
				output += "#"
			} else {
				output += "."
			}
		}
		output += "\n"
	}
	return output
}

func Run(start []Coord, maxRounds int) (map[Coord]interface{}, int) {
	var (
		elfs      = map[Coord]interface{}{}
		moveCycle = collections.Cycle(North, South, West, East)
	)

	for _, elf := range start {
		elfs[elf] = struct{}{}
	}

	for round := 0; round < maxRounds; round++ {
		if round > 0 && round%10_000 == 0 {
			fmt.Println(round)
		}
		var (
			moves   = map[Coord]Coord{} //The moves we should make it part 2
			targets = map[Coord]int{}   //The number of elfs tring to move to a square in part 2
		)

		//part 1 - find where the elfs want to go
		for elf := range elfs {
			moveCycle.Reset(round)

			neibours := map[Coord]interface{}{}
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					pos := Coord{elf.X + dx, elf.Y + dy}
					if _, found := elfs[pos]; found {
						neibours[pos] = struct{}{}
					}
				}
			}

			if len(neibours) < 2 { // We should atleast find the target elf!
				continue
			}

			var (
				checks     [3]Coord //Coords we should check
				obstructed = false  //Did we find someone
				move       Coord    //Where we should go if we can
			)
			for i := 0; i < 4; i++ {
				dir := moveCycle.Next() //Which direct to look next
				switch dir {
				case North:
					checks = [3]Coord{{-1, +1}, {+0, +1}, {+1, +1}}
					move = Coord{0, +1}
				case East:
					checks = [3]Coord{{+1, -1}, {+1, +0}, {+1, +1}}
					move = Coord{+1, 0}
				case South:
					checks = [3]Coord{{-1, -1}, {+0, -1}, {+1, -1}}
					move = Coord{0, -1}
				case West:
					checks = [3]Coord{{-1, -1}, {-1, +0}, {-1, +1}}
					move = Coord{-1, 0}
				}
				obstructed = false
				for _, c := range checks {
					check := Coord{elf.X + c.X, elf.Y + c.Y}
					_, obstructed = elfs[check]
					if obstructed {
						break
					}
				}
				if !obstructed {
					move.X += elf.X
					move.Y += elf.Y
					moves[elf] = move
					targets[move]++
					break
				}
			}
		}

		//part 2 - try and move
		didMove := false
		for from, to := range moves {
			if targets[to] == 1 {
				elfs[to] = struct{}{}
				delete(elfs, from)
				didMove = true
			}
		}
		if !didMove {
			return elfs, round + 1
		}
	}
	return elfs, maxRounds + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
