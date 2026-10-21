package main

import "fmt"

func main() {
	var (
		numberFirst, numberSecond int
		operation                 string
	)

	if _, err := fmt.Scan(&numberFirst); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&numberSecond); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(numberFirst + numberSecond)
	case "-":
		fmt.Println(numberFirst - numberSecond)
	case "*":
		fmt.Println(numberFirst * numberSecond)
	case "/":
		if numberSecond == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(numberFirst / numberSecond)
	default:
		fmt.Println("Invalid operation")
	}
}
