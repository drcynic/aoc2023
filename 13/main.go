package main

import (
	"fmt"
	"github.com/samber/lo"
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"strings"
)

func main() {
	part1()
	part2()
}

func vertSyms(lines []string, ref int) int {
	symmetries := make(map[int]bool, 0)
	for li, line := range lines {
		for i := 1; i < len(line); i++ {
			s1 := string(lo.Reverse([]rune(line[0:i])))
			s2 := line[i:]
			shorter := mathutil.Min(len(s1), len(s2))
			if s1[0:shorter] == s2[0:shorter] {
				if li == 0 && i != ref {
					symmetries[i] = true
				}
			} else {
				delete(symmetries, i)
			}
		}
	}

	for k, _ := range symmetries {
		return k
	}
	return 0
}

func horzSyms(lines []string, ref int) int {
	// transpose grid and call vertSyms
	cols := len(lines[0])
	rows := len(lines)
	// create transposed grid  9x3 -> 3x9
	grid := make([][]rune, cols)
	for i := 0; i < cols; i++ {
		grid[i] = make([]rune, rows)
	}
	// fill transposed grid
	for c, line := range lines {
		for r, ch := range line {
			grid[r][c] = ch
		}
	}
	newLines := make([]string, len(grid))
	for i, row := range grid {
		newLines[i] = string(row)
	}
	return vertSyms(newLines, ref)
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	patternGroups := strings.Split(text, "\n\n")

	sum := 0
	for _, patternGroup := range patternGroups {
		lines := strings.Split(patternGroup, "\n")
		v := vertSyms(lines, -1)
		if v != 0 {
			sum += v
		} else {
			h := horzSyms(lines, -1)
			sum += h * 100
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
	patternGroups := strings.Split(text, "\n\n")

	sum := 0
	for _, patternGroup := range patternGroups {
		lines := strings.Split(patternGroup, "\n")
		vRef := vertSyms(lines, -1)
		hRef := horzSyms(lines, -1)

		rows, cols := len(lines), len(lines[0])
		grid := make([][]rune, rows)
		for y, line := range lines {
			grid[y] = []rune(line)
		}

		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				ch := grid[y][x]
				if ch == '#' {
					grid[y][x] = '.'
				} else {
					grid[y][x] = '#'
				}

				for i, row := range grid {
					lines[i] = string(row)
				}
				v := vertSyms(lines, vRef)
				h := horzSyms(lines, hRef)
				if v != 0 && v != vRef {
					sum += v
					goto NEXT
				} else if h != 0 && h != hRef {
					sum += h * 100
					goto NEXT
				}
				grid[y][x] = ch // set back char
			}
		}
	NEXT:
	}
	fmt.Println("part2: ", sum)
}
