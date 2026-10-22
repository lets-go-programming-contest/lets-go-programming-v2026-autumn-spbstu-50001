package main

import (
	"fmt"
)

func main() {
	var N int

	_, _ = fmt.Scan(&N)

	for range N {
		var K int
		low, high := 15, 30

		_, _ = fmt.Scan(&K)

		for range K {
			var (
				value int
				op    string
			)

			_, _ = fmt.Scan(&op, &value)

			if op == ">=" && value > low {
				low = value
			} else if op == "<=" && value < high {
				high = value
			}

			if low > high {
				fmt.Println(-1)
			} else {
				fmt.Println(low)
			}
		}
	}
}
