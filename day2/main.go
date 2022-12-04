package main

import (
	"aoc-2022/utils"
	"strings"
)

// https://adventofcode.com/2022/day/2
type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 15
}
func (Sln) SolvePartOne(lines []string) int {
	scores := utils.MapFn(lines, calcLineScorePart1)
	return utils.Sum(scores)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 12
}

func (Sln) SolvePartTwo(lines []string) int {
	scores := utils.MapFn(lines, calcLineScorePart2)
	return utils.Sum(scores)
}

func main() {

	utils.Solve[int, int](Sln{})

}

const (
	rock int = iota
	paper
	scissors
)

var choiceScores map[int]int = map[int]int{
	rock:     1,
	paper:    2,
	scissors: 3,
}

var loseMap map[int]int = map[int]int{
	scissors: paper,
	paper:    rock,
	rock:     scissors,
}

var winMap map[int]int = map[int]int{
	scissors: rock,
	paper:    scissors,
	rock:     paper,
}

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

func getOpAndPlayerChoice(l string) (int, int) {
	splt := strings.Split(l, " ")
	return getChoiceFromChar(splt[0]), getChoiceFromChar(splt[1])
}

func getWinScoreFromCode(oppenentChoice int, playerChoice int) int {
	if playerChoice == oppenentChoice {
		return 3
	}
	if playerChoice == winMap[oppenentChoice] {
		return 6
	}
	return 0
}

func calcLineScorePart1(l string) int {
	op, pl := getOpAndPlayerChoice(l)
	return getWinScoreFromCode(op, pl) + choiceScores[pl]
}

// rock -> lose. paper -> draw, scissors -> win
func getNewMove(op int, pl int) int {
	switch pl {

	case rock:
		return loseMap[op]

	case paper:
		return op

	case scissors:
		return winMap[op]

	default:
		panic("unable to get new move")
	}
}

func calcLineScorePart2(l string) int {
	op, pl := getOpAndPlayerChoice(l)
	newPl := getNewMove(op, pl)

	winScore := getWinScoreFromCode(op, newPl)
	choiceScore := choiceScores[newPl]
	score := winScore + choiceScore

	return score
}
