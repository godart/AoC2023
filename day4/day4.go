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
	file, err := os.Open("day4/input-day4")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cards, err := parseInput(file, parseCard)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sum1 := 0
	for _, card := range cards {
		matches := card.matches()
		if len(matches) != 0 {
			points := len(matches) - 1
			sum1 += 1 << points
		}
	}
	fmt.Println("part1:", sum1) // should be 27454

	copies := countCopies(cards)
	sum2 := 0
	for i, numCopies := range copies {
		if i == 0 { // copies starts at 0, card numbers start at 1
			continue
		}
		fmt.Println(numCopies + 1)
		sum2 += numCopies + 1 // add one for the original card

	}
	fmt.Println("part2", sum2) // should be 6857330
}

type scratchCard struct {
	id      int
	winning []int
	given   []int
}

func (card scratchCard) matches() []int {
	result := []int{}
	for _, number := range card.winning {
		if slices.Contains(card.given, number) {
			result = append(result, number)
		}
	}
	return result
}

func parseInput[t any](input io.Reader, parseLine func(string) (t, error)) ([]t, error) {
	var result []t

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parsed, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		result = append(result, parsed)
	}
	return result, nil
}

func parseCard(line string) (scratchCard, error) {
	split1 := strings.Split(line, ":")
	if len(split1) != 2 {
		return scratchCard{}, fmt.Errorf("wrong amount of ':', %s", line)
	}
	cardNoPart := split1[0]
	split2 := strings.Split(cardNoPart, "Card")
	if len(split2) != 2 {
		return scratchCard{}, fmt.Errorf("no card number found, %s", cardNoPart)
	}
	noString := split2[1]
	number, err := strconv.Atoi(strings.TrimSpace(noString))
	if err != nil {
		return scratchCard{}, fmt.Errorf("no card number found, %s, %s", cardNoPart, err)
	}

	numbersPart := split1[1]
	split3 := strings.Split(numbersPart, "|")
	if len(split3) != 2 {
		return scratchCard{}, fmt.Errorf("invalid numbersPart, %s", numbersPart)
	}

	winningString := split3[0]
	givenString := split3[1]

	var winning []int
	for _, noString := range strings.Split(winningString, " ") {
		if noString != "" {
			no, err := strconv.Atoi(noString)
			if err != nil {
				return scratchCard{}, fmt.Errorf("winningString: %s: %s", winningString, err)
			}
			winning = append(winning, no)
		}
	}

	var given []int
	for _, noString := range strings.Split(givenString, " ") {
		if noString != "" {
			no, err := strconv.Atoi(noString)
			if err != nil {
				return scratchCard{}, fmt.Errorf("givenString: %s: %s", givenString, err)
			}
			given = append(given, no)
		}
	}
	return scratchCard{id: number, winning: winning, given: given}, nil
}

func countCopies(cards []scratchCard) []int {
	copies := make([]int, len(cards)+1)
	for _, card := range cards {
		cardCopies := copies[card.id]
		matches := card.matches()
		for i := 1; i <= len(matches); i++ {
			cardAmount := 1 + cardCopies
			copyIndex := card.id + i
			if copyIndex < len(copies) {
				copies[copyIndex] += cardAmount
			}
		}
	}
	return copies
}
