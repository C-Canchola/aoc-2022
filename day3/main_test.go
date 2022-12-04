// Only doing to be able to use debugger

package main

import (
	"aoc-2022/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	utils.Solve[int, int](Sln{})
}

func TestPriority(t *testing.T) {
	utils.Equal(t, 27, getPriority('A'))
	utils.Equal(t, 52, getPriority('Z'))
	utils.Equal(t, 1, getPriority('a'))
	utils.Equal(t, 26, getPriority('z'))
}
