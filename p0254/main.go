package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	var mp [1001]int
	ans := 0
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	fmt.Fscan(reader, &k)

	for i := range a {
		if a[i] < k {
			ans += mp[k-a[i]]
		}
		mp[a[i]]++
	}
	fmt.Fprintln(writer, ans)
}

func main() {
	solve(os.Stdin, os.Stdout)
}
