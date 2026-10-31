package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer func() {
		_ = writer.Flush()
	}()

	var departmentsCount int
	if _, err := fmt.Fscan(reader, &departmentsCount); err != nil {
		return
	}

	for range departmentsCount {
		if err := processDepartment(reader, writer); err != nil {
			return
		}
	}
}

func processDepartment(reader *bufio.Reader, writer *bufio.Writer) error {
	var employeesCount int
	if _, err := fmt.Fscan(reader, &employeesCount); err != nil {
		return fmt.Errorf("failed to scan employees count: %w", err)
	}

	minTemp, maxTemp := 15, 30
	isFailed := false

	for range employeesCount {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Fscan(reader, &operator, &temperature); err != nil {
			return fmt.Errorf("failed to scan employee preference: %w", err)
		}

		if !isFailed {
			switch operator {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			}

			if minTemp > maxTemp {
				isFailed = true
			}
		}

		if isFailed {
			_, _ = fmt.Fprintln(writer, -1)
		} else {
			_, _ = fmt.Fprintln(writer, minTemp)
		}
	}

	return nil
}
