package main

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

var input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestPart1(t *testing.T) {
	cards, err := parseInput(strings.NewReader(input), parseCard)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var matches [][]int
	for _, card := range cards {
		matches = append(matches, card.matches())
	}
	matchesExp := [][]int{
		{48, 83, 86, 17},
		{32, 61},
		{1, 21},
		{84},
		{},
		{},
	}
	if diff := cmp.Diff(matches, matchesExp); diff != "" {
		t.Fatalf(diff)
	}

}

func TestPart2(t *testing.T) {
	cards, err := parseInput(strings.NewReader(input), parseCard)
	if err != nil {
		t.Fatalf(err.Error())
	}

	copies := countCopies(cards)
	copiesExp := []int{0, 0, 1, 3, 7, 13, 0}
	if diff := cmp.Diff(copies, copiesExp); diff != "" {
		t.Error(diff)
	}
}
