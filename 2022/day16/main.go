package main

import (
	"aco/common"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flows, edges := parseInput()
	computeAllEdges(edges)

	player := maxPressure("AA", flows, edges, 30)
	fmt.Println("part1:", player.score, player)

	duo := duoMaxPressure("AA", flows, edges, 26)
	fmt.Println("part2:", duo.Score(), duo)
}

type FlowMap map[string]int
type EdgeMap map[string]map[string]int
type RouteMap map[string]map[string]Route

func EmptyRouteMap(ids ...string) RouteMap {
	empty := RouteMap{}
	for _, i := range ids {
		empty[i] = map[string]Route{}
		for _, j := range ids {
			empty[i][j] = Route{}
		}
	}
	return empty
}

func (r RouteMap) String() string {
	s := ""
	for i := range r {
		s += fmt.Sprintf("%v: %v\n", i, r[i])
	}
	return s
}

type Route struct {
	length, score int
	path          string
}

func (r Route) pos() string {
	nodes := strings.Split(r.path, "+")
	return nodes[len(nodes)-1]
}

func (r Route) visits() []string {
	return strings.Split(r.path, "+")
}

func parseInput() (FlowMap, EdgeMap) {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	flows := FlowMap{}
	dist := EdgeMap{}

	valveExp := regexp.MustCompile(`Valve (\w{2}) has flow rate=(\d+); tunnels? leads? to valves? `)
	connectionExp := regexp.MustCompile(`\w{2}`)

	for line := range input {
		valveMatch := valveExp.FindAllStringSubmatch(line, 1)

		id := valveMatch[0][1]
		flow, _ := strconv.Atoi(valveMatch[0][2])
		flows[id] = flow
		dist[id] = map[string]int{id: 0}

		line = valveExp.ReplaceAllString(line, "")
		connectionMatch := connectionExp.FindAllString(line, -1)
		for _, match := range connectionMatch {
			dist[id][match] = 1
		}
	}

	return flows, dist
}

func computeAllEdges(edgeMap EdgeMap) {
	nodes := make([]string, 0, len(edgeMap))
	for node := range edgeMap {
		nodes = append(nodes, node)
	}

	for _, k := range nodes {
		for _, i := range nodes {
			for _, j := range nodes {
				var ij, ik, kj int
				var found bool
				ij, found = edgeMap[i][j]
				if !found {
					ij = math.MaxInt / 2
				}
				ik, found = edgeMap[i][k]
				if !found {
					ik = math.MaxInt / 2
				}
				kj, found = edgeMap[k][j]
				if !found {
					kj = math.MaxInt / 2
				}

				if ik+kj < ij {
					edgeMap[i][j] = ik + kj
				}
			}
		}
	}
}

type Duo struct {
	player, elephant Route
}

func (d Duo) Score() int {
	return d.player.score + d.elephant.score
}

func duoMaxPressure(start string, flows FlowMap, edgeMap EdgeMap, maxLength int) Duo {
	add := func(r Route, start, end string) Route {
		r.length += edgeMap[start][end] + 1
		r.score += flows[end] * (maxLength - r.length)
		r.path += "+" + end
		return r
	}

	flowNodes := []string{} //nodes with positive flow values
	for node, flow := range flows {
		if flow > 0 {
			flowNodes = append(flowNodes, node)
		}
	}
	best := Duo{}
	queue := []Route{{path: start}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		duo := Duo{current, maxPressure(start, flows, edgeMap, maxLength, current.visits()...)}

		if duo.Score() > best.Score() {
			best = duo
			fmt.Println(duo, len(queue))
		}

		for _, node := range flowNodes {
			if strings.Contains(current.path, node) {
				continue
			}
			next := add(current, current.pos(), node)
			if next.length > maxLength {
				continue
			}
			queue = append(queue, next)
		}
	}
	return best
}

func maxPressure(start string, flows FlowMap, edgeMap EdgeMap, maxLength int, exclude ...string) Route {
	add := func(r Route, start, end string) Route {
		r.length += edgeMap[start][end] + 1
		r.score += flows[end] * (maxLength - r.length)
		r.path += "+" + end
		return r
	}

	flowNodes := []string{} //nodes with positive flow values
	for node, flow := range flows {
		if flow > 0 {
			shouldExclude := false
			for _, exc := range exclude {
				if node == exc {
					shouldExclude = true
				}
			}
			if !shouldExclude {
				flowNodes = append(flowNodes, node)
			}
		}
	}
	best := Route{}
	queue := []Route{{path: start}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.score > best.score {
			best = current
		}

		for _, node := range flowNodes {
			if strings.Contains(current.path, node) {
				continue
			}
			next := add(current, current.pos(), node)
			if next.length > maxLength {
				continue
			}
			queue = append(queue, next)
		}
	}
	return best
}
