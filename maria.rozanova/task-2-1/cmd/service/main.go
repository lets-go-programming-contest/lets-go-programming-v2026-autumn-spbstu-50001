package main

import "fmt"

const (
	tmin = 15
	tmax = 30
)

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

		mintemp, maxtemp := tmin, tmax

		for range numEmployee {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error: invalid input format")

				return
			}

			switch operator {
			case ">=":
				mintemp = max(mintemp, temp)
			case "<=":
				maxtemp = min(maxtemp, temp)
			default:
				fmt.Println("Unknown operator:", operator)

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
