package main

import "fmt"

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
)

type temperatureBounds struct {
	minTemp int
	maxTemp int
}

func newTemperatureBounds(minValue, maxValue int) temperatureBounds {
	return temperatureBounds{
		minTemp: minValue,
		maxTemp: maxValue,
	}
}

func (tb *temperatureBounds) updateBounds(operator string, targetTemp int) {
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
		fmt.Println("Предупреждение: неизвестный оператор:", operator)
	}
}

func (tb *temperatureBounds) getComfortTemp() (int, bool) {
	if tb.minTemp > tb.maxTemp {
		return 0, false
	}

	return tb.minTemp, true
}

func main() {
	var departmentsCount int

	if _, err := fmt.Scan(&departmentsCount); err != nil {
		fmt.Println("Ошибка чтения количества отделов:", err)

		return
	}

	for range departmentsCount {
		var employeesCount int

		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Ошибка чтения количества сотрудников:", err)

			return
		}

		bounds := newTemperatureBounds(defaultMinTemp, defaultMaxTemp)

		for range employeesCount {
			var (
				operator   string
				targetTemp int
			)

			if _, err := fmt.Scan(&operator, &targetTemp); err != nil {
				fmt.Println("Ошибка чтения данных сотрудника:", err)

				return
			}

			bounds.updateBounds(operator, targetTemp)

			if temp, ok := bounds.getComfortTemp(); ok {
				fmt.Println(temp)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
