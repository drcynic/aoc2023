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

type Pos struct {
	x, y int
}

type Dir = int

const (
	Up Dir = iota
	Down
	Left
	Right
)

var grid = make([][]rune, 0)
var energized = make(map[Pos]bool)

type CacheEntry struct {
	pos Pos
	dir Dir
}

var cache = make(map[CacheEntry]bool)

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	pos := Pos{0, 0}
	trace(pos, Right)
	//for _, line := range grid {
	//	println(string(line))
	//}
	fmt.Println("part1: ", len(energized))
}

func trace(pos Pos, dir Dir) {
	if pos.x < 0 || pos.x >= len(grid[0]) || pos.y < 0 || pos.y >= len(grid) || cache[CacheEntry{pos, dir}] {
		return
	}

	cache[CacheEntry{pos, dir}] = true
	energized[pos] = true

	ch := grid[pos.y][pos.x]
	switch ch {
	case '.':
		if dir == Left {
			trace(Pos{pos.x - 1, pos.y}, Left)
		} else if dir == Right {
			trace(Pos{pos.x + 1, pos.y}, Right)
		} else if dir == Up {
			trace(Pos{pos.x, pos.y - 1}, Up)
		} else if dir == Down {
			trace(Pos{pos.x, pos.y + 1}, Down)
		}
		break
	case '|':
		if dir == Left || dir == Right || dir == Up {
			trace(Pos{pos.x, pos.y - 1}, Up)
		}
		if dir == Left || dir == Right || dir == Down {
			trace(Pos{pos.x, pos.y + 1}, Down)
		}
		break
	case '-':
		if dir == Up || dir == Down || dir == Left {
			trace(Pos{pos.x - 1, pos.y}, Left)
		}
		if dir == Up || dir == Down || dir == Right {
			trace(Pos{pos.x + 1, pos.y}, Right)
		}
		break
	case '\\':
		if dir == Left {
			trace(Pos{pos.x, pos.y - 1}, Up)
		} else if dir == Right {
			trace(Pos{pos.x, pos.y + 1}, Down)
		} else if dir == Up {
			trace(Pos{pos.x - 1, pos.y}, Left)
		} else if dir == Down {
			trace(Pos{pos.x + 1, pos.y}, Right)
		}
		break
	case '/':
		if dir == Left {
			trace(Pos{pos.x, pos.y + 1}, Down)
		} else if dir == Right {
			trace(Pos{pos.x, pos.y - 1}, Up)
		} else if dir == Up {
			trace(Pos{pos.x + 1, pos.y}, Right)
		} else if dir == Down {
			trace(Pos{pos.x - 1, pos.y}, Left)
		}
		break
	}
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
		grid = append(grid, []rune(line))
	}

	max := 0
	for i := 0; i < len(grid[0]); i++ {
		cache = make(map[CacheEntry]bool)
		energized = make(map[Pos]bool)
		pos := Pos{i, 0}
		trace(pos, Down)
		val := len(energized)
		max = mathutil.Max(val, max)

		cache = make(map[CacheEntry]bool)
		energized = make(map[Pos]bool)
		pos = Pos{i, len(grid) - 1}
		trace(pos, Up)
		val = len(energized)
		max = mathutil.Max(val, max)
	}
	for i := 0; i < len(grid); i++ {
		cache = make(map[CacheEntry]bool)
		energized = make(map[Pos]bool)
		pos := Pos{0, i}
		trace(pos, Right)
		val := len(energized)
		max = mathutil.Max(val, max)

		cache = make(map[CacheEntry]bool)
		energized = make(map[Pos]bool)
		pos = Pos{len(grid[0]) - 1, i}
		trace(pos, Left)
		val = len(energized)
		max = mathutil.Max(val, max)
	}
	fmt.Println("part2: ", max)
}
