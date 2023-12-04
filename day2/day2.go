package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	results, err := parseInput("day2/input-day2", parseLine)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	var possibleResults []game
	for _, result := range results {
		resultPossible := true
		for _, d := range result.draws {
			if d.red > maxRed || d.green > maxGreen || d.blue > maxBlue {
				resultPossible = false
			}
		}
		if resultPossible {
			possibleResults = append(possibleResults, result)
		}
	}

	possibleSum := 0
	for _, result := range possibleResults {
		possibleSum += result.id
	}
	fmt.Println("part1:", possibleSum) // should be 2278

	powerSum := 0
	for _, gameResult := range results {
		minRed, minGreen, minBlue := 0, 0, 0
		for _, d := range gameResult.draws {
			if d.red != 0 && d.red > minRed {
				minRed = d.red
			}
			if d.green != 0 && d.green > minGreen {
				minGreen = d.green
			}
			if d.blue != 0 && d.blue > minBlue {
				minBlue = d.blue
			}
		}
		fmt.Println(gameResult.id, minRed, minGreen, minBlue)
		powerSum += minRed * minGreen * minBlue
	}
	fmt.Println("part2:", powerSum)
}

type game struct {
	id    int
	draws []draw
}

type draw struct {
	red   int
	green int
	blue  int
}

var colours = []string{"red", "green", "blue"}

func parseInput[t any](fileName string, parseLine func(string) (t, error)) ([]t, error) {
	var result []t

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
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

func parseLine(line string) (game, error) {
	pos := 0
	if !strings.HasPrefix(line, "Game ") {
		return game{}, fmt.Errorf("expecting 'Game ' at start of line, but got: %s", line)
	}
	pos += len("Game ")

	colonPos := strings.IndexRune(line, ':')
	if colonPos == -1 {
		return game{}, fmt.Errorf("expected ':', but not found, line %s", line)
	}

	numString := line[pos:colonPos]
	id, err := strconv.Atoi(numString)
	if err != nil {
		return game{}, fmt.Errorf("parse game id, err: &%s, line: %s", err, line)
	}
	pos = colonPos + 1

	var draws []draw
	for _, drawString := range strings.Split(line[pos:], ";") {
		gemStrings := strings.Split(drawString, ",")

		redCount := 0
		greenCount := 0
		blueCount := 0
		for _, gemString := range gemStrings {
			parsedColour, count, err := parseGem(gemString)
			if err != nil {
				return game{}, fmt.Errorf("%s, line %s", err, line)
			}
			if parsedColour == "red" {
				redCount = count
			}
			if parsedColour == "green" {
				greenCount = count
			}
			if parsedColour == "blue" {
				blueCount = count
			}
		}
		d := draw{
			red:   redCount,
			green: greenCount,
			blue:  blueCount,
		}
		draws = append(draws, d)
	}

	return game{
		id:    id,
		draws: draws,
	}, nil
}

func parseGem(gemString string) (string, int, error) {
	pos := 0
	count := 0
	for gemString[pos] == ' ' {
		pos++
	}
	colonPos := strings.Index(gemString[pos:], " ")
	if colonPos == -1 {
		return "", 0, fmt.Errorf("no space found in gem string: '%s'", gemString)
	}
	colonPos += pos

	count, err := strconv.Atoi(gemString[pos:colonPos])
	if err != nil {
		return "", 0, fmt.Errorf("parse count: %s, %s", err, gemString)
	}

	for _, colour := range colours {
		if strings.HasSuffix(gemString[colonPos:], colour) {
			return colour, count, nil
		}
	}
	return "", 0, fmt.Errorf("not a valid gem: %s", gemString)
}
