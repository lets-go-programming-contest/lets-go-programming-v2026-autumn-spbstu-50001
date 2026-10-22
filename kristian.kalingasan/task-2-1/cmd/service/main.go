package main

import (
	"fmt"
)

func main() {
	var N int

	_, _ = fmt.Scan(&N)

	for range N {
		var K int

		_, _ = fmt.Scan(&K)

		low, high := 15, 30

		for range K {
			var (
				value    int
				operator string
			)

			_, _ = fmt.Scan(&operator, &value)

			if operator == ">=" && value > low {
				low = value
			} else if operator == "<=" && value < high {
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
