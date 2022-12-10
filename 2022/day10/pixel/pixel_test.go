package pixel

import (
	"fmt"
	"testing"
)

func TestPixel(t *testing.T) {
	cases := map[int]Row{
		0:  0b1,
		1:  0b10,
		2:  0b100,
		3:  0b1000,
		8:  0b100000000,
		10: 0b10000000000,
		11: 0b100000000000,
	}

	for i, r := range cases {
		name := fmt.Sprintf("Pixel(%d)=>%b", i, r)
		t.Run(name, func(t *testing.T) {
			if pixel := Pixel(i); pixel != r {
				t.Errorf("%b", pixel)
			}
		})

	}
}

func TestSprite(t *testing.T) {
	cases := map[int]Row{
		0:  0b11,
		1:  0b111,
		2:  0b1110,
		3:  0b11100,
		8:  0b1110000000,
		10: 0b111000000000,
		11: 0b1110000000000,
	}

	for i, r := range cases {
		name := fmt.Sprintf("Sprite(%d)=>%b", i, r)
		t.Run(name, func(t *testing.T) {
			if sprite := Sprite(i); sprite != r {
				t.Errorf("%b", sprite)
			}
		})

	}
}

func TestRow(t *testing.T) {
	t.Run("Pos", func(t *testing.T) {
		cases := []struct {
			row Row
			pos int
			res Row
		}{
			{0b11100, 0, 0},
			{0b11100, 1, 0},
			{0b11100, 2, 0b100},
			{0b11100, 4, 0b10000},
			{0b11100, 5, 0},
		}

		for _, c := range cases {
			name := fmt.Sprintf("[%b]pos(%d)=>(%b)", c.row, c.pos, c.res)
			t.Run(name, func(t *testing.T) {
				if r := c.row.Focus(c.pos); r != c.res {
					t.Errorf("%b", r)
				}
			})
		}

	})
	cases := map[int]Row{
		0:  0b11,
		1:  0b111,
		2:  0b1110,
		3:  0b11100,
		8:  0b1110000000,
		10: 0b111000000000,
		11: 0b1110000000000,
	}

	for i, r := range cases {
		name := fmt.Sprintf("Pixel %d => %b", i, r)
		t.Run(name, func(t *testing.T) {
			if sprite := Sprite(i); sprite != r {
				t.Errorf("%b", sprite)
			}
		})

	}
}
