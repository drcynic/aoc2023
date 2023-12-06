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

func calcNumWins(time int, targetDist int) int {
	//fmt.Printf("time: %d, dist: %d\n", time, targetDist)
	wins := 0
	for i := 1; i <= time; i++ {
		dist := i * (time - i)
		if dist > targetDist {
			wins += 1
		}
	}
	return wins
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	var re = regexp.MustCompile(`(\d+)`)
	timeStrings := re.FindAllString(lines[0], -1)
	times := make([]int, len(timeStrings))
	for i, timeString := range timeStrings {
		times[i], _ = strconv.Atoi(timeString)
	}
	distStrings := re.FindAllString(lines[1], -1)
	dists := make([]int, len(distStrings))
	for i, distString := range distStrings {
		dists[i], _ = strconv.Atoi(distString)
	}

	wins := 1
	for i, time := range times {
		w := calcNumWins(time, dists[i])
		wins *= w
	}

	fmt.Printf("part1: %d\n", wins)
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	timeString := strings.Join(strings.Fields(strings.Split(lines[0], ":")[1]), "")
	distString := strings.Join(strings.Fields(strings.Split(lines[1], ":")[1]), "")
	time, _ := strconv.Atoi(timeString)
	dist, _ := strconv.Atoi(distString)

	wins := calcNumWins(time, dist)
	fmt.Printf("part2: %d\n", wins)
}
