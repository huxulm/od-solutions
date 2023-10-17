package main
import (
	"fmt"
	// "io"
	"os"
	"bufio"
)

// 求从坐标零点到坐标n的最小步数，一次只能沿横坐标轴向左或向右移动2或3。
// 注意：途经的坐标可以为负数。

// 4
// 2

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	f1, f2, f3 := 2, 1, 1
	if n == 1 {
		fmt.Fprintln(writer, f1)
		return
	}
	if n == 2 || n == 3 {
		fmt.Fprintln(writer, f2)
		return
	}
	for i := 4; i <= n; i++ {
		f4 := min(f1, f2) + 1
		f1, f2, f3 = f2, f3, f4
	} 
	fmt.Fprintln(writer, f3)
}

func min(x, y int) int {
	if x < y { return x }
	return y
}