package main

import (
	"fmt"
	"github.com/soroushj/menge"
	"io/ioutil"
	"log"
	"math"
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
		//fmt.Printf("line: %s\n", line)
		allNumbers := line[strings.Index(line, ":")+1:]
		numbers := strings.Split(allNumbers, "|")
		winnerNumbers := strings.Split(numbers[0], " ")
		winnerSet := menge.NewIntSet()
		for _, number := range winnerNumbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			winnerSet.Add(value)
		}
		//fmt.Printf("winnerSet: %s\n", winnerSet)

		ourSet := menge.NewIntSet()
		for _, number := range strings.Split(numbers[1], " ") {
			value, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			ourSet.Add(value)
		}
		//fmt.Printf("ourSet: %s\n", ourSet)

		result := winnerSet.Intersection(ourSet)
		//fmt.Printf("result: %s\n", result)
		score := int(math.Pow(2, float64(len(result)-1)))
		//fmt.Printf("score: %d\n", score)
		sum += score

		//break
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

	sum := 0
	addCards := make(map[int]int)
	for idx, line := range lines {
		//fmt.Printf("line: %s\n", lines)
		allNumbers := line[strings.Index(line, ":")+1:]
		numbers := strings.Split(allNumbers, "|")
		winnerNumbers := strings.Split(numbers[0], " ")
		winnerSet := menge.NewIntSet()
		for _, number := range winnerNumbers {
			value, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			winnerSet.Add(value)
		}
		//fmt.Printf("winnerSet: %s\n", winnerSet)

		ourSet := menge.NewIntSet()
		for _, number := range strings.Split(numbers[1], " ") {
			value, err := strconv.Atoi(number)
			if err != nil {
				continue
			}
			ourSet.Add(value)
		}
		//fmt.Printf("ourSet: %s\n", ourSet)

		result := winnerSet.Intersection(ourSet)
		//fmt.Printf("result: %s\n", result)
		score := len(result)
		for i := 1; i <= score; i++ {
			addCards[idx+1+i] += addCards[idx+1] + 1
		}
		//fmt.Printf("score: %d\n", score)
		sum += (addCards[idx+1] + 1)
		//fmt.Printf("current sum: %d\n", sum)

		//fmt.Printf("addCards: %v\n", addCards)
	}
	fmt.Printf("part2 sum: %d\n", sum)
}
