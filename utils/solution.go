package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type SolutionBetter[T_1 constraints.Ordered, T_2 constraints.Ordered] interface {
	SolvePartOne([]string) T_1
	ExampleSolutionPartOne() T_1
	SolvePartTwo([]string) T_2
	ExampleSolutionPartTwo() T_2
}

func printSection[T constraints.Ordered](name string, exSln T, exCalc T, inputCalc T) {
	fmt.Println(name)

	fmt.Println("")
	fmt.Println("expected")
	fmt.Println(exSln)
	fmt.Println("")
	fmt.Println("calculated")
	fmt.Println(exCalc)
	fmt.Println("")
	fmt.Println("pass")
	fmt.Println(exSln == exCalc)
	fmt.Println("")
	fmt.Println("input")
	fmt.Println(inputCalc)

}

func Solve[T_1 constraints.Ordered, T_2 constraints.Ordered](s SolutionBetter[T_1, T_2]) {
	exLines, inputLines := ReadExample(), ReadInput()

	exCalc1, inputCalc1 := s.SolvePartOne(exLines), s.SolvePartOne(inputLines)
	printSection("Part 1", s.ExampleSolutionPartOne(), exCalc1, inputCalc1)
	fmt.Println("-----")

	exCalc2, inputCalc2 := s.SolvePartTwo(exLines), s.SolvePartTwo(inputLines)
	printSection("Part 2", s.ExampleSolutionPartTwo(), exCalc2, inputCalc2)

}

type Solution[T constraints.Ordered] interface {
	Solve([]string) T
	ExampleSolution() T
}

func SolveOld[T constraints.Ordered](p Solution[T]) {
	exLines, inputLines := ReadExample(), ReadInput()

	exAnswer, inputAnswer := p.Solve(exLines), p.Solve(inputLines)

	fmt.Println("expected example answer:", p.ExampleSolution())
	fmt.Println("calculated example answer:", exAnswer)
	fmt.Println("example pass:", exAnswer == p.ExampleSolution())

	fmt.Println("---")

	fmt.Println("input answer")
	fmt.Println(inputAnswer)
}
