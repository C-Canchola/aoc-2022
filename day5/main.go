package main

import (
	"aoc-2022/utils"
	"strings"
)

type Sln struct {
}

var _ utils.Solution[string, string] = Sln{}

func (Sln) ExampleSolutionPartOne() string {
	return "CMZ"
}
func (Sln) SolvePartOne(lines []string) string {
	cols, moves := convertInputToColsAndMoves(lines)
	for _, m := range moves {
		cols = makeMove(cols, m)
	}
	return strings.Join(utils.MapFn(cols, func(c col) string { return c.topLetter() }), "")
}

func (Sln) ExampleSolutionPartTwo() string {
	return "MCD"
}

func (Sln) SolvePartTwo(lines []string) string {
	cols, moves := convertInputToColsAndMoves(lines)
	for _, m := range moves {
		cols = makeMoveForPart2(cols, m)
	}
	return strings.Join(utils.MapFn(cols, func(c col) string { return c.topLetter() }), "")
}

func main() {

	utils.Solve[string, string](Sln{})

}

/*
parse two sections seperately
first:
	- need array of strings PER ROW
		- reason is need to denote columns with less than max number
	- the bottom row is irrelevant as long as formatting is consistent
	- 3 chars per letter + space
	- ends with no space
*/

func extractLetters(s string) []string {
	letters := []string{}
	for len(s) > 0 {
		l := s[1:2]
		if l == " " {
			l = ""
		}
		letters = append(letters, l)
		s = s[3:]
		if len(s) > 0 {
			s = s[1:]
		}
	}
	return letters
}

// takes in entire first second unecessary bottom row included
func extractRowOfLetters(lines []string) [][]string {
	return utils.MapFn(lines[:len(lines)-1], extractLetters)
}

type col struct {
	letters []string
}

func (c col) topLetter() string {
	return c.letters[len(c.letters)-1]
}

func convertRowLettersToCols(rowLetters [][]string) []col {
	cols := []col{}
	for i := 0; i < len(rowLetters[0]); i++ {
		letters := []string{}
		for j := 0; j < len(rowLetters); j++ {
			row := rowLetters[j]
			letter := row[i]
			if letter == "" {
				continue
			}
			letters = append([]string{letter}, letters...)
		}
		cols = append(cols, col{letters: letters})
	}
	return cols
}

type move struct {
	amount int
	from   int
	to     int
}

func convertLineToMove(s string) move {
	splt := strings.Split(s, " ")
	return move{
		amount: utils.ParseIntMust(splt[1]),
		from:   utils.ParseIntMust(splt[3]),
		to:     utils.ParseIntMust(splt[5]),
	}
}

func convertInputToColsAndMoves(lines []string) ([]col, []move) {

	colLines, moveLines := []string{}, []string{}
	pastColSection := false
	for _, l := range lines {
		if l == "" {
			pastColSection = true
			continue
		}

		if pastColSection {
			moveLines = append(moveLines, l)
			continue
		}

		colLines = append(colLines, l)
	}

	rowOfLetters := extractRowOfLetters(colLines)
	cols := convertRowLettersToCols(rowOfLetters)
	moves := utils.MapFn(moveLines, convertLineToMove)

	return cols, moves
}

func makeMove(cols []col, m move) []col {
	from, to := cols[m.from-1], cols[m.to-1]
	for i := 0; i < m.amount; i++ {
		l := from.letters[len(from.letters)-1]
		from.letters = from.letters[:len(from.letters)-1]
		to.letters = append(to.letters, l)
	}
	cols[m.from-1], cols[m.to-1] = from, to
	return cols
}

func makeMoveForPart2(cols []col, m move) []col {
	from, to := cols[m.from-1], cols[m.to-1]
	moveLetters := from.letters[len(from.letters)-m.amount:]
	from.letters = from.letters[:len(from.letters)-m.amount]
	to.letters = append(to.letters, moveLetters...)
	cols[m.from-1], cols[m.to-1] = from, to
	return cols
}
