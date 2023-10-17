package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(_r io.Reader, _w io.Writer) {
	reader, writer := bufio.NewReader(_r), bufio.NewWriter(_w)
	defer writer.Flush()

	var tag string

	fmt.Fscan(reader, &tag)

	for {
		var curTag string
		fmt.Fscan(reader, &curTag)
		var curLen int
		var s string
		fmt.Fscan(reader, &s)
		v, _ := strconv.ParseInt(s, 16, 10)
		curLen |= int(v)
		fmt.Fscan(reader, &s)
		v, _ = strconv.ParseInt(s, 16, 10)
		curLen |= int(v << 8)
		var ok bool
		for i := 0; i < curLen; i++ {
			fmt.Fscan(reader, &s)
			if tag == curTag {
				ok = true
				fmt.Fprintf(writer, "%s ", s)
			}
		}
		if ok { break }
	}
}