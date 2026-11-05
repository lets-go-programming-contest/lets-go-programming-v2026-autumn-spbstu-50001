package main

import (
	"errors"
	"fmt"

	"github.com/dmitrys6/task-2-1/internal/climate"
)

const (
	maxTemp = 30
	minTemp = 15
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

		err = readRoomDataAndCalculateClimate(employeesInRoom)
		if err != nil {
			fmt.Println("error while calculating temperature:", err)

			return
		}
	}
}

func readRoomDataAndCalculateClimate(employeesInRoom int) error {
	var (
		sign          string
		preferredTemp int
		climateInRoom = climate.Climate{
			MaxTemp: maxTemp,
			MinTemp: minTemp,
		}
	)

	for range employeesInRoom {
		_, err := fmt.Scan(&sign)
		if err != nil {
			return fmt.Errorf("calculateClimateInRooms: %w", err)
		}

		_, err = fmt.Scan(&preferredTemp)
		if err != nil {
			return fmt.Errorf("calculateClimateInRooms: %w", err)
		}

		err = climate.UpdateClimate(&climateInRoom, sign, preferredTemp)
		if errors.Is(err, climate.ErrInvalidTemp) {
			climateInRoom.MinTemp = -1
		} else if err != nil {
			return fmt.Errorf("calculateClimateInRooms: %w", err)
		}

		fmt.Println(climateInRoom.MinTemp)
	}

	return nil
}
