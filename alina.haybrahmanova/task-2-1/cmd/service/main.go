package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var k int
		if _, err := fmt.Fscan(reader, &k); err != nil {
			return
		}

		lo, hi := 15, 30

		for j := 0; j < k; j++ {
			var op string
			var t int
			if _, err := fmt.Fscan(reader, &op, &t); err != nil {
				return
			}

			switch op {
			case ">=":
				if t > lo {
					lo = t
				}
			case "<=":
				if t < hi {
					hi = t
				}
			}

			if lo > hi {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, lo)
			}
		}
	}
}
