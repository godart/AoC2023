package main

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

var input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {

	validParts := parseGrid(strings.NewReader(input)).findParts()
	validPartsExp := []int{467, 35, 633, 617, 592, 755, 664, 598}
	if diff := cmp.Diff(validParts, validPartsExp); diff != "" {
		t.Fatalf("-got/+want\n%s", diff)
	}

	validParts2 := parseGrid(strings.NewReader(input)).findParts2()
	if diff := cmp.Diff(validParts2, validPartsExp); diff != "" {
		t.Fatalf("2 -got/+want\n%s", diff)
	}
}

func TestPart2(t *testing.T) {
	g := parseGrid(strings.NewReader(input))
	parts, starPositions := parsePartsAndStars(strings.NewReader(input))

	gears := g.findGears(parts, starPositions)
	gearsExp := [][2]int{{35, 467}, {598, 755}}
	if diff := cmp.Diff(gears, gearsExp); diff != "" {
		t.Fatalf("-got/+want\n%s", diff)
	}
}
