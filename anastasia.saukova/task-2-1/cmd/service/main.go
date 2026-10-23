package main

import "fmt"

const (
	minimalTemperature = 15
	maximumTemperature = 30

	minimalDepartments = 1
	maximumDepartments = 1000

	minimalEmployees = 1
	maximumEmployees = 1000

	noSolution = -1
)

func readCount(name string, minValue, maxValue int) (int, error) {
	var count int
	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to read %s: %w", name, err)
	}

	if count < minValue || count > maxValue {
		return 0, fmt.Errorf("%s must be from %d to %d, got %d", name, minValue, maxValue, count)
	}

	return count, nil
}

func narrowRange(low, high int, operation string, temperature int) (int, int, error) {
	switch operation {
	case ">=":
		if temperature > low {
			low = temperature
		}
	case "<=":
		if temperature < high {
			high = temperature
		}
	default:
		return 0, 0, fmt.Errorf("unknown operation %q", operation)
	}

	return low, high, nil
}

func main() {
	departmentsNumber, err := readCount("number of departments", minimalDepartments, maximumDepartments)
	if err != nil {
		fmt.Println(err)

		return
	}

	for department := 1; department <= departmentsNumber; department++ {
		employeesNumber, err := readCount("number of employees", minimalEmployees, maximumEmployees)
		if err != nil {
			fmt.Printf("department %d: %v\n", department, err)

			return
		}

		low, high := minimalTemperature, maximumTemperature

		for employee := 1; employee <= employeesNumber; employee++ {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				fmt.Printf("department %d, employee %d: failed to read preference: %v\n", department, employee, err)

				return
			}

			low, high, err = narrowRange(low, high, operation, temperature)
			if err != nil {
				fmt.Printf("department %d, employee %d: %v\n", department, employee, err)

				return
			}

			if low > high {
				fmt.Println(noSolution)
			} else {
				fmt.Println(low)
			}
		}
	}
}