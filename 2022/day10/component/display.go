package component

import (
	"aoc/day10/pixel"
	"time"
)

type display struct {
	width, lines, pos int
	rows              []pixel.Row

	inReg <-chan int
	out   chan<- string

	internalClock
}

func (d *display) Run(clk <-chan time.Time) (stop func(), err error) {
	action := func(_ time.Time) {
		in := <-d.inReg

		row := d.rows[d.rowNum()]
		d.rows[d.rowNum()] = row | pixel.Sprite(in).Focus(d.rowPos())

		d.nextPos()
		if d.pos == 0 {
			d.out <- d.draw()
		}
	}

	return d.onTick(clk, action)
}

func Display(width, lines int, inReg <-chan int, out chan<- string) Component {

	return &display{
		width: width,
		lines: lines,
		rows:  make([]pixel.Row, lines),

		inReg: inReg,
		out:   out,
	}
}

func (d *display) nextPos() {
	d.pos++
	if d.pos >= d.width*d.lines {
		d.pos = 0
	}
}

func (d *display) rowNum() int {
	return d.pos / d.width
}

func (d *display) rowPos() int {
	return d.pos - (d.rowNum() * d.width)
}

func (d *display) draw() string {
	screen := ""
	for y := -1; y <= d.lines; y++ {
		screen += "|"
		for x := 0; x < d.width; x++ {
			if y == -1 || y == d.lines {
				screen += "-"
				continue
			}

			if d.rows[y].Focus(x) > 0 {
				screen += "#"
			} else {
				screen += " "
			}
		}
		screen += "|\n"
	}
	return screen
}

func (d *display) String() string { return d.draw() }
