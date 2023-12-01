package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
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

	sum := 0
	for _, line := range lines {
		//fmt.Printf("line: %s\n", line)
		lineChars := []rune(line)
		front, back := 0, 0
		for _, char := range lineChars {
			if char >= '0' && char <= '9' {
				front = int(char - '0')
				break
			}
		}
		for i := len(lineChars) - 1; i >= 0; i-- {
			char := lineChars[i]
			if char >= '0' && char <= '9' {
				back = int(char - '0')
				break
			}
		}
		checksum := front*10 + back
		sum += checksum
	}
	fmt.Printf("part1 sum: %d\n", sum)
}

func matchToValue(match string) int {
	if len(match) == 1 {
		return int(match[0] - '0')
	}

	switch match {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	//fmt.Printf("lines: %d\n", len(lines))

	var re = regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|[1-9])`)

	sum := 0
	for _, line := range lines {
		//fmt.Printf("line: %s\n", line)
		matches := re.FindAllString(line, -1)
		//for _, match := range matches {
		//	fmt.Printf("match: %s\n", match)
		//}
		firstMatch, lastMatch := matches[0], matches[len(matches)-1]
		front, back := matchToValue(firstMatch), matchToValue(lastMatch)

		checksum := front*10 + back
		sum += checksum
	}
	fmt.Printf("part2 sum: %d\n", sum)
}
