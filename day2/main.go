package main

import (
	"aoc-2022/utils"
	"strings"
)

// https://adventofcode.com/2022/day/2
type Sln struct {
}

var _ utils.SolutionBetter[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 15
}
func (Sln) SolvePartOne(lines []string) int {
	return -1
}

func (Sln) ExampleSolutionPartTwo() int {
	return 0
}

func (Sln) SolvePartTwo(lines []string) int {
	return -1
}

func main() {

	utils.Solve[int, int](Sln{})

}

var responseMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

const (
	rock int = iota
	paper
	scissors
)

func getChoiceFromChar(c string) int {
	switch c {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors

	default:
		panic("invalid choice " + c)

	}
}
func getWinScore(opponent string, player string) int {
	oppenentChoice, playerChoice := getChoiceFromChar(opponent), getChoiceFromChar(player)
	if playerChoice == oppenentChoice {
		return 3
	}
	if oppenentChoice == rock {
		if playerChoice == paper {
			return 6
		}
		return 0
	}

	if oppenentChoice == paper {
		if playerChoice == scissors {
			return 6
		}
		return 0
	}

	if oppenentChoice == scissors {
		if playerChoice == rock {
			return 6
		}
		return 0
	}

	panic("unable to calc win score")
}

func calcLineScore(l string) int {
	choices := strings.Split(l, " ")
	return getWinScore(choices[0], choices[1]) + responseMap[choices[1]]
}
