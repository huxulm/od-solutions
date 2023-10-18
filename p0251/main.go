package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}

	var getScore = func(arr []int) int {
		n := len(arr)
		// f[i] = max(f[i-1], arr[i-1], f[i-2] + arr[i-1])
		f := make([]int, n + 1)
		for i := 1; i <= n; i++ {
			f[i] = max(f[i - 1], arr[i-1])
			if i >= 2 {
				f[i] = max(f[i], f[i - 2] + arr[i - 1])
			}
		}
		return f[n]
	}
	fmt.Fprintln(writer, max(getScore(a[:n-1]), getScore(a[1:])))
}

func main() {
	solve(os.Stdin, os.Stdout)
}

func max(x, y int) int { if x > y { return x } else { return y }}