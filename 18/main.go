package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
	//part1old()
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

	points := make([]Pos, 0)
	curPos := Pos{0, 0}
	points = append(points, curPos)
	boundary := 0
	for _, line := range lines {
		le := strings.Fields(line)
		d := le[0]
		l, _ := strconv.Atoi(le[1])
		boundary += l
		offset := offsets[d]
		curPos = Pos{curPos.x + offset.x*l, curPos.y + offset.y*l}
		points = append(points, curPos)
	}
	points = append(points, points[0])

	// shoelace formula: https://en.wikipedia.org/wiki/Shoelace_formula
	s1, s2 := 0, 0
	for i := 0; i < len(points)-1; i++ {
		s1 += points[i].x * points[i+1].y
		s2 += points[i].y * points[i+1].x
	}
	area := abs(s1-s2) / 2

	// pick theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
	inner := area - boundary/2 + 1

	fmt.Println("part1: ", inner+boundary)
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

	points := make([]Pos, 0)
	curPos := Pos{0, 0}
	points = append(points, curPos)
	boundary := 0
	for _, line := range lines {
		le := strings.Fields(line)[2]
		hashIdx := strings.Index(le, "#")
		hexNumberString := le[hashIdx+1 : hashIdx+6+1]
		lengthHexString := hexNumberString[:len(hexNumberString)-1]
		dirNum, _ := strconv.Atoi(hexNumberString[len(hexNumberString)-1:])
		d := numToDir[dirNum]
		l, _ := strconv.ParseInt(lengthHexString, 16, 64)
		boundary += int(l)
		offset := offsets[d]
		curPos = Pos{curPos.x + offset.x*int(l), curPos.y + offset.y*int(l)}
		points = append(points, curPos)
	}

	points = append(points, points[0])

	// shoelace formula: https://en.wikipedia.org/wiki/Shoelace_formula
	s1, s2 := 0, 0
	for i := 0; i < len(points)-1; i++ {
		s1 += points[i].x * points[i+1].y
		s2 += points[i].y * points[i+1].x
	}
	area := abs(s1-s2) / 2

	// pick theorem: https://en.wikipedia.org/wiki/Pick%27s_theorem
	inner := area - boundary/2 + 1

	fmt.Println("part2: ", inner+boundary)
}

// part 1 first version with floodfill, way to slow for part 2

func part1old() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	curPos := Pos{0, 0}
	minPos := Pos{0, 0}
	maxPos := Pos{0, 0}
	for _, line := range lines {
		fmt.Println(line)
		le := strings.Fields(line)
		d := le[0]
		l, _ := strconv.Atoi(le[1])
		offset := offsets[d]
		curPos = Pos{curPos.x + offset.x*l, curPos.y + offset.y*l}
		minPos.x = mathutil.Min(minPos.x, curPos.x)
		minPos.y = mathutil.Min(minPos.y, curPos.y)
		maxPos.x = mathutil.Max(maxPos.x, curPos.x)
		maxPos.y = mathutil.Max(maxPos.y, curPos.y)
	}
	println(minPos.x, minPos.y, maxPos.x, maxPos.y)

	grid = make([][]rune, maxPos.y-minPos.y+1)
	for i := range grid {
		grid[i] = make([]rune, maxPos.x-minPos.x+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	curPos = Pos{0, 0}
	for _, line := range lines {
		fmt.Println(line)
		le := strings.Fields(line)
		d := le[0]
		l, _ := strconv.Atoi(le[1])
		offset := offsets[d]
		for i := 0; i < l; i++ {
			curPos = Pos{curPos.x + offset.x, curPos.y + offset.y}
			grid[curPos.y-minPos.y][curPos.x-minPos.x] = '#'
			if curPos.y-minPos.y == 0 {
				println("--------------------------->", curPos.x-minPos.x, curPos.y-minPos.y)
			}
		}
	}

	// floodfill
	//fill(grid, Pos{len(grid[0]) / 2, 1})
	fill(grid, Pos{267, 1})

	sum := 0
	for i := range grid {
		println(string(grid[i]))
		for j := range grid[i] {
			if grid[i][j] == '#' {
				sum++
			}
		}
	}
	fmt.Println("part1: ", sum)
}

func fill(g [][]rune, pos Pos) {
	if pos.x < 0 || pos.y < 0 || pos.x >= len(g[0]) || pos.y >= len(g) || g[pos.y][pos.x] == '#' {
		return
	}

	g[pos.y][pos.x] = '#'

	n := []Dir{Up, Down, Left, Right}
	for _, d := range n {
		offset := offsets[d]
		newPos := Pos{pos.x + offset.x, pos.y + offset.y}
		fill(g, newPos)
	}
}
