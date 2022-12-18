package lava

import (
	"math"
)

type Droplet interface {
	Add(cube Cube)
	// TotalSurfaceArea() int
	SurfaceArea() (total int, external int)
}

func FormDroplet(cubes []Cube) Droplet {
	droplet := droplet{
		cubes: map[Cube]int{},
		minX:  math.MaxInt,
		maxX:  -math.MaxInt,
		minY:  math.MaxInt,
		maxY:  -math.MaxInt,
		minZ:  math.MaxInt,
		maxZ:  -math.MaxInt,
	}

	for _, cube := range cubes {
		droplet.Add(cube)
	}

	return &droplet
}

type droplet struct {
	cubes      map[Cube]int
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

func (droplet *droplet) Add(cube Cube) {
	droplet.cubes[cube] = 6
	droplet.minX = min(droplet.minX, cube.X)
	droplet.minY = min(droplet.minY, cube.Y)
	droplet.minZ = min(droplet.minZ, cube.Z)
	droplet.maxX = max(droplet.maxX, cube.X)
	droplet.maxY = max(droplet.maxY, cube.Y)
	droplet.maxZ = max(droplet.maxZ, cube.Z)

	// from part 1
	// for _, adj := range cube.getAdjacent() {
	// 	_, found := droplet.cubes[adj]
	// 	if found {
	// 		droplet.cubes[adj]--
	// 		droplet.cubes[cube]--
	// 	}
	// }
}

// From Part 1
// func (d droplet) TotalSurfaceArea() int {
// 	surfaceArea := 0
// 	for _, area := range d.cubes {
// 		surfaceArea += area
// 	}
// 	return surfaceArea
// }

func (d droplet) SurfaceArea() (total int, external int) {
	ext := map[Cube]bool{} //cache calculations

	for cube := range d.cubes {
		for _, adj := range cube.getAdjacent() {
			if _, found := d.cubes[adj]; !found {
				total++
				if ext[adj] {
					external++
				} else if d.isExtenal(adj) {
					ext[adj] = true // no need to check here again!
					external++
				}
			}
		}
	}
	return total, external
}

// Checks if a cube space is outside the droplet
// Start outside the droplet, BFS though adjecent (unvisited cubes) until the target is found
// Only accounds for cubes at most 1 away from the droplet, but its also an internal function so bite me (or write the extra boundry condition yourself, I dont mind, I was too busy writing this comment)
func (d droplet) isExtenal(cube Cube) bool {
	start := Cube{
		X: d.minX - 1,
		Y: d.minY - 1,
		Z: d.minZ - 1,
	} // just outside the droplet
	queue := []Cube{start}
	visited := map[Cube]bool{start: true}
	var current Cube

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:] //pop!
		for _, next := range current.getAdjacent() {
			if next == cube {
				return true
			}

			tooFar := next.X > d.maxX+2 ||
				next.Y > d.maxY+2 ||
				next.Z > d.maxZ+2 ||
				next.X < d.minX-2 ||
				next.Y < d.minY-2 ||
				next.Z < d.minZ-2

			if tooFar {
				continue //lets not check to inf!
			}
			if visited[next] {
				continue //we have been here
			}
			if _, found := d.cubes[next]; found {
				continue //not free space
			}

			visited[next] = true
			queue = append(queue, next)
		}
	}
	return false
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
