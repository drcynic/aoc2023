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

type Rating [4]int

var xmasToRatingIdx = map[string]int{"x": 0, "m": 1, "a": 2, "s": 3}

type Rule struct {
	xmas         string
	less         bool
	compareValue int
	target       string
}

type Workflow struct {
	rules         []Rule
	defaultTarget string
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	parts := strings.Split(text, "\n\n")
	workflowsPart := strings.Split(parts[0], "\n")
	ratingsPart := strings.Split(parts[1], "\n")

	workflows := make(map[string]Workflow)
	ratings := make([]Rating, 0)
	for _, line := range workflowsPart {
		a := strings.Split(line, "{")
		name := a[0]
		ruleStrings := strings.Split(a[1][:len(a[1])-1], ",")
		defaultTarget := ruleStrings[len(ruleStrings)-1]
		rules := make([]Rule, 0)
		for _, ruleString := range ruleStrings[:len(ruleStrings)-1] {
			xmas := ruleString[0:1]
			less := ruleString[1:2] == "<"
			ruleParts := strings.Split(ruleString[2:], ":")
			compareValue, _ := strconv.Atoi(ruleParts[0])
			target := ruleParts[1]
			rule := Rule{xmas, less, compareValue, target}
			rules = append(rules, rule)
		}
		workflow := Workflow{rules, defaultTarget}
		workflows[name] = workflow
	}
	//fmt.Printf("wf: %v\n", workflows)

	reRating := regexp.MustCompile(`(\d+)`)
	for _, line := range ratingsPart {
		//println(line)
		ratingGroups := reRating.FindAllString(line, -1)
		x, _ := strconv.Atoi(ratingGroups[0])
		m, _ := strconv.Atoi(ratingGroups[1])
		a, _ := strconv.Atoi(ratingGroups[2])
		s, _ := strconv.Atoi(ratingGroups[3])
		rating := Rating{x, m, a, s}
		ratings = append(ratings, rating)
	}
	//fmt.Printf("%v\n", ratings)

	sum := process(ratings, workflows)
	fmt.Printf("part1: %v\n", sum)
}

func process(ratings []Rating, workflows map[string]Workflow) int {
	aRatings := make([]Rating, 0)
	rRatings := make([]Rating, 0)

	for _, rating := range ratings {
		newTarget := "in"
		for newTarget != "A" && newTarget != "R" {
			curWorkflow := workflows[newTarget]
			for _, rule := range curWorkflow.rules {
				ruleRating := rating[xmasToRatingIdx[rule.xmas]]
				if rule.less {
					if ruleRating < rule.compareValue {
						newTarget = rule.target
						goto done
					}
				} else {
					if ruleRating >= rule.compareValue {
						newTarget = rule.target
						goto done
					}
				}
			}
			newTarget = curWorkflow.defaultTarget
		done:
		}
		if newTarget == "A" {
			aRatings = append(aRatings, rating)
		} else if newTarget == "R" {
			rRatings = append(rRatings, rating)
		} else {
			println("ouch")
		}
	}

	sum := 0
	for _, rating := range aRatings {
		for i := 0; i < len(rating); i++ {
			sum += rating[i]
		}
	}
	return sum
}

type Range struct {
	start, end int
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	parts := strings.Split(text, "\n\n")
	workflowsPart := strings.Split(parts[0], "\n")

	workflows := make(map[string]Workflow)
	for _, line := range workflowsPart {
		a := strings.Split(line, "{")
		name := a[0]
		ruleStrings := strings.Split(a[1][:len(a[1])-1], ",")
		defaultTarget := ruleStrings[len(ruleStrings)-1]
		rules := make([]Rule, 0)
		for _, ruleString := range ruleStrings[:len(ruleStrings)-1] {
			xmas := ruleString[0:1]
			less := ruleString[1:2] == "<"
			ruleParts := strings.Split(ruleString[2:], ":")
			compareValue, _ := strconv.Atoi(ruleParts[0])
			target := ruleParts[1]
			rule := Rule{xmas, less, compareValue, target}
			rules = append(rules, rule)
		}
		workflow := Workflow{rules, defaultTarget}
		workflows[name] = workflow
	}
	//fmt.Printf("wf: %v\n", workflows)

	sum := eval(workflows, "in", []Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}})
	fmt.Printf("part2: %v\n", sum)
}

func eval(workflows map[string]Workflow, workflow string, ratingRanges []Range) int {
	if workflow == "A" {
		result := 1
		for _, r := range ratingRanges {
			result *= r.end - r.start + 1
		}
		return result
	}
	if workflow == "R" {
		return 0
	}

	result := 0
	curWorkflow := workflows[workflow]
	for _, rule := range curWorkflow.rules {
		ratingRanges2 := make([]Range, len(ratingRanges))
		copy(ratingRanges2, ratingRanges[:])
		ratingIdx := xmasToRatingIdx[rule.xmas]

		if rule.less {
			ratingRanges[ratingIdx].start = rule.compareValue
			ratingRanges2[ratingIdx].end = rule.compareValue - 1
		} else {
			ratingRanges[ratingIdx].end = rule.compareValue
			ratingRanges2[ratingIdx].start = rule.compareValue + 1
		}
		result += eval(workflows, rule.target, ratingRanges2)
	}
	result += eval(workflows, curWorkflow.defaultTarget, ratingRanges)
	return result
}
