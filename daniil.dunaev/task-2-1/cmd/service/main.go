package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

var (
	errUnknownOperator = errors.New("unknown operator")
	errNoTemperature   = errors.New("no suitable temperature")
)

type temperatureRange struct {
	minTemp int
	maxTemp int
}

func newTemperatureRange(minTemp, maxTemp int) temperatureRange {
	return temperatureRange{minTemp: minTemp, maxTemp: maxTemp}
}

func (bounds *temperatureRange) add(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature > bounds.minTemp {
			bounds.minTemp = temperature
		}
	case "<=":
		if temperature < bounds.maxTemp {
			bounds.maxTemp = temperature
		}
	default:
		return fmt.Errorf("%w: %q", errUnknownOperator, operator)
	}

	return nil
}

func (bounds *temperatureRange) lowest() (int, error) {
	if bounds.minTemp > bounds.maxTemp {
		return 0, errNoTemperature
	}

	return bounds.minTemp, nil
}

func processDepartment() error {
	var employees int
	if _, err := fmt.Scan(&employees); err != nil {
		return fmt.Errorf("cannot read number of employees: %w", err)
	}

	bounds := newTemperatureRange(minTemperature, maxTemperature)

	for range employees {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Scan(&operator, &temperature); err != nil {
			return fmt.Errorf("cannot read employee preference: %w", err)
		}

		if err := bounds.add(operator, temperature); err != nil {
			return err
		}

		suitableTemp, err := bounds.lowest()
		if err != nil {
			fmt.Println(-1)

			continue
		}

		fmt.Println(suitableTemp)
	}

	return nil
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "cannot read number of departments:", err)

		return
	}

	for range departments {
		if err := processDepartment(); err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}
	}
}
