package main

import (
	"aco/common"
	"aoc/day7/mockfs"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := common.ReadInputPath(os.Args[1])
	fs := ParseFileSystem(input)

	fmt.Println("part 1:", part1(fs))
	fmt.Println("part 2:", part2(fs))
}

func part2(fs mockfs.FileSystem) int {
	fs.Move("/")
	disk := 70000000
	current := fs.CurrentDir().GetSize()
	remaining := disk - current
	update := 30000000
	required := update - remaining

	areBigEnough := func(d mockfs.Directory) bool {
		return d.GetSize() > required
	}
	candidates := fs.FindAll(areBigEnough)

	output := current

	for c := range candidates {
		if c.GetSize() < output {
			output = c.GetSize()
		}
	}

	return output
}

func part1(fs mockfs.FileSystem) int {
	total := 0

	selector := func(d mockfs.Directory) bool {
		return d.GetSize() <= 100000
	}

	for d := range fs.FindAll(selector) {
		total += d.GetSize()
	}

	return total
}

func ParseFileSystem(input <-chan string) mockfs.FileSystem {
	fs := mockfs.CreateFileSystem()

	for line := range input {
		cmd := strings.Split(line, " ")
		switch cmd[0] {
		case "$":
			switch cmd[1] {
			case "cd":
				fs.Move(cmd[2])
			}
		case "dir":
			dir := mockfs.CreateDirectory(cmd[1])
			fs.CurrentDir().AddDirectory(dir)
		default:
			size, _ := strconv.Atoi(cmd[0])
			fs.CurrentDir().AddFile(cmd[1], size)
		}
	}

	return fs
}
