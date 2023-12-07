package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

var cardToValue = map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

func handToValue(hand string) int {
	cardMap := make(map[rune]int)
	for _, card := range hand {
		cardMap[card]++
	}
	numCards := make([]int, 0)
	for _, num := range cardMap {
		numCards = append(numCards, num)
	}
	sort.Slice(numCards, func(i, j int) bool {
		return numCards[i] > numCards[j]
	})
	if len(numCards) == 1 {
		return 7 // 5 of a kind
	}
	if len(numCards) == 2 {
		if numCards[0] == 4 {
			return 6 // 4 of a kind
		}
		if numCards[0] == 3 {
			return 5 // full house
		}
	}
	if len(numCards) == 3 {
		if numCards[0] == 3 {
			return 4 // 3 of a kind
		}
		if numCards[0] == 2 {
			return 3 // 2 pair
		}
	}
	if len(numCards) == 4 {
		return 2 // pair
	}

	return 1
}

type Player struct {
	hand      string
	bid       int
	handValue int
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	players := make([]Player, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])
		handValue := handToValue(hand)
		players = append(players, Player{hand: hand, bid: bid, handValue: handValue})
	}

	sort.Slice(players[:], func(i, j int) bool {
		if players[i].handValue != players[j].handValue {
			return players[i].handValue < players[j].handValue
		}
		for r := 0; r < len(players[i].hand); r++ {
			cv1 := cardToValue[rune(players[i].hand[r])]
			cv2 := cardToValue[rune(players[j].hand[r])]
			if cv1 != cv2 {
				return cv1 < cv2
			}
		}
		return false
	})
	//fmt.Println(players)
	totalWins := 0
	for i, player := range players {
		totalWins += player.bid * (i + 1)
	}
	fmt.Printf("part1: %d\n", totalWins)
}

func handToValue2(hand string) int {
	cardMap := make(map[rune]int)
	for _, card := range hand {
		cardMap[card]++
	}
	numJokers := cardMap['J']

	numCards := make([]int, 0)
	for _, num := range cardMap {
		numCards = append(numCards, num)
	}
	sort.Slice(numCards, func(i, j int) bool {
		return numCards[i] > numCards[j]
	})
	if len(numCards) == 1 {
		return 7 // 5 of a kind
	}
	if len(numCards) == 2 {
		if numCards[0] == 4 {
			if numJokers != 0 {
				return 7 // promoted 5 of a kind
			}
			return 6 // 4 of a kind
		}
		if numCards[0] == 3 {
			if numJokers != 0 {
				return 7 // promoted 5 of a kind
			}
			return 5 // full house
		}
	}
	if len(numCards) == 3 {
		if numCards[0] == 3 {
			if numJokers != 0 {
				return 6 // promoted 4 of a kind
			}
			return 4 // 3 of a kind
		}
		if numCards[0] == 2 {
			if numJokers == 2 {
				return 6 // promoted 4 of a kind
			}
			if numJokers == 1 {
				return 5 // promoted full house
			}
			return 3 // 2 pair
		}
	}
	if len(numCards) == 4 {
		if numJokers != 0 {
			return 4 // promoted 3 of a kind
		}
		return 2 // pair
	}

	if numJokers != 0 {
		return 2 // promoted pair
	}

	return 1
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	// patch card value for joker
	cardToValue['J'] = 1

	players := make([]Player, 0)

	for _, line := range lines {
		split := strings.Split(line, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])
		handValue := handToValue2(hand)
		players = append(players, Player{hand: hand, bid: bid, handValue: handValue})
	}

	sort.Slice(players[:], func(i, j int) bool {
		if players[i].handValue != players[j].handValue {
			return players[i].handValue < players[j].handValue
		}
		for r := 0; r < len(players[i].hand); r++ {
			cv1 := cardToValue[rune(players[i].hand[r])]
			cv2 := cardToValue[rune(players[j].hand[r])]
			if cv1 != cv2 {
				return cv1 < cv2
			}
		}
		return false
	})
	//fmt.Println(players)
	totalWins := 0
	for i, player := range players {
		totalWins += player.bid * (i + 1)
	}
	fmt.Printf("part2: %d\n", totalWins)
}
