package main

import (
	"fmt"
	"os"
	// "io"
	"bufio"
	"sort"
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}

	// 单调递增，最大的 j pre[j] 使得 pre[-1] - pre[j] <= x
	pre := []int{0}
	ans := 0
	// f(j) = [false ... true]
	for _, v := range a {
		cur := pre[len(pre) - 1] + v
		pre = append(pre, cur)
		// 右边界
		j := sort.Search(len(pre), func(j int) bool {
			return cur - pre[j] < x
		})
		ans += j // 个数
	}

	fmt.Fprintln(writer, ans)
}