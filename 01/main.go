package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	fileContent, err := ioutil.ReadFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	score := 0
	score2 := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		//var s1, e1, s2, e2 int
		//_, _ = fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)

	}

	fmt.Printf("score: %d\n", score)
	fmt.Printf("score2: %d\n", score2)
}
