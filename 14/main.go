package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)

	lines := strings.Split(text, "\n")

	rows, cols := len(lines), len(lines[0])
	grid := make([][]rune, rows)
	for y, line := range lines {
		grid[y] = []rune(line)
	}

	sum := 0
	moved := true
	for moved {
		moved = false
		for y := 1; y < rows; y++ {
			for x := 0; x < cols; x++ {
				ch := grid[y][x]
				if ch == 'O' && grid[y-1][x] == '.' {
					grid[y-1][x] = 'O'
					grid[y][x] = '.'
					moved = true
				}
			}
		}
	}
	// count north
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 'O' {
				sum += rows - y
			}
		}
	}

	fmt.Println("part1: ", sum)
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)

	lines := strings.Split(text, "\n")

	rows, cols := len(lines), len(lines[0])
	grid := make([][]rune, rows)
	for y, line := range lines {
		grid[y] = []rune(line)
		//println(line)
	}

	const numCycles = 1000000000
	doNCycles(numCycles, rows, cols, grid)

	// count north
	sum := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 'O' {
				sum += rows - y
			}
		}
	}

	fmt.Println("part2: ", sum)
}

func doNCycles(numCycles int, rows int, cols int, grid [][]rune) {
	cache := make(map[uint64]int) // hash -> cycle
	cycleDetected := false
	for cycle := 0; cycle < numCycles; cycle++ {
		// move north
		moved := true
		for moved {
			moved = false
			for y := 1; y < rows; y++ {
				for x := 0; x < cols; x++ {
					ch := grid[y][x]
					if ch == 'O' && grid[y-1][x] == '.' {
						grid[y-1][x] = 'O'
						grid[y][x] = '.'
						moved = true
					}
				}
			}
		}
		// move west
		moved = true
		for moved {
			moved = false
			for x := 1; x < cols; x++ {
				for y := 0; y < rows; y++ {
					ch := grid[y][x]
					if ch == 'O' && grid[y][x-1] == '.' {
						grid[y][x-1] = 'O'
						grid[y][x] = '.'
						moved = true
					}
				}
			}
		}
		// move south
		moved = true
		for moved {
			moved = false
			for y := 0; y < rows-1; y++ {
				for x := 0; x < cols; x++ {
					ch := grid[y][x]
					if ch == 'O' && grid[y+1][x] == '.' {
						grid[y+1][x] = 'O'
						grid[y][x] = '.'
						moved = true
					}
				}
			}
		}

		// move east
		moved = true
		for moved {
			moved = false
			for x := 0; x < cols-1; x++ {
				for y := 0; y < rows; y++ {
					ch := grid[y][x]
					if ch == 'O' && grid[y][x+1] == '.' {
						grid[y][x+1] = 'O'
						grid[y][x] = '.'
						moved = true
					}
				}
			}
		}

		if !cycleDetected {
			cycleLines := make([]string, rows)
			for y := 0; y < rows; y++ {
				cycleLines[y] = string(grid[y])
			}
			cycleHash := hash(strings.Join(cycleLines, ""))
			if cache[cycleHash] > 0 {
				cycleDetected = true
				cycleLength := cycle - cache[cycleHash]
				remainder := (numCycles - cycle) % cycleLength
				fmt.Println("cycle detected", cycle, cache[cycleHash], cycleLength, remainder)
				cycle = numCycles - remainder
				println("set cycle to", cycle)
			}
			cache[cycleHash] = cycle
		}
	}
	//println("no cycle detected")
	//return 0
}

func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}
