package main

import "fmt"

func main() {
	var (
		firstOperand, secondOperand int
		operation                   string
	)

	_, error := fmt.Scanln(&firstOperand)
	if error != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, error = fmt.Scanln(&secondOperand)
	if error != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, error = fmt.Scanln(&operation)
	if error != nil {
		fmt.Println("Invalid operation")
		return
	}
}
