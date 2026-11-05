package main

import (
	"errors"
	"fmt"
)

const (
	minTemp    = 15
	maxTemp    = 30
	noSolution = -1
)

var ErrUnknownOperator = errors.New("unknown operator")

type Conditioner struct {
	low  int
	high int
}

func NewConditioner(low, high int) Conditioner {
	return Conditioner{
		low:  low,
		high: high,
	}
}

func (c *Conditioner) Apply(operator string, temp int) error {
	switch operator {
	case ">=":
		if temp > c.low {
			c.low = temp
		}
	case "<=":
		if temp < c.high {
			c.high = temp
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnknownOperator, operator)
	}

	return nil
}

func (c *Conditioner) Optimal() (int, bool) {
	if c.low > c.high {
		return 0, false
	}

	return c.low, true
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Error reading departments count:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Error reading employees count:", err)

			return
		}

		conditioner := NewConditioner(minTemp, maxTemp)

		for range employeesCount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error reading employee preference:", err)

				return
			}

			if err := conditioner.Apply(operator, temp); err != nil {
				fmt.Println("Error applying employee preference:", err)

				return
			}

			result, ok := conditioner.Optimal()
			if !ok {
				fmt.Println(noSolution)

				continue
			}

			fmt.Println(result)
		}
	}
}
