package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"modernc.org/mathutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func hasSignNeighbour(colIdx int, rowIdx int, runeGrid [][]rune) bool {
	for i := rowIdx - 1; i <= rowIdx+1; i++ {
		if i < 0 || i >= len(runeGrid) {
			continue
		}
		for j := colIdx - 1; j <= colIdx+1; j++ {
			if j < 0 || j >= len(runeGrid[i]) {
				continue
			}
			if i == rowIdx && j == colIdx {
				continue
			}
			if runeGrid[i][j] != '.' && (runeGrid[i][j] < '0' || runeGrid[i][j] > '9') {
				//fmt.Printf("neibour: %d %d %c\n", i, j, runeGrid[i][j])
				return true
			}
		}
	}
	return false
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	runeGrid := make([][]rune, len(lines))
	for i, line := range lines {
		runeGrid[i] = []rune(line)
	}

	var re = regexp.MustCompile(`(\d+)`)
	//var reSigns = regexp.MustCompile(`[^\.\d]+`)

	sum := 0
	for rowIdx, line := range lines {
		//fmt.Printf("line: %s\n", line)
		numbers := re.FindAllString(line, -1)
		indices := re.FindAllStringIndex(line, -1)
		for i, number := range numbers {
			//fmt.Printf("match: %s %d\n", number, indices[i])
			numberStartIdx, numberEndIdx := indices[i][0], indices[i][1]
			for colIdx := numberStartIdx; colIdx < numberEndIdx; colIdx++ {
				//fmt.Printf("index: %d\n", colIdx)
				if hasSignNeighbour(colIdx, rowIdx, runeGrid) {
					//fmt.Printf("number: %s hasSign \n", number)
					value, _ := strconv.Atoi(number)
					sum += value
					break
				}
			}
		}
	}
	fmt.Printf("part1 sum: %d\n", sum)
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	runeGrid := make([][]rune, len(lines))
	for i, line := range lines {
		runeGrid[i] = []rune(line)
	}

	var re = regexp.MustCompile(`(\d+)`)
	var reSigns = regexp.MustCompile(`[^\.\d]+`)

	sum := 0
	for rowIdx, line := range lines {
		//fmt.Printf("line: %s\n", line)
		signIndices := reSigns.FindAllStringIndex(line, -1)
		for _, signIndex := range signIndices {
			//fmt.Printf("sign match: %s %d\n", sign, signIndices[i])
			signIdx := signIndex[0]
			startRowIdx, endRowIdx := mathutil.Clamp(rowIdx-1, 0, len(lines)-1), mathutil.Clamp(rowIdx+1, 0, len(lines)-1)
			neighbours := make([]int, 0)
			for rowToCheck := startRowIdx; rowToCheck <= endRowIdx; rowToCheck++ {
				if len(neighbours) == 2 {
					break
				}
				numbers := re.FindAllString(lines[rowToCheck], -1)
				numberIndices := re.FindAllStringIndex(lines[rowToCheck], -1)
				for j, number := range numbers {
					if len(neighbours) == 2 {
						break
					}
					//fmt.Printf("check: %s %d signIdx: %d\n", number, numberIndices[j], signIdx)
					if signIdx >= numberIndices[j][0]-1 && signIdx <= numberIndices[j][1] {
						//fmt.Printf("number match: %s %d\n", number, numberIndices[j])
						value, _ := strconv.Atoi(number)
						neighbours = append(neighbours, value)
						continue // to next number
					}
				}
			}
			if len(neighbours) == 2 {
				//fmt.Printf("neighbours: %d %d\n", neighbours[0], neighbours[1])
				ratio := neighbours[0] * neighbours[1]
				sum += ratio
			}
		}
	}
	fmt.Printf("part2 sum: %d\n", sum)
}
