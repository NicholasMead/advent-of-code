package main

import (
	"aco/common"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type State struct {
	timeRemaining                   int
	ore, clay, obs, geo             int
	oreBot, clayBot, obsBot, geoBot int
}

type blueprint struct {
	id                      int
	oreCost                 int
	clayCost                int
	obsOreCost, obsClayCost int
	geoOreCost, geoObsCost  int
}

func (s State) nextStates(bp blueprint) []State {
	states := make([]State, 1, 5)
	states[0] = State{
		timeRemaining: s.timeRemaining - 1,
		ore:           s.ore + s.oreBot,
		clay:          s.clay + s.clayBot,
		obs:           s.obs + s.obsBot,
		geo:           s.geo + s.geoBot,
		oreBot:        s.oreBot,
		clayBot:       s.clayBot,
		obsBot:        s.obsBot,
		geoBot:        s.geoBot,
	}

	if s.ore >= bp.oreCost {
		states = append(states, State{
			timeRemaining: s.timeRemaining - 1,
			ore:           s.ore + s.oreBot - bp.oreCost,
			clay:          s.clay + s.clayBot,
			obs:           s.obs + s.obsBot,
			geo:           s.geo + s.geoBot,
			oreBot:        s.oreBot + 1,
			clayBot:       s.clayBot,
			obsBot:        s.obsBot,
			geoBot:        s.geoBot,
		})
	}
	if s.ore >= bp.clayCost {
		states = append(states, State{
			timeRemaining: s.timeRemaining - 1,
			ore:           s.ore + s.oreBot - bp.clayCost,
			clay:          s.clay + s.clayBot,
			obs:           s.obs + s.obsBot,
			geo:           s.geo + s.geoBot,
			oreBot:        s.oreBot,
			clayBot:       s.clayBot + 1,
			obsBot:        s.obsBot,
			geoBot:        s.geoBot,
		})
	}
	if s.ore >= bp.obsOreCost && s.clay >= bp.obsClayCost {
		states = append(states, State{
			timeRemaining: s.timeRemaining - 1,
			ore:           s.ore + s.oreBot - bp.obsOreCost,
			clay:          s.clay + s.clayBot - bp.obsClayCost,
			obs:           s.obs + s.obsBot,
			geo:           s.geo + s.geoBot,
			oreBot:        s.oreBot,
			clayBot:       s.clayBot,
			obsBot:        s.obsBot + 1,
			geoBot:        s.geoBot,
		})
	}
	if s.ore >= bp.geoOreCost && s.obs >= bp.geoObsCost {
		states = append(states, State{
			timeRemaining: s.timeRemaining - 1,
			ore:           s.ore + s.oreBot - bp.geoOreCost,
			clay:          s.clay + s.clayBot,
			obs:           s.obs + s.obsBot - bp.geoObsCost,
			geo:           s.geo + s.geoBot,
			oreBot:        s.oreBot,
			clayBot:       s.clayBot,
			obsBot:        s.obsBot,
			geoBot:        s.geoBot + 1,
		})
	}

	return states
}

func (s State) reduce(bp blueprint) State {
	// prevState := s
	maxOre := max(bp.oreCost, bp.clayCost, bp.obsOreCost, bp.geoOreCost) * (s.timeRemaining)
	if s.ore >= maxOre {
		s.ore = maxOre
		s.oreBot = 0
	}
	maxClay := bp.obsClayCost * (s.timeRemaining)
	if s.clay >= maxClay {
		s.clay = maxClay
		s.clayBot = 0
	}
	maxObs := bp.geoObsCost * (s.timeRemaining)
	if s.obs >= maxObs {
		s.obs = maxObs
		s.obsBot = 0
	}

	// if s != prevState {
	// 	fmt.Println("Reduced", prevState, "to", s)
	// }
	return s
}

func max(v ...int) int {
	m := math.MinInt
	for _, v := range v {
		if v > m {
			m = v
		}
	}
	return m
}

func best(bp blueprint, time int) int {
	queue := []State{{
		timeRemaining: time,
		oreBot:        1,
	}}
	visited := map[State]interface{}{}
	best := 0

	for len(queue) > 0 {
		var state State
		state, queue = queue[0], queue[1:]

		if len(visited) > 0 && len(visited)%1_000_000 == 0 {
			fmt.Printf("%v: %v, %v, v:%v, q:%v\n", bp.id, state, best, len(visited), len(queue))
		}

		if state.geo > best {
			best = state.geo
		}

		if state.timeRemaining > 0 {
			next := state.nextStates(bp)
			for _, n := range next {
				n = n.reduce(bp)
				if _, v := visited[n]; !v {
					visited[n] = new(interface{})
					queue = append(queue, n)
				}
			}
		}
	}

	return best
}

type result struct {
	num, score int
}

func main() {
	blueprints := getBlueprints()

	var p1, p2 int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { p1 = part1(blueprints); wg.Done() }()
	go func() { p2 = part2(blueprints); wg.Done() }()

	wg.Wait()
	fmt.Println("part1:", p1)
	fmt.Println("part2:", p2)
}

func part2(blueprints []blueprint) int {
	results := make(chan result, 3)
	var wg sync.WaitGroup

	run := func(i int, bp blueprint, time int) {
		score := best(bp, time)
		results <- result{i, score}
		wg.Done()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go run(i+1, blueprints[i], 32)
	}

	wg.Wait()
	close(results)
	part2 := 1
	for r := range results {
		fmt.Printf("best %v: %v\n", r.num, r.score)
		part2 *= r.score
	}
	return part2
}

func part1(blueprints []blueprint) int {
	results := make(chan result, len(blueprints))
	var wg sync.WaitGroup

	run := func(i int, bp blueprint, time int) {
		score := best(bp, time)
		results <- result{i, score}
		wg.Done()
	}

	for i, bp := range blueprints {
		fmt.Println(i+1, bp)
		wg.Add(1)
		go run(i+1, bp, 24)
	}

	wg.Wait()
	close(results)
	part1 := 0
	for r := range results {
		fmt.Printf("best %v: %v\n", r.num, r.score)
		part1 += r.num * r.score
	}
	return part1
}

func getBlueprints() []blueprint {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	id := 1
	blueprints := []blueprint{}
	for line := range input {
		line = strings.ReplaceAll(line, "  ", " ")
		words := strings.Split(line, " ")
		blueprints = append(blueprints, blueprint{
			id:          id,
			oreCost:     toInt(words[6]),
			clayCost:    toInt(words[12]),
			obsOreCost:  toInt(words[18]),
			obsClayCost: toInt(words[21]),
			geoOreCost:  toInt(words[27]),
			geoObsCost:  toInt(words[30]),
		})
		id++
	}
	return blueprints
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
