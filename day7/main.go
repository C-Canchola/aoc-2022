package main

import (
	"aoc-2022/utils"
	"strings"
)

type Sln struct {
}

var _ utils.Solution[int, int] = Sln{}

func (Sln) ExampleSolutionPartOne() int {
	return 95437
}
func (Sln) SolvePartOne(lines []string) int {
	root := getDirFromLines(&lines)
	allSizes := []int{}
	getSizeWithRecordedSubSizes(root, &allSizes)
	filteredSizes := utils.Filter(allSizes, func(i int) bool { return i <= 100000 })
	return utils.Sum(filteredSizes)

}

func (Sln) ExampleSolutionPartTwo() int {
	return 24933642
}

func (Sln) SolvePartTwo(lines []string) int {
	root := getDirFromLines(&lines)
	allSizes := []int{}
	getSizeWithRecordedSubSizes(root, &allSizes)

	utils.SortFromPredicate(allSizes, func(a, b int) bool { return a < b })

	usedSize := allSizes[len(allSizes)-1]
	unusedSize := 70000000 - usedSize
	sizeToDelete := 30000000 - unusedSize

	for _, s := range allSizes {
		if s < sizeToDelete {
			continue
		}
		return s
	}
	return -1
}

func main() {

	utils.Solve[int, int](Sln{})

}

type file struct {
	name string
	size int
}

type dir struct {
	name  string
	files []file
	dirs  []dir
}

func getDirFromLines(lines *[]string) dir {

	name := strings.Split((*lines)[0], " ")[2]

	*lines = (*lines)[2:]

	files := []file{}
	for {
		if len(*lines) == 0 {
			break
		}
		l := (*lines)[0]
		if strings.HasPrefix(l, "$") {
			break
		}

		splt := strings.Split(l, " ")
		if splt[0] != "dir" {
			files = append(files, file{size: utils.ParseIntMust(splt[0]), name: splt[1]})
		}
		*lines = (*lines)[1:]
	}

	dirs := []dir{}
	for {
		if len(*lines) == 0 {
			break
		}
		if (*lines)[0] == "$ cd .." {
			*lines = (*lines)[1:]
			break
		}
		newDir := getDirFromLines(lines)
		dirs = append(dirs, newDir)
	}

	return dir{name: name, files: files, dirs: dirs}
}

func getSizeWithRecordedSubSizes(d dir, sizes *[]int) int {
	fileSizesSum := utils.Sum(utils.MapFn(d.files, func(f file) int { return f.size }))
	dSum := 0
	for _, subD := range d.dirs {
		dSum += getSizeWithRecordedSubSizes(subD, sizes)
	}
	totalSum := dSum + fileSizesSum
	*sizes = append(*sizes, totalSum)
	return totalSum
}
