package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 用一个数组A代表程序员的工作能力，公司想通过结对编程的方式提高员工的能力，假设结对后的能力为两个员工的能力之和，求一共有多少种结对方式使结对后能力为N。

// 输入描述
// 5
// 1 2 2 2 3
// 4
// 第一行为员工的总人数，取值范围[1,1000]
// 第二行为数组 A
// A 的元素，每个元素的取值范围[1,1000]
// 第三行为 N
// N 的值，取值范围[1,1000]

// 输出描述
// 4
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
