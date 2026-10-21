package main

import "fmt"

const (
	tmin = 15
	tmax = 30
)

func main() {
	var numDepart int
	if _, err := fmt.Scan(&numDepart); err != nil {
		fmt.Println("Ошибка! Количество отделов должно быть целым числом!")
		return
	}

	for range numDepart {
		var numEmployee int
		if _, err := fmt.Scan(&numEmployee); err != nil {
			fmt.Println("Ошибка! Количество работников должно быть целым числом!")
			return
		}

		mintemp, maxtemp := tmin, tmax

		for range numEmployee {
			var op string
			var temp int
			if _, err := fmt.Scan(&op, &temp); err != nil {
				fmt.Println("Ошибка ввода условия!")
				return
			}

			switch op {
			case ">=":
				mintemp = max(mintemp, temp)
			case "<=":
				maxtemp = min(maxtemp, temp)
			default:
				fmt.Println("Неизвестный оператор:", op)
				return
			}

			if mintemp <= maxtemp {
				fmt.Println(mintemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
