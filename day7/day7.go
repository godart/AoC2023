package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day7/input-day7")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hands, bids := parseInput(file, parseHand1)

	ranking := calculateRanking(hands)
	part1 := 0
	for i := 0; i < len(ranking); i++ {
		index := ranking[i].id
		part1 += bids[index] * (i + 1)
	}

	fmt.Println("part1:", part1)

	file, err = os.Open("day7/input-day7")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hands, bids = parseInput(file, parseHand2)

	ranking = calculateRanking(hands)
	part2 := 0
	for i := 0; i < len(ranking); i++ {
		index := ranking[i].id
		part2 += bids[index] * (i + 1)
	}
	fmt.Println("part2", part2) // should be 249781879
}

func parseInput(reader io.Reader, parseHand func(string) []int) ([][]int, []int) {
	var hands [][]int
	var bids []int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		splitSpace := strings.Split(line, " ")
		if len(splitSpace) >= 2 {
			hands = append(hands, parseHand(splitSpace[0]))
			bids = append(bids, parseBid(splitSpace[1]))
		}
	}

	return hands, bids
}

func parseHand1(s string) []int {
	var result []int

	for _, char := range s {
		value := 0
		if char >= '2' && char <= '9' {
			value = int(char - '0')
		} else {
			switch char {
			case 'T':
				value = 10
			case 'J':
				value = 11
			case 'Q':
				value = 12
			case 'K':
				value = 13
			case 'A':
				value = 14
			}
		}
		result = append(result, value)
	}

	return result
}

func parseHand2(s string) []int {
	parsed := parseHand1(s)
	var result []int
	for _, card := range parsed {
		if card == 11 { // old value of J
			card = 1
		}
		result = append(result, card)
	}
	return result
}

func parseBid(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func calcScore(hand []int) int {
	cloned := slices.Clone(hand)
	slices.Sort(cloned)
	slices.Reverse(cloned)

	maxRun := 0
	secondMax := 0

	jokerCount := 0
	//	var counts map[int]int
	for i := 0; i < len(cloned); i++ {
		num := cloned[i]
		if num == 1 { // needed to count jokers in part2
			jokerCount++
			continue
		}
		runLength := 1
		for j := i + 1; j < len(cloned); j++ {
			next := cloned[j]
			if num == next {
				runLength++
				i++ // card is already counted
			} else {
				break
			}
		}
		if runLength >= maxRun {
			secondMax = maxRun
			maxRun = runLength
		} else if runLength > secondMax {
			secondMax = runLength
		}
	}

	maxRun += jokerCount
	if maxRun == 6 { // all jokers in part2
		return 7
	}
	if maxRun == 5 {
		return 7 // five
	}
	if maxRun == 4 {
		return 6 // four
	}
	if maxRun == 3 {
		if secondMax > 1 {
			return 5 // full house
		}
		return 4 // three
	}
	if maxRun == 2 {
		if secondMax > 1 {
			return 3 // two pairs
		}
		return 2 // one pair
	}

	return 1 // high card
}

type rank struct {
	id    int
	hand  []int
	score int
}

func compareScore(a rank, b rank) int {
	byScore := a.score - b.score
	if byScore != 0 {
		return byScore
	}
	for i, card := range a.hand {
		byCard := card - b.hand[i]
		if byCard != 0 {
			return byCard
		}
	}
	return 0
}

func calculateRanking(hands [][]int) []rank {
	var result []rank
	for i, hand := range hands {
		result = append(result, rank{id: i, hand: hand, score: calcScore(hand)})
	}
	slices.SortFunc(result, compareScore)

	return result
}
