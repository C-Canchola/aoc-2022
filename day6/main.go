package main

import (
	"aoc-2022/utils"
)

type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 7
}

func (Sln) SolvePartOne(lines []string) int {
	l := lines[0]
	i := 0
	for {
		sub := l[i : i+4]
		if len(utils.UniqueMap([]byte(sub))) != 4 {
			i += 1
			continue
		}

		return i + 4
	}
}

func (Sln) ExampleSolutionPartTwo() int {
	return 19
}

func (Sln) SolvePartTwo(lines []string) int {
	l := lines[0]
	i := 0
	for {
		sub := l[i : i+14]
		if len(utils.UniqueMap([]byte(sub))) != 14 {
			i += 1
			continue
		}

		return i + 14
	}
}

func main() {

	utils.Solve[int, int](Sln{})

}
