package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalcScore(t *testing.T) {
	tests := []struct {
		name string
		hand string
		exp  int
	}{
		{
			name: "high card",
			hand: "Q2345",
			exp:  1,
		},
		{
			name: "five",
			hand: "44444",
			exp:  7,
		},
		{
			name: "four",
			hand: "23222",
			exp:  6,
		},
		{
			name: "three",
			hand: "2T3TT",
			exp:  4,
		},
		{
			name: "full house",
			hand: "2J2JJ",
			exp:  5,
		},
		{
			name: "two pair",
			hand: "3QT2TQ",
			exp:  3,
		},
		{
			name: "one pair",
			hand: "32T3K",
			exp:  2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calcScore(parseHand1(test.hand))
			if got != test.exp {
				t.Errorf("got: %d, exp: %d", got, test.exp)
			}
		})
	}
}

func TestCalcScore2(t *testing.T) {
	tests := []struct {
		name string
		hand string
		exp  int
	}{
		{
			name: "high card plus joker is one pair",
			hand: "6J84Q",
			exp:  2,
		},
		{
			name: "pair plus joker is three",
			hand: "QKJK3",
			exp:  4,
		},
		{
			name: "two pair plus joker is full house",
			hand: "99J22",
			exp:  5,
		},
		{
			name: "three plus joker is four",
			hand: "TTJT3",
			exp:  6,
		},
		{
			name: "five jokers is five",
			hand: "JJJJJ",
			exp:  7,
		},
		{
			name: "four plus joker is five",
			hand: "JJJ2J",
			exp:  7,
		},
		//K4JJJ
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calcScore(parseHand2(test.hand))
			if got != test.exp {
				t.Errorf("got: %d, exp: %d", got, test.exp)
			}
		})
	}
}

var input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestPart1(t *testing.T) {
	hands, bids := parseInput(strings.NewReader(input), parseHand1)

	ranking := calculateRanking(hands)
	// (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.
	sum1 := 0
	for i := 0; i < len(ranking); i++ {
		index := ranking[i].id
		sum1 += bids[index] * (i + 1)
	}
	if sum1 != 6440 {
		t.Fatalf("wrong sum: %d, exp %d", sum1, 6440)
	}
}

func TestPart2(t *testing.T) {
	hands, bids := parseInput(strings.NewReader(input), parseHand2)

	ranking := calculateRanking(hands)
	fmt.Println(ranking)
	fmt.Println(bids)
	/// 765 * 1 + 28 + 684 * 2 + 684 * 3 + 483 * 4 + 220 * 5

	sum1 := 0
	for i := 0; i < len(ranking); i++ {
		index := ranking[i].id
		sum1 += bids[index] * (i + 1)
	}
	if sum1 != 5905 {
		t.Fatalf("wrong sum: %d, exp %d", sum1, 5905)
	}
}
