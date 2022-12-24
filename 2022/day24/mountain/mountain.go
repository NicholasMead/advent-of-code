package mountain

type Coord struct{ X, Y int }

func (c Coord) Adjecent() [5]Coord {
	adj := [5]Coord{}
	pos := [5]Coord{{0, 0}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i, p := range pos {
		adj[i].X = c.X + p.X
		adj[i].Y = c.Y + p.Y
	}
	return adj
}

type Mountain struct {
	Height, Width int
	Blizzards     []Blizzard
}

func (m *Mountain) Next() {
	next := make([]Blizzard, 0, len(m.Blizzards))
	for _, bliz := range m.Blizzards {
		bliz.Next()

		if bliz.Pos.X == -1 {
			bliz.Pos.X = m.Width - 1
		} else if bliz.Pos.X >= m.Width {
			bliz.Pos.X = 0
		}

		if bliz.Pos.Y == -1 {
			bliz.Pos.Y = m.Height - 1
		} else if bliz.Pos.Y >= m.Height {
			bliz.Pos.Y = 0
		}

		next = append(next, bliz)
	}
	m.Blizzards = next
}

func ParseMountain(raw [][]rune) Mountain {
	h, w := len(raw), len(raw[0])
	m := Mountain{
		Height:    h,
		Width:     w,
		Blizzards: make([]Blizzard, 0),
	}

	for y, line := range raw {
		for x, tile := range line {
			switch tile {
			case '>':
				b := Blizzard{
					Pos: Coord{x, y},
					Dir: Coord{1, 0},
				}
				m.Blizzards = append(m.Blizzards, b)
			case '<':
				b := Blizzard{
					Pos: Coord{x, y},
					Dir: Coord{-1, 0},
				}
				m.Blizzards = append(m.Blizzards, b)
			case '^':
				b := Blizzard{
					Pos: Coord{x, y},
					Dir: Coord{0, -1},
				}
				m.Blizzards = append(m.Blizzards, b)
			case 'v':
				b := Blizzard{
					Pos: Coord{x, y},
					Dir: Coord{0, 1},
				}
				m.Blizzards = append(m.Blizzards, b)
			}
		}
	}
	return m
}

type Blizzard struct {
	Pos Coord
	Dir Coord
}

func (b *Blizzard) Next() {
	b.Pos.X += b.Dir.X
	b.Pos.Y += b.Dir.Y
}
