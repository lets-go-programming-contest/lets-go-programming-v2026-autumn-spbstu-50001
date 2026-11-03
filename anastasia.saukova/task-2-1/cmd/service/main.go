package main

import (
	"errors"
	"fmt"
)

const (
	minimalTemperature = 15
	maximumTemperature = 30

	minimalDepartments = 1
	maximumDepartments = 1000

	minimalEmployees = 1
	maximumEmployees = 1000

	noSolution = -1
)

var (
	errOutOfRange       = errors.New("value out of range")
	errUnknownOperation = errors.New("unknown operation")
)

type temperatureState struct {
	low  int
	high int
}

func newTemperatureState() *temperatureState {
	return &temperatureState{
		low:  minimalTemperature,
		high: maximumTemperature,
	}
}

func (s *temperatureState) apply(operation string, temperature int) error {
	switch operation {
	case ">=":
		if temperature > s.low {
			s.low = temperature
		}
	case "<=":
		if temperature < s.high {
			s.high = temperature
		}
	default:
		return fmt.Errorf("%w: %q", errUnknownOperation, operation)
	}

	return nil
}

func (s *temperatureState) optimalTemperature() int {
	if s.low > s.high {
		return noSolution
	}

	return s.low
}

func readCount(name string, minValue, maxValue int) (int, error) {
	var count int
	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to read %s: %w", name, err)
	}

	if count < minValue || count > maxValue {
		return 0, fmt.Errorf("%s: %w (must be from %d to %d, got %d)", name, errOutOfRange, minValue, maxValue, count)
	}

	return count, nil
}

func main() {
	departmentsNumber, err := readCount("number of departments", minimalDepartments, maximumDepartments)
	if err != nil {
		fmt.Println(err)

		return
	}

	for department := range departmentsNumber {
		employeesNumber, err := readCount("number of employees", minimalEmployees, maximumEmployees)
		if err != nil {
			fmt.Printf("department %d: %v\n", department+1, err)

			return
		}

		state := newTemperatureState()

		for employee := range employeesNumber {
			var (
				operation   string
				temperature int
			)

			if _, err := fmt.Scan(&operation, &temperature); err != nil {
				fmt.Printf("department %d, employee %d: failed to read preference: %v\n", department+1, employee+1, err)

				return
			}

			if err := state.apply(operation, temperature); err != nil {
				fmt.Printf("department %d, employee %d: %v\n", department+1, employee+1, err)

				return
			}

			fmt.Println(state.optimalTemperature())
		}
	}
}
