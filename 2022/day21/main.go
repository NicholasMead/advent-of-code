package main

import (
	"aco/common"
	"fmt"
	"strconv"
	"strings"
)

type Jobs map[string]string

func (j Jobs) Copy() Jobs {
	copy := Jobs{}
	for mon, job := range j {
		copy[mon] = job
	}
	return copy
}

func parseJob(job string) (left, right, opp string) {
	parts := strings.Split(job, " ")

	return parts[0], parts[2], parts[1]
}

func (j Jobs) GetValue(monkey string) (val int, found bool) {
	job, found := j[monkey]
	if !found {
		return 0, false
	}
	if num, err := strconv.Atoi(job); err == nil {
		return num, true
	}

	leftMon, rightMon, opp := parseJob(job)
	left, lFound := j.GetValue(leftMon)
	right, rFound := j.GetValue(rightMon)

	if !(lFound && rFound) {
		return 0, false
	}

	switch opp {
	case "+":
		return left + right, true
	case "-":
		return left - right, true
	case "*":
		return left * right, true
	case "/":
		return left / right, true
	default:
		panic("Unknown opp: " + opp)
	}
}

func (jobs Jobs) GetHumn(monkey string, target int) int {
	job, found := jobs[monkey]
	if !found {
		return target
	}

	lMonkey, rMonkey, opp := parseJob(job)
	lVal, lFound := jobs.GetValue(lMonkey)
	rVal, rFound := jobs.GetValue(rMonkey)

	var missingMon string
	var knownVal int

	if !lFound {
		missingMon = lMonkey
		knownVal = rVal
	} else if !rFound {
		missingMon = rMonkey
		knownVal = lVal
	} else {
		panic("Unable to make calc!")
	}

	switch opp {
	case "+": // missing + known = target
		return jobs.GetHumn(missingMon, target-knownVal)
	case "*": // missing * known = target
		return jobs.GetHumn(missingMon, target/knownVal)
	case "-": // left - right = target
		if missingMon == lMonkey {
			return jobs.GetHumn(missingMon, target+rVal)
		} else {
			return jobs.GetHumn(missingMon, lVal-target)
		}
	case "/": // left / right = target
		if missingMon == lMonkey {
			return jobs.GetHumn(missingMon, target*rVal)
		} else {
			return jobs.GetHumn(missingMon, lVal/target)
		}
	default:
		panic("Unknown opp: " + opp)
	}
}

func main() {
	common.TimedMs(func() {
		jobs := parseInput()

		fmt.Println("part1:", part1(jobs))
		fmt.Println("part2:", part2(jobs))
	})
}

func part1(jobs Jobs) int {
	root, _ := jobs.GetValue("root")
	return root
}

func part2(jobs Jobs) int {
	jobs = jobs.Copy()
	root := jobs["root"]
	delete(jobs, "root")
	delete(jobs, "humn")

	lMonkey, rMonkey, _ := parseJob(root)
	lVal, lFound := jobs.GetValue(lMonkey)
	rVal, rFound := jobs.GetValue(rMonkey)

	if !lFound {
		return jobs.GetHumn(lMonkey, rVal)
	}
	if !rFound {
		return jobs.GetHumn(rMonkey, lVal)
	}
	panic("puzzel broken...")
}

func parseInput() Jobs {
	flags := common.GetFlags()
	input := common.ReadInputPath(*flags.File)

	jobs := Jobs{}

	for line := range input {
		sep := strings.Split(line, ": ")
		jobs[sep[0]] = sep[1]
	}

	return jobs
}
