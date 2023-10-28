package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

// AB两个人把苹果分为两堆
// A希望按照他的计算规则等分苹果
// 他的计算规则是按照二级制加法计算
// 并且不计算进位12+5=9(1100+0101=9),
// B的计算规则是十进制加法,
// 包括正常进位,B希望在满足A的情况下获取苹果重量最多
// 输入苹果的数量和每个苹果重量
// 输出满足A的情况下B获取的苹果总重量
// 如果无法满足A的要求 输出-1
// 数据范围：
// 1 <= 苹果数量 <= 20000
// 1 <= 每个苹果重量 <= 10000

// 输入1
// 3
// 3 5 6
// 输出1
// 11

// 输入2
// 8
// 7258 6579 2602 6716 3050 3564 5396 1773
// 输出2
// 35165
func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int 
	fmt.Fscan(in, &n)
	a := make([]int, n)
	aw, bw := 0, 0
	mw := 10001
	for i := range a {
		fmt.Fscan(in, &a[i])
		aw ^= a[i]
		bw += a[i]
		mw = min(mw, a[i])
	}

	if aw == 0 {
		fmt.Fprintln(out, bw - mw)
	} else {
		fmt.Fprintln(out, -1)
	}
}

func min(x, y int) int { if x < y { return x } else { return y }}

func main() {
	solve(os.Stdin, os.Stdout)
}

