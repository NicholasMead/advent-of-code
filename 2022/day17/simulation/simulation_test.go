package simulation

import (
	"aoc/day17/cycle"
	"aoc/day17/shapes"
	"aoc/day17/towers"
	"testing"
)

func TestSimulationCycle(t *testing.T) {
	sim := Simulation{
		Shapes: cycle.CreateCycle(
			shapes.HLine,
			shapes.Cross,
			shapes.Corner,
			shapes.VLine,
			shapes.Box,
		),
		Jets:  cycle.CreateCycle(1, 1, 2, 2),
		Tower: towers.CreateTower(7),
	}

	sim.Run(uint64(sim.CycleLength()))

	if i := sim.Shapes.Pos(); i != 0 {
		t.Error("Expected shapes to be at start:", i)
	}
	if i := sim.Jets.Pos(); i != 0 {
		t.Error("Expected jets to be at start:", i)
	}
}

// func TestSimulationPatten(t *testing.T) {
// 	sim := Simulation{
// 		Shapes: cycle.CreateCycle(
// 			shapes.Box,
// 		),
// 		Jets:  cycle.CreateCycle(0),
// 		Tower: towers.CreateTower(7),
// 	}

// 	sim.Run(10)

// 	if i, _ := sim.Shapes.Pos(); i != 0 {
// 		t.Error("Expected shapes to be at start:", i)
// 	}
// 	if i, _ := sim.Jets.Pos(); i != 0 {
// 		t.Error("Expected jets to be at start:", i)
// 	}
// }
