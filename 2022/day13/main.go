package main

import (
	"aco/common"
	"aoc/day13/packet"
	"embed"
	"flag"
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

//go:embed input.txt
var sample embed.FS

var verbose *bool

func vPrintf(template string, a ...any) {
	if *verbose {
		fmt.Printf(template, a...)
	}
}

func main() {
	file := flag.String("f", "-", "File")
	verbose = flag.Bool("v", false, "Verbose")
	flag.Parse()

	var lines <-chan string

	if *file == "-" {
		if f, err := sample.Open("input.txt"); err == nil {
			lines = common.ReadInputEmbed(f)
		} else {
			panic(err)
		}
	} else {
		lines = common.ReadInputPath(*file)
	}

	dividers := [2]packet.Packet{
		parse("[[2]]"),
		parse("[[6]]"),
	}
	allPackets := []packet.Packet{
		dividers[0].DeepCopy(),
		dividers[1].DeepCopy(),
	}
	part1 := 0
	index := 0
	var left, right string
	var open bool
	for {
		index++
		left = <-lines
		right, open = <-lines

		if !open {
			break
		}

		lPack, rPack := parse(left), parse(right)
		ordered := packet.Ordered(lPack, rPack)
		if ordered {
			part1 += index
		}
		vPrintf("%v: Ordered(%v, %v) => %v\n", index, lPack, rPack, ordered)

		allPackets = append(allPackets, []packet.Packet{lPack, rPack}...)
	}

	fmt.Println("part1:", part1)

	slices.SortFunc(allPackets, func(a, b packet.Packet) bool { return packet.Ordered(a, b) })
	codes := [2]int{}
	for i, p := range allPackets {
		found := -1
		for j, div := range dividers {
			if packet.Equals(p, div) {
				found = j
			}
		}
		if found >= 0 {
			vPrintf("***%v***\n", p)
			codes[found] = i + 1
		} else {
			vPrintf("%v\n", p)
		}
	}
	fmt.Printf("part2 codes: %v\n", codes)
	fmt.Printf("part2 answer: %v\n", codes[0]*codes[1])
}

func parse(input string) packet.Packet {
	packet, _ := parsePacket(input)

	if fmt.Sprint(packet) != input {
		panic(fmt.Sprintf("Packet %v does not match input %v", packet, input))
	}

	return packet
}

func parsePacket(input string) (packet.Packet, string) {
	switch input[0] {
	case '[':
		// vPrintf("parsing array: %v\n", input)
		return parseArray(input[1:])
	default:
		// vPrintf("parsing value: %v\n", input)
		return parseValue(input)
	}
}

func parseArray(input string) (packet.ArrayPacket, string) {
	packets := []packet.Packet{}
	for {
		switch rune(input[0]) {
		case ']':
			// vPrintf("Ending Array: %v\n", input)
			input = input[1:]
			return packet.CreatePacketArray(packets...), input
		case ',':
			// vPrintf("Continueing to next element: %v\n", input)
			input = input[1:]
			continue
		default:
			// vPrintf("Parsing next element: %v\n", input)
			var p packet.Packet
			p, input = parsePacket(input)
			packets = append(packets, p)
		}
	}
}

func parseValue(input string) (packet.ValuePacket, string) {
	vString, rString := cutString(input, ',', ']')
	vInt, _ := strconv.Atoi(vString)

	return packet.CreateValuePacket(vInt), rString
}

func cutString(input string, stops ...rune) (string, string) {
	for i := range input {
		for _, s := range stops {
			if rune(input[i]) == s {
				return input[:i], input[i:]
			}
		}
	}
	return "", input
}
