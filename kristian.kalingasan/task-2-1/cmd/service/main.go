package main

import (
	"fmt"
)

func main() {
	var departments int

	_, _ = fmt.Scan(&departments)

	for range departments {
		var count int

		_, _ = fmt.Scan(&count)

		low, high := 15, 30

		for range count {
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
