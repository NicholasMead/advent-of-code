package snafu

import (
	"fmt"
	"math"
	"strings"
)

func ItoS(i int) (a string) {
	start := i
	for n := 27; n >= 0; n-- {
		var (
			mag = pow(5, n)

			next = "0"
			diff = abs(i)

			opts = map[string]int{
				"2": mag * 2,
				"1": mag * 1,
				"0": 0,
				"-": mag * -1,
				"=": mag * -2,
			}
		)

		for opt, val := range opts {
			if abs(i-val) < diff {
				next = opt
				diff = abs(i - val)
			}
		}

		a += next
		i -= opts[next]
	}

	if i != 0 {
		err := fmt.Errorf("Unable to convert %v to Snafu", start)
		panic(err)
	} else {
		return strings.TrimLeft(a, "0")
	}
}

func StoI(a string) (i int) {
	for n := 0; n < len(a); n++ {
		var (
			mag = pow(5, len(a)-1-n)
			opp = 0
		)

		switch a[n] {
		case '2':
			opp = 2
		case '1':
			opp = 1
		case '0':
			opp = 0
		case '-':
			opp = -1
		case '=':
			opp = -2
		}
		i += mag * opp
	}
	return i
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
