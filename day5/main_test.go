// Only doing to be able to use debugger

package main

import (
	"aoc-2022/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	utils.Solve[string, string](Sln{})
}

func TestStripNonAlpha(t *testing.T) {
	s := "[N] [C]    "
	extracted := extractLetters(s)
	utils.Equal(t, []string{"N", "C", ""}, extracted)
}
