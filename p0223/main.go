package main

import (
	"io"
	"fmt"
	"bufio"
	"strings"
)

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var mp []string = []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "st", "uv", "wx", "yz"}
	var s, v string

	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &v)

	bans := 0
	for i := range v {
		bans |= 1 << (v[i] - 'a')
	}
	seq := make([]int, len(s))
	for i := range s {
		seq[i] = int(15 & s[i])
	}

	var ans []string
	var dfs func(i, mask int)
	var perm []byte

	dfs = func(i, mask int) {
		if i == len(seq) {
			if mask != bans {
				ans = append(ans, string(append([]byte(nil), perm...)))
			}
			return
		}
		for _, c := range mp[seq[i]] {
			perm = append(perm, byte(c))
			if (1 << (c - 'a')) & bans > 0 { mask |= 1 << (c - 'a') }
			dfs(i + 1, mask)
			if (1 << (c - 'a')) & bans > 0 { mask ^= 1 << (c - 'a') }
			perm = perm[:len(perm) - 1]
		}
	}
	dfs(0, 0)
	fmt.Fprintln(writer, strings.Join(ans, ","))
}
