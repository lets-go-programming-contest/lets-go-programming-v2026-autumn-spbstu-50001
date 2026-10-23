package main

import (
	"errors"
	"fmt"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

type conditioner struct {
	low  int
	high int
}

func newConditioner(low, high int) conditioner {
	return conditioner{low: low, high: high}
}

func (c *conditioner) addPreference(sign string, temperature int) error {
	switch sign {
	case ">=":
		if temperature > c.low {
			c.low = temperature
		}
	case "<=":
		if temperature < c.high {
			c.high = temperature
		}
	default:

		return fmt.Errorf("unknown sign: %s", sign)
	}

	return nil
}

func (c *conditioner) optimal() (int, error) {
	if c.low > c.high {

		return 0, errors.New("no optimal temperature")
	}

	return c.low, nil
}

func main() {
	var departmentsNum int
	if _, err := fmt.Scan(&departmentsNum); err != nil {
		fmt.Println("failed to read number of departments:", err)

		return
	}

	for range departmentsNum {
		var employeesNum int
		if _, err := fmt.Scan(&employeesNum); err != nil {
			fmt.Println("failed to read number of employees:", err)

			return
		}

		conditioner := newConditioner(minTemperature, maxTemperature)

		for range employeesNum {
			var (
				sign        string
				temperature int
			)

			_, err := fmt.Scan(&sign, &temperature)
			if err != nil {
				fmt.Println("failed to read preference:", err)

				return
			}

			if err = conditioner.addPreference(sign, temperature); err != nil {
				fmt.Println(err)

				return
			}

			temp, err := conditioner.optimal()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(temp)
		}
	}
}
