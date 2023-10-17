package main

import (
	"fmt"
	"bufio"
	"os"
	// "io"
)
/* 
input1
6 5
10 20 30 40 60

input2
15 10
10 20 30 40 60 60 70 80 90 150
*/
func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, n int
	var a []int
	fmt.Fscan(reader, &m, &n)

	a = make([]int, n) // [1, n+1]
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}

	f := make([]int, m + 1) // 1e6
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i - j >= 0 {
				f[i] = max(f[i], f[i - j] + a[j - 1])
			}
		}
	}
	fmt.Fprintln(writer, f[m])
}

func max(x, y int) int { if x > y { return x } else { return y }}