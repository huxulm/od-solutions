package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	f := make([]int, n + 1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] += f[i-1]
		if i >= 3 {
			f[i] += f[i-3]
		}
	}
	fmt.Fprintln(writer, f[n])
}

func main() {solve(os.Stdin, os.Stdout)}