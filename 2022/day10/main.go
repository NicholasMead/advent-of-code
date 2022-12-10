package main

import (
	"aco/common"
	"aoc/day10/component"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/* Circuit design:
[ReadInput]
 | file_parse
[parseInput]              tick() > [clock]
 | parse_reg                         | clk_split
[register] ------------------------[Splitter(3)]
 | reg_split(1)                      |
[spliter(2)]                         |
 # -------------#                    |
 | split_disp   | split_buffer       |
[display] ---- [memBuffer] ----------#
 | disp_out(1)
*/

func main() {
	//flags
	verbose := flag.Bool("v", false, "Extra printing")
	file := flag.String("f", "-", "File Name")
	flag.Parse()

	//channels (wires)
	file_parse := common.ReadInput(*file)
	parse_reg := parseInput(file_parse)
	reg_split := make(chan int, 1)
	split := common.Split(reg_split, 2)
	split_disp := split[0]
	split_buffer := split[1]
	disp_out := make(chan string, 1)

	//clocks (for sync)
	clk_split, tick := component.Ticker()
	clk_out := common.Split(clk_split, 3)

	//component (for work)
	reg := component.Register(parse_reg, reg_split)
	disp := component.Display(40, 6, split_disp, disp_out)
	logger := component.MemBuffer(split_buffer)

	for i, c := range []component.Component{reg, disp, logger} {
		c.Run(clk_out[i])
	}

	// Run untill fist display print
	var firstDraw string
	ticks := 0
	for firstDraw == "" {
		select {
		case firstDraw = <-disp_out:
			vPrintln(verbose, "Execution Complete")
		default:
			ticks++
			vPrintln(verbose, "tick", ticks, reg)
			vPrintln(verbose, disp)
			tick()
		}
	}

	//Get AoC outputs
	part1 := 0
	log := logger.Peak()
	for _, c := range []int{20, 60, 100, 140, 180, 220} {
		vPrintf(verbose, "%d: %d\n", c, log[c-1])
		part1 += (c * log[c-1])
	}
	fmt.Println("ticks:", ticks)
	fmt.Println("part1:", part1)
	fmt.Print("part2:\n", firstDraw)
}

func parseInput(input <-chan string) <-chan component.Command {
	cmds := make(chan component.Command, 1024)

	go func() {
		defer close(cmds)
		for line := range input {
			split := strings.Split(line, " ")
			cmd := component.Command{
				Program: component.Program(split[0]),
				Arg:     0,
			}

			if len(split) > 1 {
				cmd.Arg, _ = strconv.Atoi(split[1])
			}

			cmds <- cmd
		}
	}()

	return cmds
}

func vPrintln(v *bool, a ...any) {
	if *v {
		fmt.Println(a...)
	}
}

func vPrintf(verbose *bool, t string, a ...any) {
	if *verbose {
		fmt.Printf(t, a...)
	}
}
