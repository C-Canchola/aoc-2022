package main

import (
	"aoc-2022/utils"
	"image"
	"image/color"
	"image/gif"
	"os"
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

func makeImage(positions []pos, height int, width int) *image.Paletted {
	palette := make([]color.Color, 256)

	for i := 0; i < 256; i++ {
		palette[i] = color.Gray{255 - uint8(i)}
	}

	rect := image.Rect(0, 0, width, height)
	img := image.NewPaletted(rect, palette)
	for _, p := range positions {
		img.SetColorIndex(p.x, p.y, 255)
	}
	return img
}

func getExGif() {
	lines := utils.ReadExample()
	moves := linesToDirections(lines)
	frames := findPathFramesForGif(moves)

	dims := getFrameDims(frames)
	_ = dims

	imgs := []*image.Paletted{}
	for _, f := range frames {
		img := makeImage([]pos{f.head, f.tail}, dims.maxY, dims.maxX)
		imgs = append(imgs, img)
	}
	anim := gif.GIF{Delay: make([]int, len(imgs)), Image: imgs}
	if err := gif.EncodeAll(os.Stdout, &anim); err != nil {
		panic(err)
	}
}

func main() {
	getExGif()
	// utils.Solve[int, int](Sln{})

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

func makeMoveForGif(headPos pos, tailPos pos, move direction, tailPositions []pos, headPositions []pos) (pos, pos, []pos, []pos) {
	for i := 0; i < move.times; i++ {
		headPos = pos{x: headPos.x + move.x, y: headPos.y + move.y}
		tailPos = getUpdatedTailPosition(tailPos, headPos)

		tailPositions = append(tailPositions, tailPos)
		headPositions = append(headPositions, headPos)
	}
	return headPos, tailPos, tailPositions, headPositions
}

func findTailPath(moves []direction) []pos {
	head, tail := pos{}, pos{}
	path := []pos{}
	for _, m := range moves {
		head, tail, path = makeMove(head, tail, m, path)
	}
	return path
}

type frame struct {
	head pos
	tail pos
}

func findPathFramesForGif(moves []direction) []frame {
	head, tail := pos{}, pos{}
	headPath, tailPath := []pos{}, []pos{}
	for _, m := range moves {
		head, tail, tailPath, headPath = makeMoveForGif(head, tail, m, tailPath, headPath)
	}

	frames := []frame{frame{head: pos{}, tail: pos{}}}
	for i, h := range headPath {
		f1 := frame{head: h, tail: frames[len(frames)-1].tail}
		f2 := frame{head: h, tail: tailPath[i]}
		frames = append(frames, f1, f2)
	}

	return frames
}

type frameDimensions struct {
	minX, minY, maxY, maxX int
}

func getFrameDims(frames []frame) frameDimensions {
	minX := frames[0].head.x
	if minX > frames[0].tail.x {
		minX = frames[0].tail.x
	}
	minY := frames[0].head.y
	if minY > frames[0].tail.y {
		minY = frames[0].tail.y
	}

	maxX := frames[0].head.x
	if frames[0].tail.x > maxX {
		maxX = frames[0].tail.x
	}

	maxY := frames[0].head.y
	if frames[0].tail.y > maxY {
		maxY = frames[0].tail.y
	}
	updateMinOrMax := func(ptr *int, v int, forMax bool) {
		if forMax {
			if *ptr >= v {
				return
			}
			*ptr = v
			return
		}
		if *ptr <= v {
			return
		}
		*ptr = v
	}
	for _, f := range frames {
		updateMinOrMax(&maxY, f.head.y, true)
		updateMinOrMax(&maxY, f.tail.y, true)
		updateMinOrMax(&maxX, f.head.x, true)
		updateMinOrMax(&maxX, f.tail.x, true)

		updateMinOrMax(&minX, f.head.x, false)
		updateMinOrMax(&minX, f.tail.x, false)
		updateMinOrMax(&minY, f.head.y, false)
		updateMinOrMax(&minY, f.tail.y, false)
	}

	return frameDimensions{minX: minX, maxY: maxY, maxX: maxX, minY: minY}
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
