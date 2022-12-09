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
	return 21
}
func (Sln) SolvePartOne(lines []string) int {
	heights := convertLinesToInt(lines)

	globalMax := createGlobalMaxGrid(heights)

	return getVisibleCount(heights, globalMax)
}

func (Sln) ExampleSolutionPartTwo() int {
	return 8
}

func (Sln) SolvePartTwo(lines []string) int {
	heights := convertLinesToInt(lines)

	distances := getAllViewDistances(heights)

	maxPerRow := utils.MapFn(distances, func(r []int) int {
		m, _ := utils.FindMax(r, func(i int) int { return i })
		return m
	})
	max, _ := utils.FindMax(maxPerRow, func(i int) int { return i })

	return max
}

func main() {

	utils.Solve[int, int](Sln{})

}

/*
create 4 copies of grid
each copy contains the max height of the tree found in the specific direction

30373
25512
65332
33549
35390

for up would be
30373
35573
65573
65579
65599

once each has been found, the original grid can be traversed and visibility can be confirmed by
checking the following:
- the tree is on an edge
- check the specific direction grid at the tree's location. if all of the maxes are smaller than the tree, it is visible
*/

type treeGrid struct {
	heights [][]int
}

func convertLinesToInt(lines []string) [][]int {
	return utils.MapFn(lines, func(l string) []int {
		is := []int{}
		for _, c := range l {
			is = append(is, utils.ParseIntMust(string(c)))
		}
		return is
	})
}

func convertIntsToString(is [][]int) string {
	maxRowCharCounts := utils.MapFn(is, func(r []int) int {
		v, _ := utils.FindMax(r, func(i int) int { return len(fmt.Sprintf("%d", i)) })
		return v
	})

	totalMax, _ := utils.FindMax(maxRowCharCounts, func(i int) int { return i })
	padStr := strings.Repeat(" ", totalMax+1)
	rows := utils.MapFn(is, func(r []int) string {
		a := utils.MapFn(r, func(i int) string { return fmt.Sprintf("%s%d", padStr, i) })
		return strings.Join(a, "")
	})
	return strings.Join(rows, "\n")
}

// better safe than sorry
func copyGrid(heights [][]int) [][]int {
	c := make([][]int, 0, len(heights))
	for _, r := range heights {
		cr := make([]int, len(r))
		copy(cr, r)
		c = append(c, cr)
	}
	return c
}

func handleMaxHorizontalRow(heightRow []int, isRight bool) {
	l := len(heightRow)
	if !isRight {
		for i, v := range heightRow {
			if i == 0 {
				continue
			}
			if heightRow[i-1] < v {
				continue
			}
			heightRow[i] = heightRow[i-1]
		}
		return
	}

	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			continue
		}
		if heightRow[i+1] < heightRow[i] {
			continue
		}
		heightRow[i] = heightRow[i+1]

	}
}

func handleMaxGridHorizontal(heights [][]int, isRight bool) [][]int {
	for _, r := range heights {
		handleMaxHorizontalRow(r, isRight)
	}
	return heights
}

func handleMaxVerticalCol(heights [][]int, col int, isDown bool) {
	l := len(heights)
	if !isDown {
		for i := 0; i < l; i++ {
			if i == 0 {
				continue
			}
			if heights[i-1][col] < heights[i][col] {
				continue
			}
			heights[i][col] = heights[i-1][col]
		}
		return
	}

	for i := l - 1; i >= 0; i-- {
		if i == l-1 {
			continue
		}
		if heights[i+1][col] < heights[i][col] {
			continue
		}
		heights[i][col] = heights[i+1][col]
	}
}

func handleMaxVerticalGrid(heights [][]int, isDown bool) [][]int {
	for i := range heights[0] {
		handleMaxVerticalCol(heights, i, isDown)
	}
	return heights
}

func createMaxGrid(heights [][]int, isRightOrDown bool, isVertical bool) [][]int {
	maxGrid := copyGrid(heights)
	if !isVertical {
		return handleMaxGridHorizontal(maxGrid, isRightOrDown)
	}
	return handleMaxVerticalGrid(maxGrid, isRightOrDown)
}
func createGlobalMaxGrid(heights [][]int) [][]int {
	maxLeft, maxRight, maxUp, maxDown := createMaxGrid(heights, false, false), createMaxGrid(heights, true, false), createMaxGrid(heights, false, true), createMaxGrid(heights, true, true)

	globalMaxGrid := copyGrid(heights)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights); j++ {
			if i == 0 || i == len(heights)-1 || j == 0 || j == len(heights)-1 {
				globalMaxGrid[i][j] = -1
				continue
			}
			min, _ := utils.FindMinEl([]int{maxLeft[i][j-1], maxRight[i][j+1], maxUp[i-1][j], maxDown[i+1][j]})
			globalMaxGrid[i][j] = min
		}
	}
	return globalMaxGrid
}

func getVisibleCount(heights [][]int, globalMaxes [][]int) int {
	c := 0
	for i := range heights {
		for j := range heights {
			if heights[i][j] <= globalMaxes[i][j] {
				continue
			}
			c += 1
		}
	}
	return c
}

// part 2:
/*
like above examine every direction
*/

func getDirectionDistance(heights [][]int, i int, j int, fn func(int, int) (int, int)) int {
	d := 1
	x, y := i, j
	for {
		x, y = fn(x, y)
		if x == 0 || y == 0 || x == len(heights)-1 || y == len(heights)-1 {
			return d
		}
		if heights[x][y] >= heights[i][j] {
			return d
		}
		d++
	}
}

func getViewDistance(heights [][]int, i int, j int) int {
	if i == 0 || j == 0 || i == len(heights)-1 || j == len(heights)-1 {
		return 0
	}

	maxLeft := getDirectionDistance(heights, i, j, func(x, y int) (int, int) { return x - 1, y })
	maxRight := getDirectionDistance(heights, i, j, func(x, y int) (int, int) { return x + 1, y })
	maxUp := getDirectionDistance(heights, i, j, func(x, y int) (int, int) { return x, y - 1 })
	maxDown := getDirectionDistance(heights, i, j, func(x, y int) (int, int) { return x, y + 1 })

	return maxLeft * maxRight * maxUp * maxDown

}

func getAllViewDistances(heights [][]int) [][]int {
	vs := copyGrid(heights)
	for i := range heights {
		for j := range heights {
			vs[i][j] = getViewDistance(heights, i, j)
		}
	}
	return vs
}
