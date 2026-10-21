package main

import (
	"fmt"

	"github.com/kirillkirsanov2007/task-2-1/internal/climate"
)

const (
	minTemp = 15
	maxTemp = 30
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		currentRange := climate.NewTemperatureRange(minTemp, maxTemp)

		for j := 0; j < k; j++ {
			var (
				cmp  string
				temp int
			)
			if _, err := fmt.Scan(&cmp, &temp); err != nil {
				return
			}

			constraint, err := climate.NewConstraint(cmp, temp)
			if err != nil {
				return
			}

			currentRange.Apply(constraint)

			optimal, err := currentRange.GetOptimal()
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(optimal)
			}
		}
	}
}
