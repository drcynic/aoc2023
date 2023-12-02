package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func valid(match string) bool {
	split := strings.Split(match, " ")
	color := split[1]
	value, _ := strconv.Atoi(split[0])

	switch color {
	case "red":
		return value <= maxRed
	case "green":
		return value <= maxGreen
	case "blue":
		return value <= maxBlue
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

	var reValues = regexp.MustCompile(`(\d+ (red|green|blue))`)

	sum := 0
	for i, line := range lines {
		//fmt.Printf("line: %s\n", line)
		matches := reValues.FindAllString(line, -1)
		allValid := true
		for _, match := range matches {
			//fmt.Printf("match: %s\n", match)
			if !valid(match) {
				allValid = false
				break
			}
		}
		if allValid {
			//fmt.Printf("add game: %d\n", i+1)
			sum += i + 1
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

	var reValues = regexp.MustCompile(`(\d+ (red|green|blue))`)

	sum := 0
	for _, line := range lines {
		//fmt.Printf("line: %s\n", line)
		matches := reValues.FindAllString(line, -1)
		var maxValues = make(map[string]int)
		for _, match := range matches {
			//fmt.Printf("match: %s\n", match)
			split := strings.Split(match, " ")
			color := split[1]
			value, _ := strconv.Atoi(split[0])
			if maxValues[color] < value {
				maxValues[color] = value
			}
		}
		power := maxValues["red"] * maxValues["green"] * maxValues["blue"]
		//fmt.Printf("add game %d power %d\n", i+1, power)
		sum += power
	}
	fmt.Printf("part2 sum: %d\n", sum)
}
