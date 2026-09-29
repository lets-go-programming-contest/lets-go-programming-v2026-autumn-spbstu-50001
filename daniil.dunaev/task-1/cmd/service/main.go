package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	if !input.Scan() {
		fmt.Println("Invalid first operand")
		return
	}
	first, err := strconv.Atoi(strings.TrimSpace(input.Text()))
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if !input.Scan() {
		fmt.Println("Invalid second operand")
		return
	}
	second, err := strconv.Atoi(strings.TrimSpace(input.Text()))
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if !input.Scan() {
		fmt.Println("Invalid operation")
		return
	}
	operation := strings.TrimSpace(input.Text())

	switch operation {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
