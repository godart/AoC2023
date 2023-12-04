package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"unicode"
)

func main() {
	file, err := os.Open("day3/input-day3")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	schematic := parseGrid(file)

	sum1wrong := 0
	for _, partNo := range schematic.findParts() {
		sum1wrong += partNo
	}
	fmt.Println("part1 try1:", sum1wrong) // should be 535078 // had a bug here, which got fixed during refactoring for the second try

	sum1 := 0
	for _, partNo := range schematic.findParts2() {
		sum1 += partNo
	}
	fmt.Println("part1:", sum1) // should be 535078

	file, _ = os.Open("day3/input-day3")
	partList, stars := parsePartsAndStars(file)
	gears := schematic.findGears(partList, stars)
	sum2 := 0
	for _, gear := range gears {
		ratio := gear[0] * gear[1]
		sum2 += ratio
	}
	fmt.Println("part2:", sum2) // should be 75312571
}

type pos [2]int

type grid [][]rune

func (g grid) hasAdjacentSymbol(row int, col int) bool {
	right := pos{row, col + 1}
	rightDown := pos{row + 1, col + 1}
	down := pos{row + 1, col}
	leftDown := pos{row + 1, col - 1}
	left := pos{row, col - 1}
	leftUp := pos{row - 1, col - 1}
	up := pos{row - 1, col}
	rightUp := pos{row - 1, col + 1}
	adjacent := []pos{right, rightDown, down, leftDown, left, leftUp, up, rightUp}
	for i, adjPos := range adjacent {
		if adjPos[0] >= 0 && adjPos[0] < len(g) && adjPos[1] >= 0 && adjPos[1] < len(g[i]) {
			value := g[adjPos[0]][adjPos[1]]
			if value != '.' && !unicode.IsDigit(value) {
				return true
			}
		}
	}
	return false
}

func (g grid) adjacent(row int, col int) []pos {
	right := pos{row, col + 1}
	rightDown := pos{row + 1, col + 1}
	down := pos{row + 1, col}
	leftDown := pos{row + 1, col - 1}
	left := pos{row, col - 1}
	leftUp := pos{row - 1, col - 1}
	up := pos{row - 1, col}
	rightUp := pos{row - 1, col + 1}
	adjacent := []pos{right, rightDown, down, leftDown, left, leftUp, up, rightUp}

	var result []pos
	for i, adjPos := range adjacent {
		if adjPos[0] >= 0 && adjPos[0] < len(g) && adjPos[1] >= 0 && adjPos[1] < len(g[i]) {
			result = append(result, adjPos)
		}
	}
	return result
}

func parseGrid(input io.Reader) grid {
	var result [][]rune

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, []rune(line))
	}
	return result
}

func (g grid) findParts() []int {
	var result []int

	partNo := 0
	isValid := false
	for i, row := range g {
		for j, value := range row {
			if unicode.IsDigit(value) {
				partNo = partNo*10 + int(value-'0')
				if !isValid {
					isValid = g.hasAdjacentSymbol(i, j)
				}
			} else {
				if partNo > 0 && isValid {
					result = append(result, partNo)
				}
				partNo = 0
				isValid = false
			}
		}
	}

	return result
}

func (g grid) findParts2() []int {
	var results []int

	partNo := 0
	isValid := false
	for i, row := range g {
		for j, value := range row {
			if unicode.IsDigit(value) {
				partNo = partNo*10 + int(value-'0')
				for _, adjPos := range g.adjacent(i, j) {
					r := g[adjPos[0]][adjPos[1]]
					if isSymbol(r) {
						isValid = true
					}
				}
			} else {
				if partNo > 0 && isValid {
					results = append(results, partNo)
				}
				isValid = false
				partNo = 0
			}
		}
	}

	return results
}

func isSymbol(r rune) bool {
	if r != '.' && !unicode.IsDigit(r) {
		return true
	}
	return false
}

type part struct {
	no  int
	pos []pos
}

func parsePartsAndStars(reader io.Reader) ([]part, []pos) {
	var stars []pos
	var parts []part

	scanner := bufio.NewScanner(reader)
	i := 0
	partNo := 0
	var partPos []pos
	for scanner.Scan() {
		line := scanner.Text()
		for j, char := range line {
			if unicode.IsDigit(char) {
				partNo = partNo*10 + int(char-'0')
				partPos = append(partPos, pos{i, j})
			} else {
				if char == '*' {
					stars = append(stars, pos{i, j})
				}
				if partNo > 0 {
					parts = append(parts, part{no: partNo, pos: partPos})
				}
				partNo = 0
				partPos = nil
			}
		}
		i++
	}
	return parts, stars
}

func (g grid) findGears(parts []part, starPositions []pos) [][2]int {
	var gears [][2]int
	for _, starPos := range starPositions {
		var adjPartNos []int
		for _, adjPos := range g.adjacent(starPos[0], starPos[1]) {
			for _, part := range parts {
				if slices.Contains(part.pos, adjPos) && !slices.Contains(adjPartNos, part.no) {
					adjPartNos = append(adjPartNos, part.no)
				}
			}
		}
		if len(adjPartNos) == 2 {
			gears = append(gears, [2]int{adjPartNos[0], adjPartNos[1]})
		}
	}

	return gears
}
