package main

import (
	"testing"
	"os"
	"strings"
)

func TestXxx(t *testing.T) {
	inputs := []string{
		`3
		3 5 6`,
		`8
		7258 6579 2602 6716 3050 3564 5396 1773`,
	}

	for _, in := range inputs {
		solve(strings.NewReader(in), os.Stdout)
	}
}