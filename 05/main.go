package main

import (
	"fmt"
	"github.com/go-camp/interval"
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

func transform(block string, seeds menge.IntSet) menge.IntSet {
	lines := strings.Split(block, "\n")
	seeds2 := menge.NewIntSet()
	for _, line := range lines[1:] {
		//fmt.Printf("line: %s\n", line)
		soilStrings := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(soilStrings[0])
		sourceStart, _ := strconv.Atoi(soilStrings[1])
		rng, _ := strconv.Atoi(soilStrings[2])
		//fmt.Printf("destStart: %d, sourceStart: %d, range: %d\n", destStart, sourceStart, rng)
		for _, seed := range seeds.AsSlice() {
			if seed >= sourceStart && seed < sourceStart+rng {
				//fmt.Printf("move seed: %d\n", seed)
				seeds2.Add(destStart + (seed - sourceStart))
				seeds.Remove(seed)
			}
		}
		//fmt.Printf("seeds: %v\n", seeds)
	}
	seeds2 = seeds2.Union(seeds)
	//fmt.Printf("seeds2: %v\n", seeds2)
	return seeds2
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	blocks := strings.Split(text, "\n\n")
	var seeds = menge.NewIntSet()

	for i, block := range blocks {
		//fmt.Printf("block: %s\n", block)
		if i == 0 {
			startIdx := strings.Index(block, " ")
			seedStrings := strings.Split(block[startIdx+1:], " ")
			for _, seedString := range seedStrings {
				seed, _ := strconv.Atoi(seedString)
				seeds.Add(seed)
			}

			//fmt.Printf("seeds: %v\n", seeds)
		} else {
			seeds = transform(block, seeds)
		}
	}
	//fmt.Printf("seeds: %v\n", seeds)
	minLoc := math.MaxInt32
	for _, seed := range seeds.AsSlice() {
		if seed < minLoc {
			minLoc = seed
		}
	}
	fmt.Printf("minLoc: %d\n", minLoc)
}

func transform2(block string, seeds interval.OrderedSet) interval.OrderedSet {
	lines := strings.Split(block, "\n")
	seeds2 := interval.OrderedSet{}
	for _, line := range lines[1:] {
		//fmt.Printf("line: %s\n", line)
		soilStrings := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(soilStrings[0])
		sourceStart, _ := strconv.Atoi(soilStrings[1])
		rng, _ := strconv.Atoi(soilStrings[2])
		//fmt.Printf("destStart: %d, sourceStart: %d, range: %d\n", destStart, sourceStart, rng)
		sourceRange := interval.Interval{Begin: sourceStart, IncBegin: true, End: sourceStart + rng, IncEnd: false}
		//fmt.Printf("sourceRange: %v\n", sourceRange)
		for _, seed := range seeds.Intervals() {
			intersection := seed.Intersect(sourceRange)
			if !intersection.IsEmpty() {
				//fmt.Printf("move seed: %v\n", seed)
				seeds2.Add(interval.Interval{Begin: destStart + (intersection.Begin - sourceStart), IncBegin: true, End: destStart + (intersection.End - sourceStart), IncEnd: false})
				seeds.Remove(intersection)
			}
		}

		//fmt.Printf("seeds: %v\n", seeds)
	}

	for _, seed := range seeds.Intervals() {
		seeds2.Add(seed)
	}
	//fmt.Printf("seeds2: %v\n", seeds2)
	return seeds2
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	blocks := strings.Split(text, "\n\n")

	seeds := interval.OrderedSet{}
	//fmt.Printf("a: %s\n", seeds)

	for i, block := range blocks {
		//fmt.Printf("block: %s\n", block)
		if i == 0 {
			startIdx := strings.Index(block, " ")
			seedStrings := strings.Split(block[startIdx+1:], " ")
			for j := 0; j < len(seedStrings); j += 2 {
				start, _ := strconv.Atoi(seedStrings[j])
				length, _ := strconv.Atoi(seedStrings[j+1])
				seeds.Add(interval.Interval{Begin: start, IncBegin: true, End: start + length, IncEnd: false})
			}

			//fmt.Printf("seeds: %v\n", seeds)
		} else {
			seeds = transform2(block, seeds)
		}
	}
	//fmt.Printf("seeds: %v\n", seeds)
	minLoc := seeds.Intervals()[0].Begin
	fmt.Printf("minLoc: %d\n", minLoc)
}
