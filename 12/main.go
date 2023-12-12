package main

import (
	"fmt"
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
		//fmt.Println("l: ", line)
		split := strings.Fields(line)
		springs := split[0]
		gns := strings.Split(split[1], ",")
		gn := make([]int, len(gns))
		for i, gs := range gns {
			gn[i], _ = strconv.Atoi(gs)
		}
		val := dfs(springs, gn)
		sum += val
	}

	fmt.Println("part1: ", sum)
}

func dfs(springs string, gn []int) int {
	if len(springs) == 0 {
		if len(gn) == 0 {
			return 1
		}
		return 0
	}

	if len(gn) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}
		return 1
	}

	r := 0

	if springs[0] == '.' || springs[0] == '?' {
		r += dfs(springs[1:], gn)
	}

	if springs[0] == '#' || springs[0] == '?' {
		if gn[0] <= len(springs) && !strings.Contains(springs[:gn[0]], ".") && (gn[0] == len(springs) || springs[gn[0]] != '#') {
			if gn[0] == len(springs) {
				r += dfs("", gn[1:])
			} else {
				r += dfs(springs[gn[0]+1:], gn[1:])
			}
		}
	}

	return r
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
		//fmt.Println("l", i, ": ", line)
		split := strings.Fields(line)
		springs := split[0]
		gnsgen := split[1]
		for i := 0; i < 4; i++ {
			springs += "?"
			springs += split[0]
			gnsgen += ","
			gnsgen += split[1]
		}
		gns := strings.Split(gnsgen, ",")
		gn := make([]int, len(gns))
		for i, gs := range gns {
			gn[i], _ = strconv.Atoi(gs)
		}

		var cache [][]int
		for i := 0; i < len(springs); i++ {
			cache = append(cache, make([]int, len(gn)+1))
			for j := 0; j < len(gn)+1; j++ {
				cache[i][j] = -1
			}
		}

		val := dfs2(0, 0, springs, gn, cache)
		sum += val
	}

	fmt.Println("part2: ", sum)
}

func dfs2(i, j int, springs string, gn []int, cache [][]int) int {
	if i >= len(springs) {
		if j < len(gn) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	r := 0
	if springs[i] == '.' {
		r = dfs2(i+1, j, springs, gn, cache)
	} else {
		if springs[i] == '?' {
			r += dfs2(i+1, j, springs, gn, cache)
		}
		if j < len(gn) {
			count := 0
			for k := i; k < len(springs); k++ {
				if count > gn[j] || springs[k] == '.' || (count == gn[j] && springs[k] == '?') {
					break
				}
				count += 1
			}

			if count == gn[j] {
				if i+count < len(springs) && springs[i+count] != '#' {
					r += dfs2(i+count+1, j+1, springs, gn, cache)
				} else {
					r += dfs2(i+count, j+1, springs, gn, cache)
				}
			}
		}
	}

	cache[i][j] = r
	return r
}
