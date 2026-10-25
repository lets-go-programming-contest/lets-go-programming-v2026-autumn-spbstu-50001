package main

import (
	"fmt"

	"github.com/kirillkirsanov2007/task-2-1/internal/climate"
)

const (
	initialMinTemp = 15
	initialMaxTemp = 30
)

func main() {
	var numberOfDepartments int
	if _, err := fmt.Scan(&numberOfDepartments); err != nil {
		fmt.Println("Invalid number of departments:", err)
		return
	}

	for i := 0; i < numberOfDepartments; i++ {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Println("Invalid number of employees:", err)
			return
		}

		tempRange := climate.NewTemperatureRange(initialMinTemp, initialMaxTemp)

		for j := 0; j < employeesCount; j++ {
			var (
				sign       string
				limitValue int
			)

			if _, err := fmt.Scan(&sign, &limitValue); err != nil {
				fmt.Println("Invalid temperature condition:", err)
				return
			}

			if sign == ">=" {
				tempRange.SetMinLimit(limitValue)
			} else if sign == "<=" {
				tempRange.SetMaxLimit(limitValue)
			} else {
				fmt.Println("Invalid operation sign")
				return
			}

			optimal, err := tempRange.GetOptimalTemperature()
			if err != nil {
				fmt.Println(-1)
			} else {
				fmt.Println(optimal)
			}
		}
	}
}
