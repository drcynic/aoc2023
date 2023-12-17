package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

var grid = make([][]int, 0)
var turnDirs = map[Dir][]Dir{Up: {Left, Right}, Down: {Left, Right}, Left: {Up, Down}, Right: {Up, Down}}
var offsets = map[Dir]Pos{Up: {0, -1}, Down: {0, 1}, Left: {-1, 0}, Right: {1, 0}}

type CacheEntry struct {
	pos  Pos
	step int
	dir  Dir
}

var cache = make(map[CacheEntry]int)
var globalMinFound = 1000000000

func canWalk(pos Pos) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.x < len(grid[0]) && pos.y < len(grid)
}

type QueueEntry struct {
	pos  Pos
	step int
	dir  Dir
}

var queueCache = make(map[QueueEntry]int) // heatLoss for walk pos/step/dir

func eval(pos Pos, step int, dir Dir, heatLoss int, queue *[]QueueEntry) {
	step++
	heatLoss += grid[pos.y][pos.x]

	if pos.x == len(grid[0])-1 && pos.y == len(grid)-1 {
		//println("found", heatLoss)
		if heatLoss < globalMinFound {
			globalMinFound = heatLoss
		}
		return
	}

	if heatLoss > globalMinFound {
		return
	}

	cacheEntry, ok := cache[CacheEntry{pos, step, dir}]
	if ok && cacheEntry < heatLoss {
		return
	}

	cache[CacheEntry{pos, step, dir}] = heatLoss

	newPositions := make([]Pos, 0)
	if step < 3 {
		newPos := Pos{pos.x + offsets[dir].x, pos.y + offsets[dir].y}
		newPositions = append(newPositions, newPos)
		if canWalk(newPos) {
			newEntry := QueueEntry{newPos, step, dir}
			cv, ok2 := queueCache[newEntry]
			if !ok2 {
				*queue = append(*queue, newEntry)
				queueCache[newEntry] = heatLoss
			} else if cv > heatLoss {
				queueCache[newEntry] = heatLoss
			}
		}
	}
	for _, newDir := range turnDirs[dir] {
		newPos := Pos{pos.x + offsets[newDir].x, pos.y + offsets[newDir].y}
		if canWalk(newPos) {
			newEntry := QueueEntry{newPos, 0, newDir}
			cv, ok2 := queueCache[newEntry]
			if !ok2 {
				*queue = append(*queue, newEntry)
				queueCache[newEntry] = heatLoss
			} else if cv > heatLoss {
				queueCache[newEntry] = heatLoss
			}
		}
	}
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		gridLine := make([]int, 0)
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			gridLine = append(gridLine, v)
		}
		grid = append(grid, gridLine)
	}

	globalMinFound = (len(grid) * len(grid[0]) * 9) / 2

	queue := make([]QueueEntry, 0)
	queue = append(queue, QueueEntry{Pos{0, 1}, 0, Down})
	queue = append(queue, QueueEntry{Pos{1, 0}, 0, Right})

	for len(queue) > 0 {
		entry := queue[0]
		heatLoss := queueCache[entry]
		queue = queue[1:]
		delete(queueCache, entry)
		eval(entry.pos, entry.step, entry.dir, heatLoss, &queue)
	}

	fmt.Println("part1: ", globalMinFound)
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	grid = make([][]int, 0)

	for _, line := range lines {
		gridLine := make([]int, 0)
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			gridLine = append(gridLine, v)
		}
		grid = append(grid, gridLine)
	}

	globalMinFound = (len(grid) * len(grid[0]) * 9) / 2
	cache = make(map[CacheEntry]int)

	queue := make([]QueueEntry, 0)
	queueCache = make(map[QueueEntry]int)
	queue = append(queue, QueueEntry{Pos{0, 1}, 0, Down})
	queue = append(queue, QueueEntry{Pos{1, 0}, 0, Right})

	for len(queue) > 0 {
		entry := queue[0]
		heatLoss := queueCache[entry]
		queue = queue[1:]
		delete(queueCache, entry)
		eval2(entry.pos, entry.step, entry.dir, heatLoss, &queue)
	}

	fmt.Println("part2: ", globalMinFound)
}

func eval2(pos Pos, step int, dir Dir, heatLoss int, queue *[]QueueEntry) {
	step++
	heatLoss += grid[pos.y][pos.x]

	distX := len(grid[0]) - pos.x - 1
	distY := len(grid) - pos.y - 1
	if distX == 0 && step+distY < 4 {
		return
	}
	if distY == 0 && step+distX < 4 {
		return
	}

	if pos.x == len(grid[0])-1 && pos.y == len(grid)-1 {
		//println("found", heatLoss)
		if heatLoss < globalMinFound {
			globalMinFound = heatLoss
		}
		return
	}

	if heatLoss > globalMinFound {
		return
	}

	cacheEntry, ok := cache[CacheEntry{pos, step, dir}]
	if ok && cacheEntry < heatLoss {
		return
	}

	cache[CacheEntry{pos, step, dir}] = heatLoss

	newPositions := make([]Pos, 0)
	if step < 10 {
		newPos := Pos{pos.x + offsets[dir].x, pos.y + offsets[dir].y}
		newPositions = append(newPositions, newPos)
		checkMove(newPos, step, dir, queue, heatLoss)
	}
	if step >= 4 {
		for _, newDir := range turnDirs[dir] {
			newPos := Pos{pos.x + offsets[newDir].x, pos.y + offsets[newDir].y}
			checkMove(newPos, 0, newDir, queue, heatLoss)
		}
	}
}

func checkMove(newPos Pos, step int, dir Dir, queue *[]QueueEntry, heatLoss int) {
	if canWalk(newPos) {
		newEntry := QueueEntry{newPos, step, dir}
		cv, ok2 := queueCache[newEntry]
		if !ok2 {
			*queue = append(*queue, newEntry)
			queueCache[newEntry] = heatLoss
		} else if cv > heatLoss {
			queueCache[newEntry] = heatLoss
		}
	}
}
