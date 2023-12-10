package main

import "C"
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day8/input-day8")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	instructions, nodeIndex := parseInput(file)

	start := "AAA"
	count, _, _ := countSteps(instructions, nodeIndex, 0, start)
	fmt.Println("part1:", count) // should be 16271

	var startPositions []string
	for _, n := range nodeIndex {
		if strings.HasSuffix(n.value, "A") {
			startPositions = append(startPositions, n.value)
		}
	}

	part2(instructions, nodeIndex)
}

func part2(instructions []int, nodeIndex map[string]node) {
	startPositions := findStarts(nodeIndex)
	findLoopLength(instructions, nodeIndex, startPositions)
}

type node struct {
	value string
	next  []string
}

func parseInput(reader io.Reader) ([]int, map[string]node) {
	var instructions []int
	nodeIndex := make(map[string]node)

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		instructionLine := scanner.Text()
		for _, char := range instructionLine {
			if char == 'L' {
				instructions = append(instructions, 0)
			} else if char == 'R' {
				instructions = append(instructions, 1)
			}
		}
	}
	scanner.Scan() // skip one line
	for scanner.Scan() {
		n := parseNode(scanner.Text())
		nodeIndex[n.value] = n
	}

	return instructions, nodeIndex
}

func parseNode(s string) node {
	split := strings.Split(s, "=")
	value := split[0]

	leftRight := strings.Trim(split[1], " ()")
	split2 := strings.Split(leftRight, ",")
	next := []string{strings.TrimSpace(split2[0]), strings.TrimSpace(split2[1])}
	return node{
		value: strings.TrimSpace(value),
		next:  next,
	}
}

func foundEnd(position string) bool {
	if !strings.HasSuffix(position, "Z") {
		return false
	}

	return true
}

// countSteps returns the number of steps, and additionally the next instruction index, and the encoding of the position
func countSteps(instructions []int, nodeIndex map[string]node, startInstruction int, startPosition string) (int, int, string) {
	position := startPosition
	i := startInstruction
	count := 0
	for !foundEnd(position) {
		if i >= len(instructions) {
			i = 0
		}

		instruction := instructions[i]
		position = nodeIndex[position].next[instruction]
		count++

		i++
	}

	return count, i, position
}

func findStarts(nodeIndex map[string]node) []string {
	var startPositions []string
	for _, n := range nodeIndex {
		if strings.HasSuffix(n.value, "A") {
			startPositions = append(startPositions, n.value)
		}
	}
	return startPositions
}

func findLoopLength(instructions []int, nodeIndex map[string]node, startPositions []string) {
	fmt.Println("# nodes:", len(nodeIndex))
	fmt.Println("len instr", len(instructions))

	part2 := 1
	for i, startPos := range startPositions {
		fmt.Printf("%d start: %s -> %s from node %v\n", i+1, startPos, nodeIndex[startPos].next[instructions[0]], nodeIndex[startPos])

		count, nextI, endPosition := countSteps(instructions, nodeIndex, 0, startPos)
		fmt.Println("- count:", count, "next instr", nextI, "or", nextI%len(instructions))

		// next progression / start of loop
		fmt.Println("-- count / len(inst)", float32(count)/float32(len(instructions)))
		part2 *= count / len(instructions)

		// progress once
		instruction := instructions[nextI%len(instructions)]
		nextPos := nodeIndex[endPosition].next[instruction]

		fmt.Printf("--- end: %s -> %s from node %v\n", endPosition, nextPos, nodeIndex[endPosition])

		//// count the steps in the cycle
		count2, _, loopEnd := countSteps(instructions, nodeIndex, nextI+1, nextPos)
		fmt.Println("-loop end", loopEnd)
		fmt.Println("-- cycle count", count2)
		fmt.Println("-- start len", count-count2)
		fmt.Println("---- cycle divided len(instr)", float32(count2)/float32(len(instructions)))
		fmt.Println("---- startLen divided len(instr)", float32(count-count2)/float32(len(instructions)))
	}
	fmt.Println(part2)

	// found that a) the length of the instruction set is a prime (and a factor of the loops)
	//            b) all path lengths divided by this prime are prime again
	// so the LCM of the path lengths is the product of all of these primes
	fmt.Println(part2 * len(instructions)) // This is the least common multiplier of all path lenghts

	// the amount of nodes is exactly 2 * len(instructions) --> limited amount of states you can end up with
}
