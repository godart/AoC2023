package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type pos struct {
	row int
	col int
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

type grid [][]rune

func main() {
	file, _ := os.Open("day10/input-day10")
	pipeMap, start := parseInput(file)

	loopTiles, err := pipeMap.findLoopTiles(start, 'L', south) // only possible pipe at start is 'L', just pick a valid direction
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("part1:", len(loopTiles)/2) // should be 6947

	// -------- Part2 ------

	var loopOnlyMap [][]rune
	for rowCnt, mapRow := range pipeMap {
		cleanedRow := slices.Clone(mapRow)
		for colCnt, _ := range cleanedRow {
			if !slices.Contains(loopTiles, pos{row: rowCnt, col: colCnt}) {
				cleanedRow[colCnt] = '.'
			}
		}
		loopOnlyMap = append(loopOnlyMap, cleanedRow)
	}

	inside := false
	insideCount := 0
	for _, mapRow := range loopOnlyMap {
		openTurn := ' '
		for _, tile := range mapRow {
			switch tile {
			case '|':
				inside = !inside
			case 'F':
				openTurn = 'F'
			case 'L':
				openTurn = 'L'
			case '7':
				if openTurn == 'L' {
					inside = !inside
				}
				openTurn = ' '
			case 'J':
				if openTurn == 'F' {
					inside = !inside
				}
				openTurn = ' '
			case '.':
				if inside {
					insideCount++
				}
			}
		}
	}
	fmt.Println("part2:", insideCount) // should be 273
}

func (g grid) findLoopTiles(start pos, startTile rune, lastDir direction) ([]pos, error) {
	var loopTiles []pos

	curPos := start
	g[start.row][start.col] = startTile
	var err error
	for {
		loopTiles = append(loopTiles, curPos)
		curPos, lastDir, err = g.followPipe(curPos, lastDir)
		if err != nil {
			return nil, err
		}
		if curPos == start {
			break
		}
	}

	return loopTiles, nil
}

func parseInput(input io.Reader) (grid, pos) {
	grid := grid{}
	var start pos
	scanner := bufio.NewScanner(input)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		gridLine := []rune(line)
		grid = append(grid, gridLine)

		animalColumn := strings.Index(line, "S")
		if animalColumn != -1 {
			start = pos{row: i, col: animalColumn}
		}
		i++
	}

	return grid, start
}

func (g grid) followPipe(curPos pos, lastDir direction) (pos, direction, error) {
	curTile := '.' // set boundary to '.'
	if curPos.row >= 0 || curPos.row < len(g) || curPos.col >= 0 || curPos.col < len(g[curPos.row]) {
		curTile = g[curPos.row][curPos.col]
	}
	comingFrom := (lastDir + 2) % 4

	var goingTo direction
	switch curTile {
	case '|':
		switch comingFrom {
		case north:
			goingTo = south
		case south:
			goingTo = north
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case '-':
		switch comingFrom {
		case west:
			goingTo = east
		case east:
			goingTo = west
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case 'L':
		switch comingFrom {
		case north:
			goingTo = east
		case east:
			goingTo = north
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case 'J':
		switch comingFrom {
		case north:
			goingTo = west
		case west:
			goingTo = north
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case 'F':
		switch comingFrom {
		case south:
			goingTo = east
		case east:
			goingTo = south
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case '7':
		switch comingFrom {
		case south:
			goingTo = west
		case west:
			goingTo = south
		default:
			return pos{}, 0, fmt.Errorf("wrong connection: %s - %d", string(curTile), comingFrom)
		}
	case '.':
		return pos{}, 0, fmt.Errorf("%v is a ground tile, not a pipe", curPos)
	default:
		return pos{}, 0, fmt.Errorf("invalid tile %s", string(curTile))
	}

	var nextPos pos
	if goingTo == north {
		nextPos = pos{row: curPos.row - 1, col: curPos.col}
	} else if goingTo == south {
		nextPos = pos{row: curPos.row + 1, col: curPos.col}
	} else if goingTo == west {
		nextPos = pos{row: curPos.row, col: curPos.col - 1}
	} else if goingTo == east {
		nextPos = pos{row: curPos.row, col: curPos.col + 1}
	}

	return nextPos, goingTo, nil
}
