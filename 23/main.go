package main

import (
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"strings"
)

func main() {
	part1()
	part2()
}

type Pos struct {
	x, y int
}

type Dir = string

const (
	Up    Dir = "U"
	Down      = "D"
	Left      = "L"
	Right     = "R"
)

var numToDir = map[int]Dir{0: Right, 1: Down, 2: Left, 3: Up}
var offsets = map[Dir]Pos{Up: Pos{0, -1}, Down: Pos{0, 1}, Left: Pos{-1, 0}, Right: Pos{1, 0}}
var slopeToDir = map[rune]Dir{'<': Left, '>': Right, '^': Up, 'v': Down}
var grid = make([][]rune, 0)
var endPos = Pos{-1, -1}

var cache = make(map[Pos]int)
var globalMaxFound = 0

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	grid = make([][]rune, 0)

	for _, line := range lines {
		//println(line)
		grid = append(grid, []rune(line))
	}
	startPos := Pos{strings.IndexByte(lines[0], '.'), 0}
	endPos = Pos{strings.IndexByte(lines[len(lines)-1], '.'), len(lines) - 1}
	println(startPos.x, startPos.y)
	println(endPos.x, endPos.y)

	step := -1
	path := make(map[Pos]bool)
	dfs(startPos, Pos{startPos.x, startPos.y}, step, &path)

	for k, _ := range globalPath {
		grid[k.y][k.x] = 'O'
	}

	for _, line := range grid {
		println(string(line))
	}

	println("part1:", globalMaxFound)
}

var globalPath = make(map[Pos]bool)

func dfs(pos, nextPos Pos, step int, path *map[Pos]bool) int {
	if !canWalk(pos, nextPos) {
		return step
	}
	step++

	if nextPos.x == 11 && nextPos.y == 4 {
		println("11,4")
	}

	if _, ok := (*path)[nextPos]; ok {
		return step // already visited
	}

	(*path)[nextPos] = true
	if nextPos == endPos {
		if step > globalMaxFound {
			globalMaxFound = step
			for k, v := range *path {
				globalPath[k] = v
			}
		}
		//fmt.Printf("found path: %v\n", path)
		println("found", step)
		delete(*path, nextPos)
		return step
	}

	c, ok := cache[nextPos]
	if ok && c >= step {
		return step
	}
	cache[nextPos] = step

	result := step
	for i := 0; i < 4; i++ {
		offset := offsets[numToDir[i]]
		nextNextPos := Pos{nextPos.x + offset.x, nextPos.y + offset.y}
		result = mathutil.Max(dfs(nextPos, nextNextPos, step, path), result)
	}
	delete(*path, nextPos)

	return result
}

func isSlope(r rune) bool {
	return r == '<' || r == '>' || r == '^' || r == 'v'
}

func canWalk(pos, nextPos Pos) bool {
	validField := nextPos.x >= 0 && nextPos.x < len(grid[0]) && nextPos.y >= 0 && nextPos.y < len(grid) && grid[nextPos.y][nextPos.x] != '#'
	if !validField {
		return false
	}
	curChar := grid[pos.y][pos.x]
	if isSlope(curChar) {
		validOffsetForSlope := offsets[slopeToDir[curChar]]
		if validOffsetForSlope.x != nextPos.x-pos.x || validOffsetForSlope.y != nextPos.y-pos.y {
			return false // next pos is not in slope direction
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	grid = make([][]rune, 0)

	for _, line := range lines {
		//println(line)
		grid = append(grid, []rune(line))
	}
	startPos := Pos{strings.IndexByte(lines[0], '.'), 0}
	endPos = Pos{strings.IndexByte(lines[len(lines)-1], '.'), len(lines) - 1}
	println(startPos.x, startPos.y)
	println(endPos.x, endPos.y)

	step := -1
	path := make(map[Pos]bool)
	dfs2(Pos{startPos.x, startPos.y}, step, &path)

	for k, _ := range globalPath {
		grid[k.y][k.x] = 'O'
	}

	for _, line := range grid {
		println(string(line))
	}

	println("part2:", globalMaxFound)
}

func dfs2(pos Pos, step int, path *map[Pos]bool) int {
	if !canWalk2(pos) {
		return step
	}
	step++

	if _, ok := (*path)[pos]; ok {
		return step // already visited
	}

	if pos == endPos {
		if step > globalMaxFound {
			globalMaxFound = step
			globalPath = make(map[Pos]bool)
			for k, v := range *path {
				globalPath[k] = v
			}
		}
		// simply print out the current max found, after a few seconds we have the global max found but still iterating
		println("found", step, "globalMaxFound", globalMaxFound)
		return step
	}
	(*path)[pos] = true

	// this cache doesn't work
	//c, ok := cache[pos]
	//if ok && c < step {
	//	return c
	//}
	//cache[pos] = step

	result := step
	for i := 0; i < 4; i++ {
		offset := offsets[numToDir[i]]
		nextPos := Pos{pos.x + offset.x, pos.y + offset.y}
		result = mathutil.Max(dfs2(nextPos, step, path), result)
	}
	delete(*path, pos)

	return result
}

func canWalk2(nextPos Pos) bool {
	validField := nextPos.x >= 0 && nextPos.x < len(grid[0]) && nextPos.y >= 0 && nextPos.y < len(grid) && grid[nextPos.y][nextPos.x] != '#'
	if !validField {
		return false
	}
	return true
}
