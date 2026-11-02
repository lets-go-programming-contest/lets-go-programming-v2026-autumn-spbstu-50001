package main

import (
	"errors"
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

var (
	ErrNoComfortTemp   = errors.New("no comfortable temperature")
	ErrUnknownOperator = errors.New("unknown operator")
)

type temperatureBounds struct {
	minTemp int
	maxTemp int
}

func newTemperatureBounds(lower, upper int) temperatureBounds {
	return temperatureBounds{
		minTemp: lower,
		maxTemp: upper,
	}
}

func (tb *temperatureBounds) updateBounds(operator string, targetTemp int) error {
	switch operator {
	case ">=":
		if targetTemp > tb.minTemp {
			tb.minTemp = targetTemp
		}
	case "<=":
		if targetTemp < tb.maxTemp {
			tb.maxTemp = targetTemp
		}
	default:
		return fmt.Errorf("%w: %s", ErrUnknownOperator, operator)
	}

	return nil
}

func (tb *temperatureBounds) getComfortTemp() (int, error) {
	if tb.minTemp > tb.maxTemp {
		return 0, ErrNoComfortTemp
	}

	return tb.minTemp, nil
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

		bounds := newTemperatureBounds(defaultMinTemp, defaultMaxTemp)

		for range employeesCount {
			var (
				operator   string
				targetTemp int
			)

			if _, err := fmt.Scan(&operator, &targetTemp); err != nil {
				fmt.Println("Error reading employee data:", err)

				return
			}

			if err := bounds.updateBounds(operator, targetTemp); err != nil {
				fmt.Println("Error:", err)

				return
			}

			temp, err := bounds.getComfortTemp()
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(temp)
			}
		}
	}
}
