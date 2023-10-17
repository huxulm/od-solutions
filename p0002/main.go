package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"sort"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

type pair struct{
	i, j, val int
}

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var n, m, k int
	var arr1, arr2 []int

	var nodes []pair
	fmt.Fscan(reader, &n)
	arr1 = make([]int, n)
	for i := range arr1 {
		fmt.Fscan(reader, &arr1[i])
	}

	fmt.Fscan(reader, &m)
	arr2 = make([]int, m)
	for i := range arr2 {
		fmt.Fscan(reader, &arr2[i])
	}

	fmt.Fscan(reader, &k)

	for i, x := range arr1 {
		for j, y := range arr2 {
			nodes = append(nodes, pair{i, j, x + y})
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val
	})

	used := make(map[pair]bool, m * n)
	var ans, cnt int

	for _, p := range nodes {
		if used[pair{p.i, p.j, 0}] {
			continue
		}
		ans += p.val
		used[pair{p.i, p.j, 0}] = true
		cnt++
		if cnt == k {
			break
		}
	}

	fmt.Fprintln(writer, ans)
}