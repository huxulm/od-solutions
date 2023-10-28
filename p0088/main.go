package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"sort"
)

func main() { solve1(os.Stdin, os.Stdout) }

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, x int
	var a []int
	fmt.Fscan(in, &n, &x)
	a = make([]int, n)
	ans := 0
	pre := []int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		if i == 0 {
			pre = append(pre, a[i])
		} else {
			pre = append(pre, pre[i-1] + a[i])
		}
		// 最大的 pre[j], pre[i] - pre[j] >= x
		// [0, j) 都 <= pre[i] - x
		j := sort.SearchInts(pre, pre[i] - x + 1) // 
		if pre[i] >= x {
			ans += j + 1
		}
	}
	fmt.Fprintln(out, ans)
}

// 滑动窗口
func solve1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, x int
	fmt.Fscan(in, &n, &x)
	a := make([]int, n)

	var s, j, ans int
	for i := range a {
		fmt.Fscan(in, &a[i])
		s += a[i]
		for s - a[j] >= x {
			s -= a[j]
			j++
		}
		if s >= x {
			ans += j + 1
		}
	}
	fmt.Fprintln(out, ans)
}