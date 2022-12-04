package main

import (
	"aoc-2022/utils"
	"strconv"
)

type Sln struct {
}

var _ utils.SolutionBetter[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 24000
}
func (Sln) SolvePartOne(lines []string) int {
	elves := getElves(lines)
	return maxCalories(elves)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 45000
}

func (Sln) SolvePartTwo(lines []string) int {
	elves := getElves(lines)
	return maxThreeCalories(elves)
}

func main() {

	utils.Solve[int, int](Sln{})

}

type elf struct {
	calories []int
	position int
}

func (e elf) totalCalories() int {
	return utils.Sum(e.calories)
}

// assumes end on valid number
func getElves(lines []string) []elf {
	elves := make([]elf, 0)

	curPosition := 1
	curCalories := []int{}

	for _, l := range lines {
		calorieCount, err := strconv.Atoi(l)
		if err != nil {
			newElf := elf{calories: curCalories, position: curPosition}
			elves = append(elves, newElf)

			curCalories = []int{}
			curPosition += 1

		}

		curCalories = append(curCalories, calorieCount)
	}
	newElf := elf{calories: curCalories, position: curPosition}
	elves = append(elves, newElf)
	return elves
}

func maxCalories(elves []elf) int {
	maxCals, _ := utils.FindMax(elves, func(e elf) int { return e.totalCalories() })
	return maxCals
}

func maxThreeCalories(elves []elf) int {
	utils.Sort(elves, func(e elf) int { return e.totalCalories() })
	maxThreeSum := 0
	for _, elf := range elves[len(elves)-3:] {
		maxThreeSum += elf.totalCalories()
	}
	return maxThreeSum
}
