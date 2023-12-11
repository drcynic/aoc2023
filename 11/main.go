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

type Point struct {
	x, y int
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	var galaxies = make([]Point, 0)
	var grid = make([][]rune, 0)
	for _, line := range lines {
		//fmt.Println("l: ", line)
		chars := []rune(line)
		grid = append(grid, chars)
		if !strings.Contains(line, "#") {
			grid = append(grid, chars)
		}
		for x, c := range chars {
			if c == '#' {
				galaxies = append(galaxies, Point{x, len(grid) - 1})
			}
		}
	}

	add := 0
	for x := 0; x < len(grid[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for i, g := range galaxies {
				if g.x-add > x {
					g.x++
					galaxies[i] = g
				}
			}
			add++
		}
	}

	sum := 0
	for s := 0; s < len(galaxies); s++ {
		for e := s + 1; e < len(galaxies); e++ {
			start := galaxies[s]
			end := galaxies[e]
			dx := Abs(start.x - end.x)
			dy := Abs(start.y - end.y)
			dist := dx + dy
			sum += dist
		}
	}

	fmt.Println("part1: ", sum)
}

func Abs(x int) int {
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
	replaceDist := 1000000

	var galaxies = make([]Point, 0)
	var grid = make([][]rune, 0)
	add := 0
	for _, line := range lines {
		chars := []rune(line)
		grid = append(grid, chars)
		if !strings.Contains(line, "#") {
			add += replaceDist - 1
		}
		for x, c := range chars {
			if c == '#' {
				galaxies = append(galaxies, Point{x, len(grid) - 1 + add})
			}
		}
	}

	add = 0
	for x := 0; x < len(grid[0]); x++ {
		hasGalaxy := false
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for i, g := range galaxies {
				if g.x-add > x {
					g.x += replaceDist - 1
					galaxies[i] = g
				}
			}
			add += replaceDist - 1
		}
	}

	sum := 0
	for s := 0; s < len(galaxies); s++ {
		for e := s + 1; e < len(galaxies); e++ {
			start := galaxies[s]
			end := galaxies[e]
			dx := Abs(start.x - end.x)
			dy := Abs(start.y - end.y)
			dist := dx + dy
			sum += dist
		}
	}

	fmt.Println("part2: ", sum)
}
