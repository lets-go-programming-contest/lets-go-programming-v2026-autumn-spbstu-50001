package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var departments int
	if _, err := fmt.Fscan(in, &departments); err != nil {
		fmt.Fprintln(os.Stderr, "cannot read number of departments:", err)
		return
	}

	for i := 0; i < departments; i++ {
		var employees int
		if _, err := fmt.Fscan(in, &employees); err != nil {
			fmt.Fprintln(os.Stderr, "cannot read number of employees:", err)
			return
		}

		minTemp, maxTemp := 15, 30
		for j := 0; j < employees; j++ {
			var operator string
			var temperature int
			if _, err := fmt.Fscan(in, &operator, &temperature); err != nil {
				fmt.Fprintln(os.Stderr, "cannot read employee preference:", err)
				return
			}

			switch operator {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			case "==":
				if temperature > minTemp {
					minTemp = temperature
				}
				if temperature < maxTemp {
					maxTemp = temperature
				}
			default:
				fmt.Fprintf(os.Stderr, "unknown operator: %q\n", operator)
				return
			}

			if minTemp > maxTemp {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, minTemp)
			}
		}
	}
}
