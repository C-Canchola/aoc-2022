package utils

import (
	"fmt"
	"time"

	"golang.org/x/exp/constraints"
)

type Solution[T_1 comparable, T_2 comparable] interface {
	SolvePartOne([]string) T_1
	ExampleSolutionPartOne() T_1
	SolvePartTwo([]string) T_2
	ExampleSolutionPartTwo() T_2
}

func solveTimed[T comparable](fn func([]string) T, lines []string) (T, time.Duration) {
	start := time.Now()
	v := fn(lines)
	end := time.Now()

	return v, end.Sub(start)
}

func printSection[T constraints.Ordered](name string, exSln T, exCalc T, inputCalc T, exDur time.Duration, inputDur time.Duration) {
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
	fmt.Println("example time")
	fmt.Println(exDur.Seconds())
	fmt.Println("input time")
	fmt.Println(inputDur.Seconds())
	fmt.Println("")
	fmt.Println("input")
	fmt.Println(inputCalc)

}

func Solve[T_1 constraints.Ordered, T_2 constraints.Ordered](s Solution[T_1, T_2]) {
	exLines, inputLines := ReadExample(), ReadInput()

	exCalc1, exCalc1Dur := solveTimed(s.SolvePartOne, exLines)
	inputCalc1, inputCalc1Dur := solveTimed(s.SolvePartOne, inputLines)

	// exCalc1, inputCalc1 := s.SolvePartOne(exLines), s.SolvePartOne(inputLines)
	printSection("Part 1", s.ExampleSolutionPartOne(), exCalc1, inputCalc1, exCalc1Dur, inputCalc1Dur)
	fmt.Println("-----")

	exCalc2, exCalc2Dur := solveTimed(s.SolvePartTwo, exLines)
	inputCalc2, inputCalc2Dur := solveTimed(s.SolvePartTwo, inputLines)

	printSection("Part 2", s.ExampleSolutionPartTwo(), exCalc2, inputCalc2, exCalc2Dur, inputCalc2Dur)

}
