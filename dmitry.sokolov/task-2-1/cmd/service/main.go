package main

import (
	"fmt"

	"github.com/dmitrys6/task-2-1/internal/climate"
)

func main() {
	var numberOfRooms, employeesInRoom int

	_, err := fmt.Scan(&numberOfRooms)
	if err != nil {
		fmt.Println("Incorrect input")

		return
	}

	for range numberOfRooms {
		_, err = fmt.Scan(&employeesInRoom)
		if err != nil {
			fmt.Println("Incorrect input")

			return
		}

		temperature, err := climate.ProcessRoom(employeesInRoom)

		if err != nil {
			fmt.Println("Incorrect input")

			return
		}

		fmt.Println(temperature)
	}
}
