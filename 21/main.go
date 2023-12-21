package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
var grid = make([][]rune, 0)

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	grid = make([][]rune, 0)

	startPos := Pos{-1, -1}
	for _, line := range lines {
		//println(line)
		grid = append(grid, []rune(line))
		if i := strings.IndexByte(line, 'S'); i != -1 {
			startPos = Pos{i, len(grid) - 1}
			grid[startPos.y][startPos.x] = '.'
		}
	}

	curState := map[Pos]bool{startPos: true}
	nextState := make(map[Pos]bool)
	for step := 0; step < 64; step++ {
		for k, _ := range curState {
			for i := 0; i < 4; i++ {
				offset := offsets[numToDir[i]]
				nextPos := Pos{k.x + offset.x, k.y + offset.y}
				if canWalk(nextPos) {
					nextState[nextPos] = true
				}
			}
		}
		curState = nextState
		nextState = make(map[Pos]bool)
	}
	fmt.Println("part1: ", len(curState))
}

func canWalk(pos Pos) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid) && grid[pos.y][pos.x] == '.'
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

	startPos := Pos{-1, -1}
	for _, line := range lines {
		//println(line)
		grid = append(grid, []rune(line))
		if i := strings.IndexByte(line, 'S'); i != -1 {
			startPos = Pos{i, len(grid) - 1}
			grid[startPos.y][startPos.x] = '.'
		}
	}

	gridSize := len(grid)
	rest := 26501365 % gridSize
	//startPos = Pos{0, len(grid) / 2}
	//const STEPS = 0
	//const STEPS = 6
	//const STEPS = 10
	//const STEPS = 50
	//const STEPS = 64
	//const STEPS = 100
	//const STEPS = 500
	//const STEPS = 1000
	//const STEPS = 5000
	//const STEPS = 65
	//const STEPS = 65 + 131
	const STEPS = 26501365
	// 590104708070703

	//w := len(grid[0])
	//numRight := (STEPS - w/2) / w
	//restRight := (STEPS - w/2) % w
	//println("numRight", numRight)
	//println("restRight", restRight)
	//numCorner := (STEPS - w) / w
	//restCorner := (STEPS - w) % w
	//println("numCorner", numCorner)
	//println("restCorner", restCorner)
	//
	//startQuadPlots := calc(2*w, startPos)
	//leftCenterStartQuadPlots := calc(gridSize+rest, Pos{0, w / 2})
	//leftCenterStartQuadPlots2 := calc(2*gridSize+rest, Pos{0, w / 2})
	//restLeftCenter := calc(restRight, Pos{0, w / 2})
	//leftTopStartQuadPlots1 := calc(2*gridSize-1, Pos{0, 0})
	//leftTopStartQuadPlots2 := calc(2*gridSize, Pos{0, 0})
	//restTopLeft := calc(restCorner, Pos{0, 0})
	//
	//println("startQuadPlots", startQuadPlots)
	//println("leftCenterStartQuadPlots", leftCenterStartQuadPlots)
	//println("leftCenterStartQuadPlots2", leftCenterStartQuadPlots2)
	//println("restLeftCenter", restLeftCenter)
	//println("leftTopStartQuadPlots", leftTopStartQuadPlots1)
	//println("leftTopStartQuadPlots2", leftTopStartQuadPlots2)
	//println("restTopLeft", restTopLeft)
	//
	//sumCenter := leftCenterStartQuadPlots*numRight/2 + leftCenterStartQuadPlots2*numRight/2 + restLeftCenter
	//println("sumCenter", sumCenter)
	//
	//allCorner := numCorner * (numCorner + 1) / 2
	//println("allCorner", allCorner)
	//sumCorner := leftTopStartQuadPlots1 * allCorner / 2
	//sumCorner = sumCorner + leftTopStartQuadPlots2*allCorner/2
	//sumCorner = sumCorner + (numCorner+1)*restTopLeft
	//println("sumCorner", sumCorner)
	//sumQuarter := sumCenter + sumCorner
	//println("sumQuarter", sumQuarter)
	//sum := sumQuarter*4 + startQuadPlots
	//println("sum", sum)

	// the stuff above didn't work, so I did it the quadratic equation way taken from reddit by hand with
	// the following first 3 samples
	samples := [3]int{0, 1, 2}
	for i := 0; i < 3; i++ {
		steps := i*gridSize + rest
		samples[i] = calc(steps, startPos)
	}
	fmt.Printf("samples: %v\n", samples)

	// solved the quadratic equation (ax^2 + bx +c) by hand with the first 3 samples to this
	var a = 14419
	var b = 14590
	var c = samples[0] // 3703

	f := func(x int) int { return a*x*x + b*x + c }

	// double-check the first 3 results
	for i, s := range samples {
		fmt.Println(s, f(i))
	}
	// cals finals plots
	num := STEPS / gridSize
	plots := abs(f(num))

	println("part2:", plots)
}

func calc(steps int, startPos Pos) int {
	curState := map[Pos]bool{startPos: true}
	nextState := make(map[Pos]bool)
	for step := 1; step <= steps; step++ {
		for k, _ := range curState {
			for i := 0; i < 4; i++ {
				offset := offsets[numToDir[i]]
				nextPos := Pos{k.x + offset.x, k.y + offset.y}
				if canWalk2(nextPos) {
					nextState[nextPos] = true
				}
			}
		}
		curState = nextState
		nextState = make(map[Pos]bool)
	}
	return len(curState)
}

func canWalk2(pos Pos) bool {
	x := (pos.x%len(grid[0]) + len(grid[0])) % len(grid[0])
	y := (pos.y%len(grid) + len(grid)) % len(grid)
	return grid[y][x] == '.'
}
