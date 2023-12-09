package main

import (
	"fmt"
	"github.com/samber/lo"
	"io/ioutil"
	"log"
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

	sum := 0
	for _, line := range lines {
		ns := strings.Split(line, " ")
		numbers := make([]int, len(ns))
		for i, n := range ns {
			numbers[i], _ = strconv.Atoi(n)
		}
		val := step(numbers)
		sum += val
	}

	fmt.Println("part1: ", sum)
}

func step(numbers []int) int {
	diffs := lo.Map(numbers[1:], func(n int, i int) int {
		return n - numbers[i]
	})
	if isSame(diffs) {
		return numbers[len(numbers)-1] + diffs[len(diffs)-1]
	}

	return numbers[len(numbers)-1] + step(diffs)
}

func isSame(numbers []int) bool {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1] {
			return false
		}
	}
	return true
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	sum := 0
	for _, line := range lines {
		ns := strings.Split(line, " ")
		numbers := make([]int, len(ns))
		for i, n := range ns {
			numbers[i], _ = strconv.Atoi(n)
		}
		val := step2(numbers)
		sum += val
	}

	fmt.Println("part2: ", sum)
}

func step2(numbers []int) int {
	diffs := lo.Map(numbers[1:], func(n int, i int) int {
		return n - numbers[i]
	})
	if isSame(diffs) {
		return numbers[0] - diffs[0]
	}

	return numbers[0] - step2(diffs)
}
