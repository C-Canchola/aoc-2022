package main

import (
	"aoc-2022/utils"
	"strings"
)

type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 13
}
func (Sln) SolvePartOne(lines []string) int {
	moves := linesToDirections(lines)
	path := findTailPath(moves)
	m := map[pos]bool{}
	for _, p := range path {
		m[p] = true
	}
	return len(m)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 1
}

func (Sln) SolvePartTwo(lines []string) int {
	moves := linesToDirections(lines)
	path := findTailPathPart2(moves)
	m := map[pos]bool{}
	for _, p := range path {
		m[p] = true
	}
	return len(m)
}

func main() {

	utils.Solve[int, int](Sln{})

}

type direction struct {
	x     int
	y     int
	times int
}

func newDirection(s string) direction {
	splt := strings.Split(s, " ")
	x, y := 0, 0

	switch splt[0] {
	case "R":
		x = 1
	case "L":
		x = -1
	case "D":
		y = -1
	case "U":
		y = 1
	}

	times := utils.ParseIntMust(splt[1])

	return direction{x: x, y: y, times: times}
}

func linesToDirections(lines []string) []direction {
	return utils.MapFn(lines, newDirection)
}

type pos struct {
	x, y int
}

func getUpdatedTailPosition(tailPos pos, headPos pos) pos {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			checkPos := pos{x: headPos.x - i, y: headPos.y - j}
			if checkPos != tailPos {
				continue
			}
			return tailPos
		}
	}

	if headPos.x == tailPos.x {
		if headPos.y > tailPos.y {
			return pos{x: tailPos.x, y: tailPos.y + 1}
		}
		return pos{x: tailPos.x, y: tailPos.y - 1}
	}

	if headPos.y == tailPos.y {
		if headPos.x > tailPos.x {
			return pos{x: tailPos.x + 1, y: tailPos.y}
		}
		return pos{x: tailPos.x - 1, y: tailPos.y}
	}

	// diagonal move
	yOffset := -1
	if headPos.y > tailPos.y {
		yOffset = 1
	}
	xOffset := -1
	if headPos.x > tailPos.x {
		xOffset = 1
	}

	return pos{x: tailPos.x + xOffset, y: tailPos.y + yOffset}
}

func makeMove(headPos pos, tailPos pos, move direction, tailPositions []pos) (pos, pos, []pos) {
	curHead, curTail := headPos, tailPos
	for i := 0; i < move.times; i++ {
		curHead = pos{x: curHead.x + move.x, y: curHead.y + move.y}
		curTail = getUpdatedTailPosition(curTail, curHead)
		tailPositions = append(tailPositions, curTail)
	}
	return curHead, curTail, tailPositions
}

func findTailPath(moves []direction) []pos {
	head, tail := pos{}, pos{}
	path := []pos{}
	for _, m := range moves {
		head, tail, path = makeMove(head, tail, m, path)
	}
	return path
}

func makeMoveForChain(headPos pos, chain []pos, move direction, tailPositions []pos) (pos, []pos, []pos) {
	for i := 0; i < move.times; i++ {
		headPos = pos{x: headPos.x + move.x, y: headPos.y + move.y}
		for j := 0; j < len(chain); j++ {
			if j == 0 {
				chain[j] = getUpdatedTailPosition(chain[j], headPos)
				continue
			}
			chain[j] = getUpdatedTailPosition(chain[j], chain[j-1])
		}
		tailPositions = append(tailPositions, chain[len(chain)-1])
	}
	return headPos, chain, tailPositions
}

func findTailPathPart2(moves []direction) []pos {
	head := pos{}
	chain := make([]pos, 9)
	paths := []pos{}

	for _, m := range moves {
		head, chain, paths = makeMoveForChain(head, chain, m, paths)
	}

	return paths

}
