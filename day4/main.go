package main

import (
	"aoc-2022/utils"
	"strings"
)

// https://adventofcode.com/2022/day/4
type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 2
}
func (Sln) SolvePartOne(lines []string) int {
	sectionPairs := getAllSectionPairs(lines)
	sectionsWithSub := utils.Filter(sectionPairs, func(pair []section) bool {
		return sectionIsSub(pair[0], pair[1]) || sectionIsSub(pair[1], pair[0])
	})
	return len(sectionsWithSub)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 4
}

func (Sln) SolvePartTwo(lines []string) int {
	sectionPairs := getAllSectionPairs(lines)
	sectionsWithOverlap := utils.Filter(sectionPairs, func(pair []section) bool {
		return sectionsHaveOverlap(pair[0], pair[1])
	})
	return len(sectionsWithOverlap)
}

func main() {

	utils.Solve[int, int](Sln{})

}

type section struct {
	start int
	end   int
}

func (s section) nums() []int {
	a := []int{}
	for i := s.start; i <= s.end; i++ {
		a = append(a, i)
	}
	return a
}

func newSectionFromId(id string) section {
	components := strings.Split(id, "-")
	return section{start: utils.ParseIntMust(components[0]), end: utils.ParseIntMust(components[1])}
}

func getAllSectionPairs(lines []string) [][]section {
	return utils.MapFn(lines, func(s string) []section {
		splt := strings.Split(s, ",")
		return []section{
			newSectionFromId(splt[0]),
			newSectionFromId(splt[1]),
		}
	})
}

func sectionIsSub(sub section, check section) bool {
	if sub.start < check.start {
		return false
	}
	if sub.end > check.end {
		return false
	}
	return true
}

func sectionsHaveOverlap(a section, b section) bool {
	aNums := a.nums()
	bNums := b.nums()

	return len(utils.Intersect(aNums, bNums)) > 0
}
