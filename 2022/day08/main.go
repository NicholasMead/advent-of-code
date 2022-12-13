package main

import (
	"aco/common"
	"aoc/day8/grid"
	"fmt"
	"os"
)

func main() {
	input := common.ReadInputPath(os.Args[1])
	verbose := len(os.Args) > 2 && os.Args[2] == "-v"
	trees := [][]int{}

	x := 0
	for input := range input {
		trees = append(trees, make([]int, len(input)))

		for y := range trees[x] {
			trees[x][y] = int(rune(input[y]) - '0')
			if verbose {
				fmt.Print(trees[x][y])
			}
		}
		x++
		if verbose {
			fmt.Print("\n")
		}
	}

	grid := grid.NewGrid(trees)
	fmt.Println("Pased grid of size (w,h): ", grid.Width(), grid.Height())

	treeCount := 0
	count := 0
	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			treeCount++
			if grid.IsVisible(x, y) {
				count++
				if verbose {
					fmt.Print(1)
				}
			} else {
				if verbose {
					fmt.Print(0)
				}
			}
		}
		if verbose {
			fmt.Print("\n")
		}
	}

	fmt.Println("Part 1: Of ", treeCount, "trees, ", count, "are visible")

	maxScenic := 0
	for x := 0; x < grid.Width(); x++ {
		for y := 0; y < grid.Height(); y++ {
			scenic := grid.SceanicScore(x, y)
			if verbose {
				fmt.Print(scenic)
				fmt.Print("\t")
			}

			if scenic > maxScenic {
				maxScenic = scenic
			}
		}
		if verbose {
			fmt.Print("\n")
		}
	}

	fmt.Println("Part 2:", maxScenic)
}
