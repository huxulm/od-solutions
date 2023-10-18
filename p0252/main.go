package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)
func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var k int
	var s string
	fmt.Fscan(reader, &k)
	fmt.Fscan(reader, &s)
	hasFirst, cnt := 0, 0
	for i := range s {
		if hasFirst == 0 {
			fmt.Fprintf(writer, "%s", s[i:i+1])
			if s[i] == '-' {
				hasFirst = 1
			}
		} else {
			if s[i] == '-' { continue }
			fmt.Fprintf(writer, "%s", string(toupper(s[i])))
			cnt++
			if cnt % k == 0 {
				if i != len(s) - 1 {
					fmt.Fprint(writer, "-")
				}
				cnt %= k
			}
		}
	}
	fmt.Fprint(writer, "\n")
}

func toupper(ch byte) byte {
	if 'a' <= ch && ch <= 'z' {
		return byte(ch - 'a' + 'A')
	}
	return ch
}
func main() {
	solve(os.Stdin, os.Stdout)
}