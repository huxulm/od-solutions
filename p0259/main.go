package main

import (
	"fmt"
	"os"
	"bufio"
)
func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string

	fmt.Fscan(reader, &s)

	tot := 0
	for _, c := range s {
		tot += int(c)
	}
	if tot % 3 != 0 {
		fmt.Fprintln(writer, "0 0")
		return
	}

	target := tot / 3
	mp := map[int]int{}
	pre := 0
	for i, c := range s {
		if pre % 2 != 0 {
			continue
		}
		if j, ok := mp[target]; ok && target * 2 == pre {
			fmt.Fprintf(writer, "%d %d", j, i)
			return
		}
		pre += int(c)
		if i > 0 {
		  mp[pre] = i
		}
	}
	fmt.Fprintln(writer, "0 0")
}