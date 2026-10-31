package main

import (
	"fmt"
)

type Conditioner struct {
	low  int
	high int
}

func newConditioner() *Conditioner {
	return &Conditioner{
		low:  LOW,
		high: HIGH,
	}
}

func (c *Conditioner) apply(operator string, value int) {
	if operator == ">=" && value > c.low {
		c.low = value
	} else if operator == "<=" && value < c.high {
		c.high = value
	}
}

func (c *Conditioner) isValid() bool {
	return c.low <= c.high
}

func (c *Conditioner) optimal() int {
	return c.low
}

const (
	LOW  = 15
	HIGH = 30
)

func main() {
	var departments int

	_, err := fmt.Scan(&departments)
	if err != nil {
		fmt.Println("Invalid number of departments")

		return
	}

	for range departments {
		var count int

		_, err = fmt.Scan(&count)
		if err != nil {
			fmt.Println("Invalid number of employees")

			return
		}

		conditioner := newConditioner()

		for range count {
			var (
				value    int
				operator string
			)

			_, err = fmt.Scan(&operator, &value)
			if err != nil {
				fmt.Println("Invalid operator or value")

				return
			}

			if !(operator == ">=" || operator == "<=") {
				fmt.Println("Operator is not allowed")

				return
			}

			if value > HIGH || value < LOW {
				fmt.Println("Value is out of range")

				return
			}

			conditioner.apply(operator, value)

			if !conditioner.isValid() {
				fmt.Println(-1)
			} else {
				fmt.Println(conditioner.optimal())
			}
		}
	}
}
