package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	fmt.Fscan(in, &s, &t)

	cnt := map[byte]int{} // s 字符个数
	for i := range s {
		cnt[s[i]]++
	}
	var check = func(c1, c2 map[byte]int) bool { // c2 每个字符个数 <= c1 每个字符个数
		for ch, v := range c2 {
			if v > c1[ch] {
				return false
			}
		}
		return true
	}

	var checkEqual = func(c1, c2 map[byte]int) bool {
		for ch := byte('a'); ch <= 'z'; ch++ {
			if c1[ch] != c2[ch] {
				return false
			}
		}
		return true
	}

	win := map[byte]int{}
	j := 0
	for i := range t {
		win[t[i]]++
		for !check(cnt, win) {
			win[t[j]]--
			j++
		}
		if checkEqual(cnt, win) && j <= i {
			fmt.Fprintln(out, j)
			return
		}
	}
	fmt.Fprintln(out, -1)
}
func main() {
	solve(os.Stdin, os.Stdout)
}