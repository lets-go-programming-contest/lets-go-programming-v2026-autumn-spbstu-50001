package main

import (
	"fmt"

	"github.com/dmitrys6/task-2-1/internal/climate"
)

func main() {
	var numberOfRooms, employeesInRoom int

	_, err := fmt.Scan(&numberOfRooms)
	if err != nil {
		fmt.Println("Incorrect input:", err)

		return
	}

	if numberOfRooms < 0 {
		fmt.Println("Incorrect input: number of rooms cannot be negative")

		return
	}

	for range numberOfRooms {
		_, err := fmt.Scan(&employeesInRoom)
		if err != nil {
			fmt.Println("Incorrect input:", err)

			return
		}

		if employeesInRoom < 0 {
			fmt.Println("Incorrect input: there cannot be a negative number of employees")

			return
		}

		temperatures, err := readRoomDataAndCalculateClimate(employeesInRoom)
		if err != nil {
			fmt.Println("error while calculating temperature:", err)

			return
		}

		for _, val := range temperatures {
			fmt.Println(val)
		}
	}
}

func readRoomDataAndCalculateClimate(employeesInRoom int) ([]int, error) {
	var (
		sign          string
		preferredTemp int
		signs         = make([]string, employeesInRoom)
		temps         = make([]int, employeesInRoom)
	)

	for index := range employeesInRoom {
		_, err := fmt.Scan(&sign)
		if err != nil {
			return nil, fmt.Errorf("calculateClimateInRooms: %w", err)
		}

		signs[index] = sign

		_, err = fmt.Scan(&preferredTemp)
		if err != nil {
			return nil, fmt.Errorf("calculateClimateInRooms: %w", err)
		}

		temps[index] = preferredTemp
	}

	temperatures, err := climate.CalculateClimateInRoom(signs, temps)
	if err != nil {
		return nil, fmt.Errorf("calculateClimateInRooms: %w", err)
	}

	return temperatures, nil
}
