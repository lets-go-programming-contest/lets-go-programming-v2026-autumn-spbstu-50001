package main

import "fmt"

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

type TemperatureBounds struct {
	minTemp int
	maxTemp int
}

func NewTemperatureBounds() TemperatureBounds {
	return TemperatureBounds{
		minTemp: DefaultMinTemp,
		maxTemp: DefaultMaxTemp,
	}
}

func (tb *TemperatureBounds) UpdateBounds(operator string, targetTemp int) {
	if operator == ">=" {
		if targetTemp > tb.minTemp {
			tb.minTemp = targetTemp
		}
	} else if operator == "<=" {
		if targetTemp < tb.maxTemp {
			tb.maxTemp = targetTemp
		}
	}
}

func (tb *TemperatureBounds) GetComfortTemp() int {
	if tb.minTemp > tb.maxTemp {
		return -1
	}
	return tb.minTemp
}

func main() {
	var departmentsCount int
	if _, err := fmt.Scan(&departmentsCount); err != nil {
		return
	}

	for deptIndex := 0; deptIndex < departmentsCount; deptIndex++ {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			return
		}

		bounds := NewTemperatureBounds()

		for empIndex := 0; empIndex < employeesCount; empIndex++ {
			var (
				operator   string
				targetTemp int
			)

			if _, err := fmt.Scan(&operator, &targetTemp); err != nil {
				return
			}

			bounds.UpdateBounds(operator, targetTemp)
			fmt.Println(bounds.GetComfortTemp())
		}
	}
}
