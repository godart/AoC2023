package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day6/input-day6")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	raceData, err := parseInput(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// UGLY HACKY STUFF STARTS HERE ********
	var resultsPerRace [][]int
	for _, race := range raceData {
		resultsPerRace = append(resultsPerRace, race.possibleResults(1, race.time-1))
	}

	result := 1
	for i, possibleDists := range resultsPerRace {
		raceSum := 0
		for _, dist := range possibleDists {
			if dist > raceData[i].distance {
				raceSum += 1
			}
		}
		if raceSum != 0 {
			result *= raceSum
		}
		fmt.Println("race", i, "better dists", raceSum)
	}
	fmt.Println("part1:", result) // should be 608902

	file, err = os.Open("day6/input-day6")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	raceResult := parseResult2(file)
	fmt.Println("part2:", raceResult)

	// **** EVEN MORE UGLY STUFF HERE
	results := raceResult.possibleResults(1, raceResult.time-1)
	count := 0
	for _, result := range results {
		if result > raceResult.distance {
			count++
		}
	}
	fmt.Println("count", count) // should be 71503

}

type raceRecord struct {
	time     int
	distance int
}

func (r raceRecord) possibleResults(min int, max int) []int {
	var result []int

	for hold := min; hold <= max; hold++ {
		dist := (r.time - hold) * hold
		result = append(result, dist)
	}

	return result
}

func parseInput(input io.Reader) ([]raceRecord, error) {
	contents, _ := io.ReadAll(input)
	lines := strings.Split(string(contents), "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("invalid file: got %d lines, expected 2", len(lines))
	}
	timesString := strings.Split(lines[0], ":")[1]
	distancesString := strings.Split(lines[1], ":")[1]
	times, err := parseNumbers(timesString)
	if err != nil {
		return nil, err
	}
	distances, err := parseNumbers(distancesString)
	if err != nil {
		return nil, err
	}
	if len(times) != len(distances) {
		return nil, fmt.Errorf("number of columns does not match: %d, %d", len(times), len(distances))
	}
	result := []raceRecord{}
	for i, _ := range times {
		result = append(result, raceRecord{
			time:     times[i],
			distance: distances[i],
		})
	}
	return result, nil
}

func parseNumbers(numString string) ([]int, error) {
	result := []int{}
	for _, subString := range strings.Split(numString, " ") {
		trimmed := strings.TrimSpace(subString)
		if trimmed != "" {
			number, err := strconv.Atoi(trimmed)
			if err != nil {
				return nil, fmt.Errorf("invalid number string: %s, %s", numString, err)
			}
			result = append(result, number)
		}
	}
	return result, nil
}

func parseResult2(reader io.Reader) raceRecord {
	time := 0
	distance := 0
	scanner := bufio.NewScanner(reader)

	linesRead := 0
	for scanner.Scan() {
		line := scanner.Text()
		numString := strings.Split(line, ":")[1]
		numString = strings.ReplaceAll(numString, " ", "")
		readValue, err := strconv.Atoi(numString)
		if err != nil {
			fmt.Println(err)
		}
		if linesRead == 0 {
			time = readValue
			linesRead++
		} else {
			distance = readValue
			break
		}
	}

	return raceRecord{
		time:     time,
		distance: distance,
	}
}
