package snafu

import (
	"fmt"
	"testing"
)

func TestStafu(t *testing.T) {

	cases := map[int]string{
		1:         "1",
		2:         "2",
		3:         "1=",
		4:         "1-",
		5:         "10",
		6:         "11",
		7:         "12",
		8:         "2=",
		9:         "2-",
		10:        "20",
		15:        "1=0",
		20:        "1-0",
		2022:      "1=11-2",
		12345:     "1-0---0",
		314159265: "1121-1110-1=0",

		1747: "1=-0-2",
		906:  "12111",
		198:  "2=0=",
		11:   "21",
		201:  "2=01",
		31:   "111",
		1257: "20012",
		32:   "112",
		353:  "1=-1=",
		107:  "1-12",
		37:   "122",
	}

	for i, s := range cases {
		forward := fmt.Sprintf("text-to-stafu:(%v>%v)", s, i)
		t.Run(forward, func(t *testing.T) {
			val := StoI(s)
			if val != i {
				t.Errorf("Got %v expected %v", val, i)
			}
		})
	}

	for i, s := range cases {

		backwards := fmt.Sprintf("stafu to text: %v -> %v", s, i)
		t.Run(backwards, func(t *testing.T) {
			val := ItoS(i)
			if val != s {
				t.Errorf("Got %v expected %v", val, s)
			}
		})
	}

}
