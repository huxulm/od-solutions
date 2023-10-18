package main

import (
	"os"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {

	t.Run("#test1", func(t *testing.T) {
		in := strings.NewReader(`5
		1 2 2 2 3
		4`)
		solve(in, os.Stdout)
	})
}
