package main

import (
	"testing"
	"os"
	"strings"
)

func TestXxx(t *testing.T) {
	inputs := []string{
		"abc efghicabiii",
		"abc efghicaibii",
	}

	for _, input := range inputs {
		solve(strings.NewReader(input), os.Stdout)
	}
}
