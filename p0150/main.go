package main

import (
	"fmt"
	// "io"
	"os"
	"sort"
	"bufio"
)
/*
输入1
5 3
4 5 3 5 5
输出1
5

输入2
5 2
4 5 3 5 5
输出2
4
*/
func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	var mx int = 10000000
	for i := range a {
		fmt.Fscan(reader, &a[i])
		mx = min(mx, a[i])
	}
	// f(x) = [false ... true]
	ans := sort.Search(mx + m + 1, func(x int) bool {
		tot := 0
		for _, v := range a {
			if v < x {
				tot += x - v
			}
		}
		return tot > m
	})
	fmt.Fprintln(writer, ans - 1)
}

func min(x, y int) int { if x < y { return x } else { return y }}