package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"sort"
	"math/big"
)

// ABA
// 3

// ABCDEFGHHA
// 907200

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	s, _ := reader.ReadString('\n')
	
	// 10  * 9 * 8 * 7 * 6
	bs := []byte(s[:len(s)-2])
	n := len(bs)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	var perm []byte
	var ans int
	var vis = make([]bool, n)
    var backtrack func(int)
    backtrack = func(idx int) {
        if idx == n {
            ans += 1
            return
        }
        for i, v := range bs {
            if vis[i] || i > 0 && !vis[i-1] && bs[i] == bs[i-1] {
                continue
            }
            perm = append(perm, byte(v))
            vis[i] = true
            backtrack(idx + 1)
            vis[i] = false
            perm = perm[:len(perm)-1]
        }
    }
    backtrack(0)
	
	fmt.Fprintln(writer, ans)
}


func factorial(n int64) *big.Int {
	res := big.NewInt(1)
	for i := int64(2); i <= n; i++ {
		res.Mul(res, big.NewInt(i))
	}
	return res
}
func solve1(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()
	
	for {
		var s string
		fmt.Fscan(reader, &s)
		var freq [26]int
		for _, c := range s {
			freq[c - 'A']++
		}
		ans := factorial(int64(len(s)))
		for _, f := range freq {
			if f > 1 { ans.Div(ans, factorial(int64(f))) }
		}
		fmt.Fprintln(writer, ans)
		break
	}
}
func main() {
	solve1(os.Stdin, os.Stdout)
}