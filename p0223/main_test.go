package main
import (
	"testing"
	"os"
	"strings"
)
func TestXXX(t *testing.T) {
	t.Run("#test1", func(t *testing.T) {
		in := strings.NewReader(`78
		ux`)
		solve(in, os.Stdout)
	})
	t.Run("#test2", func(*testing.T) {
		in := strings.NewReader(`78
		x`)
		solve(in, os.Stdout)
	})
}