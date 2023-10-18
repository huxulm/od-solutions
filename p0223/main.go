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

	var mp [][]byte = [][]byte{
		{'a','b','c'},
		{'d','e','f'},
		{'g','h','i'},
		{'j','k','l'},
		{'m','n','o'},
		{'p','q','r'},
		{'s','t'},
		{'u','v'},
		{'w','x'},
		{'y','z'},
	}

	var s string
	var v string

	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &v)

	bans := [26]bool{}
	for i := range v {
		bans[v[i] - 'a'] = true
	}
	seq := make([]int, len(s))
	for i := range s {
		seq[i] = int(15 & s[i])
	}

	var ans []string
	var dfs func(i int)
	var perm []byte

	dfs = func(i int) {
		if i == len(seq) {
			// check
			bs := [26]bool{}
			for j := range perm {
				bs[perm[j] - 'a'] = true
			}
			if bs != bans {
				ans = append(ans, string(append([]byte(nil), perm...)))
			}
			return
		}
		for _, c := range mp[seq[i]] {
			perm = append(perm, c)
			dfs(i + 1)
			perm = perm[:len(perm) - 1]
		}
	}
	dfs(0)
	fmt.Fprintln(writer, strings.Join(ans, ","))
}
