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
	file, err := os.Open("day9/input-day9")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	report := parseInput(file)

	sum := 0
	for _, line := range report {
		diffCascade := createDiffCascade(line)

		predictedIncrease := 0
		for i := len(diffCascade) - 2; i >= 0; i-- {
			add := diffCascade[i][len(diffCascade[i])-1]
			predictedIncrease += add
		}
		prediction := line[len(line)-1] + predictedIncrease
		sum += prediction
	}

	sum2 := 0
	for _, line := range report {
		diffCascade := createDiffCascade(line)

		predictedDecrease := 0
		for i := len(diffCascade) - 2; i >= 0; i-- {
			sub := diffCascade[i][0]
			predictedDecrease = sub - predictedDecrease
		}
		prediction := line[0] - predictedDecrease
		sum2 += prediction
	}
	fmt.Println("part1", sum)  // should be 1953784198
	fmt.Println("part2", sum2) // should be 957
}

func createDiffCascade(readings []int) [][]int {
	var result [][]int

	for {
		allZero := true
		var diffs []int
		for i := 1; i < len(readings); i++ {
			diff := readings[i] - readings[i-1]
			diffs = append(diffs, diff)
			if allZero && diff != 0 {
				allZero = false
			}
		}
		result = append(result, diffs)

		readings = diffs
		if allZero {
			break
		}
	}

	return result
}

func parseInput(input io.Reader) [][]int {
	var result [][]int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var line []int
		text := scanner.Text()
		fields := strings.Fields(text)
		for _, field := range fields {
			value, _ := strconv.Atoi(field)
			line = append(line, value)
		}
		result = append(result, line)
	}

	return result
}
