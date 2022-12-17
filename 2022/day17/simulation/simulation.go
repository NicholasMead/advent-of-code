package simulation

import (
	"aoc/day17/cycle"
	"aoc/day17/shapes"
	"aoc/day17/towers"
)

type Jets cycle.Cycle[int]
type Shapes cycle.Cycle[shapes.Patten]

type Simulation struct {
	Tower  towers.Tower
	Jets   Jets
	Shapes Shapes
}

func (s *Simulation) Run(length uint64) uint64 {

	var skippedSize uint64 = 0
	var count uint64 = 0
	// cycle := uint64(s.CycleLength()) <- Why did I think this was a good idea... check every 1 full itteration cycle?
	cycle := uint64(1)
	store := []struct {
		patten           shapes.Patten
		size             int
		shapePos, jetPos int
	}{
		{
			s.Tower.AsPatten(),
			s.Tower.Size(),
			s.Shapes.Pos(),
			s.Jets.Pos(),
		},
	}

	for count+cycle < length {
		s.run(cycle)
		count += cycle
		cNum := count / cycle

		patten := s.Tower.AsPatten()
		match := false
		for prevNum, prev := range store {
			match = prev.shapePos == s.Shapes.Pos() &&
				prev.jetPos == s.Jets.Pos() &&
				patten.Match(prev.patten)

			if match {
				// Super maths time!
				pattenLength := (cNum - uint64(prevNum)) * cycle
				pattenSize := s.Tower.Size() - prev.size

				remainder := length - count
				skippedCycles := remainder / pattenLength
				skippedRemainder := remainder % pattenLength
				skippedSize = uint64(pattenSize) * skippedCycles

				count = length - skippedRemainder
				break
			}
		}

		if match {
			break
		} else {
			store = append(store, struct {
				patten           shapes.Patten
				size             int
				shapePos, jetPos int
			}{
				s.Tower.AsPatten(),
				s.Tower.Size(),
				s.Shapes.Pos(),
				s.Jets.Pos(),
			})
		}
	}

	s.run(length - count)
	return uint64(s.Tower.Size()) + skippedSize
}

func (s Simulation) run(length uint64) {
	var i uint64 = 0
	for ; i < length; i++ {
		shape, _ := s.Shapes.Next()
		s.Tower.Drop(shape, s.Jets)
	}
}

func (s Simulation) CycleLength() int {
	return s.Shapes.Len() * s.Jets.Len()
}
