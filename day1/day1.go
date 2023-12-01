package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "day1/input-day1"
	calibrationValues1, err := parseCalibrations(fileName, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resultPart1 := 0
	for _, value := range calibrationValues1 {
		resultPart1 += value
	}
	calibrationValues2, err := parseCalibrations(fileName, true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resultPart2 := 0
	for _, value := range calibrationValues2 {
		resultPart2 += value
	}

	fmt.Println("part1:", resultPart1) // should be 54667
	fmt.Println("part2:", resultPart2) // should be 54203
}

func parseCalibrations(fileName string, parseSpelled bool) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var numStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var calibrationValues []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []int
		for pos := 0; pos < len(line); {
			char := line[pos]
			if char >= '0' && char <= '9' {
				numbers = append(numbers, int(char-'0'))
			}
			if parseSpelled {
				restOfLine := line[pos:]
				for i, numString := range numStrings {
					if strings.HasPrefix(restOfLine, numString) {
						numbers = append(numbers, i+1)
					}
				}
			}
			pos++
		}

		first := 0
		last := 0
		if len(numbers) > 0 {
			first, last = numbers[0], numbers[len(numbers)-1]
		}

		value := first*10 + last
		calibrationValues = append(calibrationValues, value)
	}

	return calibrationValues, nil
}
