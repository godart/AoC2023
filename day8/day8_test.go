package main

import (
	"testing"
)

func TestFoundEnd2(t *testing.T) {
	tests := []struct {
		name     string
		position string
		exp      bool
	}{
		{
			name:     "one Z",
			position: "AAZ", //"ABZ", "BAZ",
			exp:      true,
		},
		{
			name:     "z middle",
			position: "AZA", //"BZA", "XZA", "ZZA", "ZZB",
			exp:      false,
		},
		{
			name:     "two Z",
			position: "AZZ", //"BZZ", "CZZ", "XZZ",
			exp:      true,
		},
		{
			name:     "three Z",
			position: "ZZZ",
			exp:      true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := foundEnd(test.position)
			if got != test.exp {
				t.Fatalf("%s: %t, got %t", test.position, got, test.exp)
			}
		})
	}
}
