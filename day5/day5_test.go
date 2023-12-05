package main

import (
	"github.com/google/go-cmp/cmp"
	"strings"
	"testing"
)

var input = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestParse(t *testing.T) {
	seeds, maps, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parse input: %s", err.Error())
	}
	seedsExp := []int{79, 14, 55, 13}
	if diff := cmp.Diff(seeds, seedsExp); diff != "" {
		t.Fatalf("seeds -got/+exp\n%s", diff)
	}

	if len(maps) != 7 {
		t.Fatalf("Wrong amount of categories: got %d - exp 7, %v", len(maps), maps)
	}
	seedToSoilExp := categoryMap{
		From: "light",
		To:   "temperature",
		Mappings: []mapping{
			{FromStart: 77, ToStart: 45, Length: 23},
			{FromStart: 45, ToStart: 81, Length: 19},
			{FromStart: 64, ToStart: 68, Length: 13},
		},
	}
	if diff := cmp.Diff(maps[4], seedToSoilExp); diff != "" {
		t.Fatalf("%s", diff)
	}
}

func TestResolve(t *testing.T) {
	seeds, maps, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parse input: %s", err)
	}

	mapped := maps[0].resolve(seeds)

	mappedExp := []int{81, 14, 57, 13}
	if diff := cmp.Diff(mapped, mappedExp); diff != "" {
		t.Fatalf(diff)
	}
}

func TestPart1(t *testing.T) {
	seeds, maps, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parse input: %s", err)
	}

	locations := mapToLocations(seeds, maps)
	locationsExp := []int{82, 43, 86, 35}

	if diff := cmp.Diff(locations, locationsExp); diff != "" {
		t.Errorf("-got/+exp:\n%s", diff)
	}
}

func TestPart2(t *testing.T) {
	seeds, maps, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parse input: %s", err)
	}

	location := 0
	for ; location <= 255; location++ {
		if locationMapsSeed(location, maps, seeds) {
			break
		}
	}
	if location != 46 {
		t.Fatalf("minLocation: %d, expected 46", location)
	}
}
