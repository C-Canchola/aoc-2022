// Only doing to be able to use debugger

package main

import (
	"aoc-2022/utils"
	"testing"
)

func TestSolution(t *testing.T) {
	utils.Solve[int, int](Sln{})
}

func TestGetPrimeFactors(t *testing.T) {
	helper := func(v int, keys []int, t *testing.T) {
		m := getPrimeFactors(v)
		ks := utils.MapKeys(m)
		utils.Sort(ks, func(i int) int { return i })
		utils.Equal(t, keys, ks)
	}

	helper(2, []int{2}, t)
	helper(144, []int{2, 19}, t)
	helper(19+3, []int{2, 3}, t)
	helper(19+3*19, []int{2, 3}, t)
	helper(19+3*3, []int{2, 3}, t)
	helper(19+3*4, []int{2, 3}, t)

}

func TestDivisibility(t *testing.T) {
	utils.Equal(t, true, isDivisible(19*2*3, 19))
	utils.Equal(t, true, isDivisible(7565, 17))
}
