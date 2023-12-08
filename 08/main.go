package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	part1()
	part2()
}

var entries = make(map[string]Entry)

type Entry struct {
	left  string
	right string
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	re := regexp.MustCompile(`([A-Z]{3})`)
	curEntry := "AAA"

	for _, line := range lines[2:] {
		//fmt.Printf("%s\n", line)
		s := re.FindAllString(line, -1)
		entries[s[0]] = Entry{s[1], s[2]}
	}

	//fmt.Println(entries)
	instr := lines[0]
	instrIdx := 0
	steps := 0
	for curEntry != "ZZZ" {
		ins := instr[instrIdx]
		if ins == 'L' {
			curEntry = entries[curEntry].left
		} else {
			curEntry = entries[curEntry].right
		}
		instrIdx = (instrIdx + 1) % len(instr)
		steps = steps + 1
	}

	fmt.Printf("part1: %d\n", steps)
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	re := regexp.MustCompile(`([A-Z12]{3})`)
	curEntries := make([]string, 0)

	for _, line := range lines[2:] {
		s := re.FindAllString(line, -1)
		entries[s[0]] = Entry{s[1], s[2]}
		if s[0][len(s[0])-1:len(s)] == "A" {
			curEntries = append(curEntries, s[0])
		}
	}

	entrySteps := make([]int, len(curEntries))
	instr := lines[0]

	for entryIdx, curEntry := range curEntries {
		instrIdx := 0
		steps := 0
		for curEntry[2] != 'Z' {
			ins := instr[instrIdx]
			if ins == 'L' {
				curEntry = entries[curEntry].left
			} else {
				curEntry = entries[curEntry].right
			}
			curEntries[entryIdx] = curEntry
			instrIdx = (instrIdx + 1) % len(instr)
			steps = steps + 1
		}

		entrySteps[entryIdx] = steps
	}

	sort.Ints(entrySteps)
	fmt.Println(entrySteps)
	lcm := LCM(entrySteps[0], entrySteps[1], entrySteps[2:]...)
	fmt.Println("Part2: ", lcm)
}

// LCM taken from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
