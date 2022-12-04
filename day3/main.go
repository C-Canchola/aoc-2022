package main

import (
	"aoc-2022/utils"
	"fmt"
)

// https://adventofcode.com/2022/day/3
type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 157
}
func (Sln) SolvePartOne(lines []string) int {
	priorities := utils.MapFn(lines, func(l string) int {
		mid := int(len(l) / 2)
		c1, c2 := l[:mid], l[mid:]
		return getPriorityFromStrings(c1, c2)
	})
	return utils.Sum(priorities)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 70
}

func (Sln) SolvePartTwo(lines []string) int {
	groupedLines := utils.Reduce(lines, nil, func(s string, groupedLines [][]string) [][]string {
		if groupedLines == nil {
			groupedLines = [][]string{{}}
		}

		lastIndex := len(groupedLines) - 1
		indexAdder := 0

		if len(groupedLines[lastIndex]) == 3 {
			groupedLines = append(groupedLines, []string{})
			indexAdder = 1
		}

		newLastIndex := lastIndex + indexAdder
		groupedLines[newLastIndex] = append(groupedLines[newLastIndex], s)

		return groupedLines

	})

	priorities := utils.MapFn(groupedLines, func(g []string) int {
		return getPriorityFromStrings(g...)
	})

	return utils.Sum(priorities)
}

func main() {
	fmt.Println(int('a'))
	utils.Solve[int, int](Sln{})

}

//https://adventofcode.com/2022/day/3

// returns numeric value for character.
// a-z -> 1-26
// A-Z -> 27-52
//
// A is lower byte than a
func getPriority(c byte) int {
	if c >= 'a' {
		return int(c) - int('a') + 1
	}
	return int(c) - int('A') + 27
}

// assumes only one common characters in provided strings
func getCommonCharacter(groupedText ...string) byte {
	groupedBytes := utils.MapFn(groupedText, func(s string) []byte {
		return []byte(s)
	})
	byteIntersect := utils.Intersect(groupedBytes...)
	return byteIntersect[0]
}

func getPriorityFromStrings(a ...string) int {
	commonChar := getCommonCharacter(a...)
	return getPriority(commonChar)
}
