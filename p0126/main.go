package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var cur int = 1
	for i := 0; i < n; i++ {
		indent := (n - 1 - i) * 4
		// print indent
		for j := 0; j < indent; j++ {
			fmt.Fprint(writer, " ") 
		}
		for k := 0; k < i+1; k++ {
			s := strconv.Itoa(cur)
			for t := 0; t < 5 - len(s); t++ {
				s += "*"
			}
			fmt.Fprintf(writer, "%s    ", s)
			cur++
		}
		fmt.Fprint(writer, "\n")
	}
}