package main

import (
	"fmt"

	"github.com/Alehandro-de-bah/task-2-1/internal/climate"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var departmentsNumber int
	if _, err := fmt.Scan(&departmentsNumber); err != nil {
		fmt.Println("Invalid number of departments:", err.Error())

		return
	}

	for range departmentsNumber {
		var employeesNumber int
		if _, err := fmt.Scan(&employeesNumber); err != nil {
			fmt.Println("Invalid number of employees:", err.Error())

			return
		}

		currentRange, err := climate.NewTemperatureRange(minTemperature, maxTemperature)
		if err != nil {
			fmt.Println("Error:", err.Error())

			return
		}

		for range employeesNumber {
			var (
				comparison  string
				temperature int
			)

			if _, err := fmt.Scan(&comparison, &temperature); err != nil {
				fmt.Println("Invalid temperature:", err.Error())

				return
			}

			constraint, err := climate.NewConstraint(comparison, temperature)
			if err != nil {
				fmt.Println("Error:", err.Error())

				return
			}

			currentRange.ApplyConstraint(constraint)

			optimalTemperature, err := currentRange.GetOptimalTemperature()
			if err != nil {
				fmt.Println(-1)

				continue
			}

			fmt.Println(optimalTemperature)
		}
	}
}
