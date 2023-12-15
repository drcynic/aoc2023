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

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	line := lines[0]
	tokens := strings.Split(line, ",")

	sum := 0
	for _, token := range tokens {
		currentVal := 0
		for _, ch := range token {
			currentVal = (int(ch) + currentVal) * 17 % 256
		}
		sum += currentVal
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
	line := lines[0]
	tokens := strings.Split(line, ",")
	re := regexp.MustCompile(`([a-z]+)([=-])(\d*)`)
	boxes := make([][]string, 256)

	for _, token := range tokens {
		matches := re.FindStringSubmatch(token)
		label := matches[1]
		op := matches[2]
		lensId := matches[3]
		boxIdx := 0
		for _, ch := range label {
			boxIdx = (int(ch) + boxIdx) * 17 % 256
		}

		if op == "=" {
			storedLabel := label + " " + lensId
			i := findIndex(boxes[boxIdx], label)
			if i == -1 {
				boxes[boxIdx] = append(boxes[boxIdx], storedLabel)
			} else {
				boxes[boxIdx][i] = storedLabel
			}
		} else {
			box := boxes[boxIdx]
			idxToRemove := findIndex(box, label)
			if idxToRemove != -1 {
				boxes[boxIdx] = append(box[:idxToRemove], box[idxToRemove+1:]...)
			}
		}
	}

	sum := 0
	for boxIdx, box := range boxes {
		for slotIdx, entry := range box {
			flString := strings.Fields(entry)[1]
			fl, _ := strconv.Atoi(flString)
			val := (boxIdx + 1) * (slotIdx + 1) * fl
			sum += val
		}
	}
	fmt.Println("part2: ", sum)
}

func findIndex(box []string, label string) int {
	for i, entry := range box {
		if strings.HasPrefix(entry, label) {
			return i
		}
	}
	return -1
}
