package main

import (
	"fmt"
	"strings"
	"testing"
)

var input = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	raceData, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatalf(err.Error())
	}

	var resultsPerRace [][]int
	for _, race := range raceData {
		resultsPerRace = append(resultsPerRace, race.possibleResults(1, race.time-1))
	}

	result := 0
	for i, possibleDists := range resultsPerRace {
		raceSum := 0
		for _, dist := range possibleDists {
			if dist > raceData[i].distance {
				raceSum += 1
			}
		}
		if result == 0 {
			result = raceSum
		} else {
			if raceSum != 0 {
				result *= raceSum
			}
		}
		fmt.Println("race", i, "better dists", raceSum)
	}
	fmt.Println("result", result)
}

func TestPart2(t *testing.T) {
	raceResult := parseResult2(strings.NewReader(input))
	results := raceResult.possibleResults(1, raceResult.time-1)
	count := 0
	for _, result := range results {
		if result > raceResult.distance {
			fmt.Println(result)
			count++
		}
	}
	fmt.Println("count", count) // should be 71503
}
