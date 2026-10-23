package main

import "fmt"

const (
	tmin = 15
	tmax = 30
)

type Conditioner struct {
	low  int
	high int
}

func (c *Conditioner) Apply(operator string, temp int) int {
	switch operator {
	case ">=":
		c.low = max(c.low, temp)
	case "<=":
		c.high = min(c.high, temp)
	}

	if c.low > c.high {
		return -1
	}

	return c.low
}

func main() {
	var numDepart int

	if _, err := fmt.Scan(&numDepart); err != nil {
		fmt.Println("Error: number of departments must be an integer")

		return
	}

	for range numDepart {
		var numEmployee int

		if _, err := fmt.Scan(&numEmployee); err != nil {
			fmt.Println("Error: number of employees must be an integer")

			return
		}

		cond := Conditioner{low: tmin, high: tmax}

		for range numEmployee {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error: invalid input format")

				return
			}

			fmt.Println(cond.Apply(operator, temp))
		}
	}
}
