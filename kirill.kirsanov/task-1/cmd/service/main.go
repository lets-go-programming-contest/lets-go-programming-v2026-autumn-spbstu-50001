package main

import "fmt"

func main() {
	var (
		firstOperand  int
		secondOperand int
		operation     string
	)

	_, err := fmt.Scanln(&firstOperand)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondOperand)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
			return
		}
	case "*":
		fmt.Println(firstOperand * secondOperand)

	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	default:
		fmt.Println("Invalid operation")
	}
}
