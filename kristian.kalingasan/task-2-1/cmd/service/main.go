package main

import (
	"fmt"
)

const (
	minTemp = 15
	maxTemp = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner(low, high int) conditioner {
	return conditioner{
		low:  low,
		high: high,
	}
}

func (c *conditioner) apply(operator string, value int) error {
	switch operator {
	case ">=":
		if value > c.low {
			c.low = value
		}
	case "<=":
		if value < c.high {
			c.high = value
		}
	default:
		return fmt.Errorf("operator is not allowed")
	}

	return nil
}

func (c *conditioner) optimal() (int, bool) {
	if c.low > c.high {
		return 0, false
	}

	return c.low, true
}

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments", err)

		return
	}

	for range departments {
		var count int

		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("Invalid number of employees", err)

			return
		}

		cond := newConditioner(minTemp, maxTemp)

		for range count {
			var (
				value    int
				operator string
			)

			_, err = fmt.Scan(&operator, &value)
			if err != nil {
				fmt.Println("Invalid operator or value", err)

				return
			}

			err = cond.apply(operator, value)
			if err != nil {
				fmt.Println(err)

				return
			}

			val, ok := cond.optimal()
			if !ok {
				fmt.Println(-1)
			} else {
				fmt.Println(val)
			}
		}
	}
}
