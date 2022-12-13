package main

import (
	"aco/common"
	"aoc/day11/monkey"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

func parseMonkeys(input <-chan string) []monkey.Monkey {
	monkeys := []monkey.Monkey{}
	cache := []string{}
	for i := range input {
		isMonkey := regexp.MustCompile(`Monkey \d+:`).MatchString(i)
		if isMonkey && len(cache) > 0 {
			if monkey, err := monkey.ParseMonkey(strings.Join(cache, "\n")); err == nil {
				monkeys = append(monkeys, monkey)
			}
			cache = []string{i}
		} else {
			cache = append(cache, i)
		}
	}
	if len(cache) > 0 {
		if monkey, err := monkey.ParseMonkey(strings.Join(cache, "\n")); err == nil {
			monkeys = append(monkeys, monkey)
		}
	}
	return monkeys
}

func GetScores(rounds int, worryDiv uint64, monkeys []monkey.Monkey) []int {
	monkeyScore := map[monkey.Monkey]int{}
	base := monkey.CommonBase(monkeys)

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for {
				if item, next := monkey.PopNext(); next {
					item := monkey.Inspect(item) / worryDiv
					item = item % base
					target := monkey.GetTarget(item)
					monkeys[target].Push(item)
					monkeyScore[monkey]++
				} else {
					break
				}
			}
		}
	}
	fmt.Println(monkeyScore)
	scores := []int{}
	for _, s := range monkeyScore {
		scores = append(scores, s)
	}
	return scores
}

func main() {
	filePath := flag.String("f", "-", "File path")
	flag.Parse()

	//part 1
	{
		input := common.ReadInputPath(*filePath)
		monkeys := parseMonkeys(input)
		scores := GetScores(20, 3, monkeys)
		slices.Sort(scores)
		fmt.Println("part1:", scores[len(scores)-1]*scores[len(scores)-2])
	}

	//part 2
	{
		input := common.ReadInputPath(*filePath)
		monkeys := parseMonkeys(input)
		scores := GetScores(10000, 1, monkeys)
		slices.Sort(scores)
		fmt.Println("part2:", scores[len(scores)-1]*scores[len(scores)-2])
	}
}
