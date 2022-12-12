package main

import (
	"aoc-2022/utils"
	"fmt"
	"strings"
)

type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 13140
}
func (Sln) SolvePartOne(lines []string) int {
	c := newCycleCounter()
	for _, l := range lines {
		c.processSignal(l)
	}
	return c.getSumOfCycles(20, 60, 100, 140, 180, 220)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 0
}

func (Sln) SolvePartTwo(lines []string) int {
	c := newCycleCounter()
	for _, l := range lines {
		c.processSignal(l)
	}
	fmt.Println(getImageString(c))
	return -1
}

func main() {

	utils.Solve[int, int](Sln{})

}

type cycleCounter struct {
	values    []int
	addValues []int
}

func newCycleCounter() cycleCounter {
	return cycleCounter{
		values:    []int{1},
		addValues: []int{1},
	}
}

func (c *cycleCounter) processSignal(s string) {
	last := c.values[len(c.values)-1]
	if s == "noop" {
		c.addValues = append(c.addValues, 0)
		c.values = append(c.values, last)
		return
	}

	intStr := strings.Split(s, " ")[1]
	v := utils.ParseIntMust(intStr)
	c.addValues = append(c.addValues, 0, v)
	addVal := last + v
	c.values = append(c.values, last, addVal)
}

func (c cycleCounter) extractCycleValue(n int) int {

	v := c.values[n-1]
	return v * n
}

func (c cycleCounter) getSumOfCycles(ns ...int) int {
	s := 0
	for _, n := range ns {
		s += c.extractCycleValue(n)
	}
	return s
}

func getImageString(c cycleCounter) string {
	i := 0
	rows := [][]string{}
	row := []string{}

	for range c.values {
		i++
		if len(row) == 40 {
			rows = append(rows, row)
			row = []string{}
		}
		pos := len(row)
		regX := c.values[i-1]
		c := "."
		if regX >= pos-1 && regX <= pos+1 {
			c = "#"
		}
		row = append(row, c)
	}
	rowStrs := utils.MapFn(rows, func(l []string) string {
		return strings.Join(l, "")
	})
	return strings.Join(rowStrs, "\n") + "\n"
}
