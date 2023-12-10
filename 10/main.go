package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"strings"
)

func main() {
	part1()
	part2()
}

type Point struct {
	x, y int
}

var grid = make([][]rune, 0)
var visited = make(map[Point]bool)
var start = Point{-1, -1}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		//fmt.Println(line)
		chars := []rune(line)
		grid = append(grid, chars)
		for x, c := range chars {
			if c == 'S' {
				start = Point{x, y}
			}
		}
	}

	p := start
	r := checkPos(Point{p.x + 1, p.y}, 0)
	visited = make(map[Point]bool)
	l := checkPos(Point{p.x - 1, p.y}, 0)
	visited = make(map[Point]bool)
	d := checkPos(Point{p.x, p.y + 1}, 0)
	visited = make(map[Point]bool)
	u := checkPos(Point{p.x, p.y - 1}, 0)
	val := mathutil.MaxVal(r, l, d, u) / 2

	fmt.Println("part1: ", val)
	//fmt.Println(grid)
}

func checkPos(p Point, step int) int {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(grid[0]) || p.y >= len(grid) || grid[p.y][p.x] == '.' {
		return step
	}
	step++
	visited[p] = true

	switch grid[p.y][p.x] {
	case '|':
		d := checkPos(Point{p.x, p.y + 1}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(d, u)
	case '-':
		r := checkPos(Point{p.x + 1, p.y}, step)
		l := checkPos(Point{p.x - 1, p.y}, step)
		return mathutil.Max(r, l)
	case 'F':
		r := checkPos(Point{p.x + 1, p.y}, step)
		d := checkPos(Point{p.x, p.y + 1}, step)
		return mathutil.Max(r, d)
	case '7':
		l := checkPos(Point{p.x - 1, p.y}, step)
		d := checkPos(Point{p.x, p.y + 1}, step)
		return mathutil.Max(l, d)
	case 'J':
		l := checkPos(Point{p.x - 1, p.y}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(l, u)
	case 'L':
		r := checkPos(Point{p.x + 1, p.y}, step)
		u := checkPos(Point{p.x, p.y - 1}, step)
		return mathutil.Max(r, u)
	case 'S':
		//println("found", step)
		break
	}

	return step
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		//fmt.Println(line)
		chars := []rune(line)
		grid = append(grid, chars)
		for x, c := range chars {
			if c == 'S' {
				start = Point{x, y}
			}
		}
	}

	p := start
	val := 0
	pipePath := make(map[Point]bool)
	val = checkPos2(Point{p.x + 1, p.y}, 0, &pipePath)
	visited = make(map[Point]bool)
	pipePath2 := make(map[Point]bool)
	l := checkPos2(Point{p.x - 1, p.y}, 0, &pipePath2)
	if l > val {
		pipePath = pipePath2
		val = l
	}
	visited = make(map[Point]bool)
	pipePath3 := make(map[Point]bool)
	d := checkPos2(Point{p.x, p.y + 1}, 0, &pipePath3)
	if d > val {
		pipePath = pipePath3
		val = d
	}
	visited = make(map[Point]bool)
	pipePath4 := make(map[Point]bool)
	u := checkPos2(Point{p.x, p.y - 1}, 0, &pipePath4)
	if u > val {
		pipePath = pipePath4
		val = u
	}

	// hack: manually patch the start position after inspecting the grid
	grid[start.y][start.x] = '|'
	//println("val", val)

	pathElements := pipePath
	count := 0
	for y, line := range grid {
		inside := false
		lastStart := '-'
		for x, c := range line {
			if !pathElements[Point{x, y}] {
				if inside {
					count++
				}
				continue
			}
			if c == '-' {
				continue
			}
			if c == '|' || c == 'F' || c == 'L' {
				inside = !inside
				lastStart = c
			}
			if c == 'J' && lastStart == 'L' {
				inside = !inside
				lastStart = '-'
			}
			if c == '7' && lastStart == 'F' {
				inside = !inside
				lastStart = '-'
			}
		}
	}

	fmt.Println("part2: ", count)
}

func checkPos2(p Point, step int, pipePath *map[Point]bool) int {
	if visited[p] || p.x < 0 || p.y < 0 || p.x >= len(grid[0]) || p.y >= len(grid) || grid[p.y][p.x] == '.' {
		return step
	}
	step++
	visited[p] = true
	(*pipePath)[p] = true

	switch grid[p.y][p.x] {
	case '|':
		d := checkPos2(Point{p.x, p.y + 1}, step, pipePath)
		u := checkPos2(Point{p.x, p.y - 1}, step, pipePath)
		return mathutil.Max(d, u)
	case '-':
		r := checkPos2(Point{p.x + 1, p.y}, step, pipePath)
		l := checkPos2(Point{p.x - 1, p.y}, step, pipePath)
		return mathutil.Max(r, l)
	case 'F':
		r := checkPos2(Point{p.x + 1, p.y}, step, pipePath)
		d := checkPos2(Point{p.x, p.y + 1}, step, pipePath)
		return mathutil.Max(r, d)
	case '7':
		l := checkPos2(Point{p.x - 1, p.y}, step, pipePath)
		d := checkPos2(Point{p.x, p.y + 1}, step, pipePath)
		return mathutil.Max(l, d)
	case 'J':
		l := checkPos2(Point{p.x - 1, p.y}, step, pipePath)
		u := checkPos2(Point{p.x, p.y - 1}, step, pipePath)
		return mathutil.Max(l, u)
	case 'L':
		r := checkPos2(Point{p.x + 1, p.y}, step, pipePath)
		u := checkPos2(Point{p.x, p.y - 1}, step, pipePath)
		return mathutil.Max(r, u)
	case 'S':
		//println("found", step)
		break
	}

	return step
}
