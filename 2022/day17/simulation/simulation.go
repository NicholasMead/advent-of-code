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
	store := []struct {
		signature        shapes.Patten
		size             int
		shapePos, jetPos int
	}{
		{
			s.Tower.Signature(),
			s.Tower.Size(),
			s.Shapes.Pos(),
			s.Jets.Pos(),
		},
	}

	for count+1 < length {
		s.run(1)
		count++

		signature := s.Tower.Signature()
		match := false
		for prevNum, prev := range store {
			match = prev.shapePos == s.Shapes.Pos() &&
				prev.jetPos == s.Jets.Pos() &&
				signature.Match(prev.signature)

			if match {
				// Super maths time!
				pattenLength := count - uint64(prevNum)  //number of cycles between matches
				pattenSize := s.Tower.Size() - prev.size //high diff between matches

				remainder := length - count                      //how many cycles are left
				skippedCycles := remainder / pattenLength        //how may pattens we can skip
				skippedRemainder := remainder % pattenLength     //how many cycles will be left after skip
				skippedSize = uint64(pattenSize) * skippedCycles //how much hight we skipped up

				count = length - skippedRemainder //skip forward in time
				break
			}
		}

		if match {
			break
		} else {
			store = append(store, struct {
				signature        shapes.Patten
				size             int
				shapePos, jetPos int
			}{
				s.Tower.Signature(),
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
