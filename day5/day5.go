package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day5/input-day5")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	seeds, maps, err := parseInput(file)
	locations := mapToLocations(seeds, maps)

	min1 := math.MaxInt
	for _, location := range locations {
		if location < min1 {
			min1 = location
		}
	}
	fmt.Println("part1:", min1) // should be 57075758

	min2 := 0
	for ; ; min2++ {
		if locationMapsSeed(min2, maps, seeds) {
			break
		}
	}
	fmt.Println("part2:", min2)
}

type mapping struct {
	FromStart int
	ToStart   int
	Length    int
}

type categoryMap struct {
	From     string
	To       string
	Mappings []mapping
}

func (cMap categoryMap) resolve(values []int) []int {
	result := values
	for i, value := range values {
		for _, m := range cMap.Mappings {
			if value >= m.FromStart && value < m.FromStart+m.Length {
				result[i] = m.ToStart + value - m.FromStart
			}
		}
	}
	return result
}

func (cMap categoryMap) resolveBackwards(value int) int {
	result := value
	for _, m := range cMap.Mappings {
		if result >= m.ToStart && result < m.ToStart+m.Length {
			result = m.FromStart + result - m.ToStart
			break
		}
	}
	return result
}

func parseInput(input io.Reader) ([]int, []categoryMap, error) {
	var seeds []int
	var maps []categoryMap

	var mapLines []string
	var from string
	var to string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		var err error
		if strings.HasPrefix(line, "seeds: ") {
			seeds, err = parseSeeds(line[len("seeds: "):])
			if err != nil {
				return nil, nil, err
			}
		} else if strings.HasSuffix(line, " map:") {
			toParse := strings.Split(line, " ")[0]
			mapHeader := strings.Split(toParse, "-to-")
			if len(mapHeader) != 2 {
				return nil, nil, fmt.Errorf("invalid header line: %s", line)
			}
			from = mapHeader[0]
			to = mapHeader[1]
		} else if line == "" {
			if mapLines != nil {
				curMapping, err := parseMapping(mapLines)
				if err != nil {
					return nil, nil, err
				}
				maps = append(maps, categoryMap{From: from, To: to, Mappings: curMapping})
			}

			from = ""
			to = ""
			mapLines = nil
		} else {
			mapLines = append(mapLines, line)
		}
	}
	if mapLines != nil {
		curMapping, err := parseMapping(mapLines)
		if err != nil {
			return nil, nil, err
		}
		maps = append(maps, categoryMap{From: from, To: to, Mappings: curMapping})
	}

	return seeds, maps, nil
}

func parseSeeds(seedLine string) ([]int, error) {
	var result []int
	for _, numString := range strings.Split(seedLine, " ") {
		seed, err := strconv.Atoi(numString)
		if err != nil {
			return nil, fmt.Errorf("seedLine: %s, %s", seedLine, err)
		}
		result = append(result, seed)
	}
	return result, nil
}

func parseMapping(lines []string) ([]mapping, error) {
	var result []mapping
	for _, line := range lines {
		splitSpace := strings.Split(line, " ")
		if len(splitSpace) != 3 {
			return nil, fmt.Errorf("illegal mapping line: %s", line)
		}
		fromStart, err := strconv.Atoi(splitSpace[1])
		toStart, err := strconv.Atoi(splitSpace[0])
		length, err := strconv.Atoi(splitSpace[2])
		if err != nil {
			return nil, fmt.Errorf("illegal mapping numbers: %s", line)
		}
		result = append(result, mapping{
			FromStart: fromStart,
			ToStart:   toStart,
			Length:    length,
		})
	}
	return result, nil
}

func mapToLocations(seeds []int, maps []categoryMap) []int {
	curValues := slices.Clone(seeds)
	for _, aMapping := range maps {
		curValues = aMapping.resolve(curValues)
	}

	return curValues
}

func locationMapsSeed(location int, maps []categoryMap, seedRanges []int) bool {
	mapped := location
	mapsLocal := slices.Clone(maps)
	slices.Reverse(mapsLocal)
	for _, m := range mapsLocal {
		mapped = m.resolveBackwards(mapped)
	}

	for i := 0; i < len(seedRanges); i = i + 2 {
		start := seedRanges[i]
		end := start + seedRanges[i+1]
		if mapped >= start && mapped < end {
			return true
		}
	}
	return false
}
