package mountain

import "fmt"

type state struct {
	pos  Coord
	time int
}

func FindShortestPath(m Mountain, start, end Coord, time int) int {
	var (
		period       = m.Height * m.Width
		hazzCache    = make([]map[Coord]int, 0, period)
		getHazzState = func(t int) map[Coord]int {
			for t+1 > len(hazzCache) {
				hazz := map[Coord]int{}
				for _, b := range m.Blizzards {
					hazz[b.Pos]++
				}
				hazzCache = append(hazzCache, hazz)
				m.Next()
			}
			return hazzCache[t]
		}
	)

	var (
		sState  = state{start, time}
		queue   = []state{sState}
		seen    = map[state]interface{}{sState: struct{}{}}
		current state
	)

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, adj := range current.pos.Adjecent() {
			isStart, isEnd := adj == start, adj == end
			inBounds := isStart || isEnd ||
				(adj.X >= 0 && adj.X < m.Width &&
					adj.Y >= 0 && adj.Y < m.Height)

			if !inBounds {
				continue
			}

			if isEnd {
				return current.time + 1
			} else {
				var (
					nTime  = current.time + 1
					pTime  = nTime % period
					nState = state{adj, nTime}
					pState = state{adj, pTime}
				)

				if getHazzState(pTime)[adj] > 0 {
					continue
				}

				if _, hasSeen := seen[pState]; !hasSeen {
					seen[nState] = struct{}{}
					queue = append(queue, state{adj, current.time + 1})
				}
			}
		}
	}

	fmt.Println("No route found. Seen:", len(seen))
	return 0
}
