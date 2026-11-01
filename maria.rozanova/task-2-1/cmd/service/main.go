package main

import (
	"fmt"
)

const (
	tmin = 15
	tmax = 30
)

type Conditioner struct {
	low  int
	high int
}

func NewConditioner() *Conditioner {
	return &Conditioner{
		low:  tmin,
		high: tmax,
	}
}

func (c *Conditioner) AddPreference(operator string, temp int) (int, error) {
	switch operator {
	case ">=":
		c.low = max(c.low, temp)
	case "<=":
		c.high = min(c.high, temp)
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}

	if c.low > c.high {
		return 0, fmt.Errorf("no valid temperature")
	}

	return c.low, nil
}

func main() {
	var numDepart int

	if _, err := fmt.Scan(&numDepart); err != nil {
		fmt.Println("Error reading number of departments:", err)

		return
	}

	for range numDepart {
		var numEmployee int

		if _, err := fmt.Scan(&numEmployee); err != nil {
			fmt.Println("Error reading number of employees:", err)

			return
		}

		cond := NewConditioner()

		for range numEmployee {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error reading preference:", err)

				return
			}

			result, err := cond.AddPreference(operator, temp)
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(result)
		}
	}
}
