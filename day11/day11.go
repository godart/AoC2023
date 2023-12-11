package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct {
	row int
	col int
}

func main() {
	file, _ := os.Open("day11/input-day11")
	scanner := bufio.NewScanner(file)

	galaxyLocs, rowEmpty, colEmpty := parseInput(scanner)

	sum := sumDistances(galaxyLocs, rowEmpty, colEmpty, 2) // should be 9521776
	fmt.Println("part1", sum)
	sum2 := sumDistances(galaxyLocs, rowEmpty, colEmpty, 1000000) // should be 553224415344
	fmt.Println("part2", sum2)
}

func parseInput(scanner *bufio.Scanner) ([]pos, []int, []int) {
	var galaxyLocs []pos
	var rowEmpty []int
	noGalaxiesInCol := make(map[int]int)
	for i := 0; scanner.Scan(); i++ {
		galaxyInRow := false
		for j, char := range scanner.Text() {
			if noGalaxiesInCol[j] == 0 { // explicitly fill the map with a 0 value for counting later
				noGalaxiesInCol[j] = 0
			}
			if char == '#' {
				galaxyLocs = append(galaxyLocs, pos{row: i, col: j})
				galaxyInRow = true
				noGalaxiesInCol[j]++
			}
		}
		if !galaxyInRow {
			rowEmpty = append(rowEmpty, i)
		}
	}

	var colEmpty []int
	for i, count := range noGalaxiesInCol {
		if count == 0 {
			colEmpty = append(colEmpty, i)
		}
	}
	return galaxyLocs, rowEmpty, colEmpty
}

func sumDistances(galaxyLocs []pos, rowsEmpty []int, colsEmpty []int, expansionFactor int) int {
	sum := 0

	expandTerm := expansionFactor - 1 // number of additional rows/cols per expansion
	for i := 0; i < len(galaxyLocs); i++ {
		start := galaxyLocs[i]

		for j := i + 1; j < len(galaxyLocs); j++ {
			end := galaxyLocs[j]

			startRow, endRow := start.row, end.row
			if startRow > endRow {
				startRow, endRow = endRow, startRow
			}
			startCol, endCol := start.col, end.col
			if startCol > endCol {
				startCol, endCol = endCol, startCol
			}

			dist := endRow - startRow + endCol - startCol

			expansion := 0
			for _, rowNo := range rowsEmpty {
				if startRow < rowNo && rowNo < endRow {
					expansion += expandTerm
				}
			}
			for _, colNo := range colsEmpty {
				if startCol < colNo && colNo < endCol {
					expansion += expandTerm
				}
			}

			sum += dist + expansion
		}
	}

	return sum
}
