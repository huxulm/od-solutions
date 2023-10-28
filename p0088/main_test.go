package main

import (
	"testing"
	"strings"
	"os"
)

func TestXxx1(t *testing.T) {
	inputs := []string{
		`3 7
		3 4 7`,
		`10 10000000
		1 2 3 4 5 6 7 8 9 10`,
	}

	for _, input := range inputs {
		solve(strings.NewReader(input), os.Stdout)
	}
}

func TestXxx2(t *testing.T) {
	inputs := []string{
		`3 7
		3 4 7`,
		`10 10000000
		1 2 3 4 5 6 7 8 9 10`,
	}

	for _, input := range inputs {
		solve1(strings.NewReader(input), os.Stdout)
	}
}