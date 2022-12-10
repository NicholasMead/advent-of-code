package signal_test

import (
	"aoc/day6/signal"
	"testing"
)

func TestReaderFindStart(t *testing.T) {
	smallCases := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
		"nppdvjthqldpwncqszvftbrmjlhg":      6,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
	}
	bigCases := map[string]int{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
		"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
		"nppdvjthqldpwncqszvftbrmjlhg":      23,
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
	}

	small := signal.CreateReader(4)
	large := signal.CreateReader(14)

	testCases(t, smallCases, small)
	testCases(t, bigCases, large)
}

func testCases(t *testing.T, cases map[string]int, reader signal.Reader) {
	for i, r := range cases {
		input := make(chan byte, len(i))
		for _, i := range i {
			input <- byte(i)
		}
		result, _ := reader.FindStart(input)

		if result != r {
			t.Errorf("Expected %d for %d for input %s", r, result, i)
		}
	}
}
