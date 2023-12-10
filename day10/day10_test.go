package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	var input = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

	pipeMap, start := parseInput(strings.NewReader(input))

	loopTiles, err := pipeMap.findLoopTiles(start, 'F', west)
	if err != nil {
		t.Fatalf(err.Error())
	}

	length := len(loopTiles)
	if length != 16 {
		t.Errorf("length %d, expected 16", length)
	}
}
