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
		var employeesCount int
		if _, err := fmt.Fscan(reader, &employeesCount); err != nil {
			return
		}

		minTemp, maxTemp := 15, 30
		isFailed := false

		for range employeesCount {
			var operator string
			var temperature int

			if _, err := fmt.Fscan(reader, &operator, &temperature); err != nil {
				return
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
	}
}
