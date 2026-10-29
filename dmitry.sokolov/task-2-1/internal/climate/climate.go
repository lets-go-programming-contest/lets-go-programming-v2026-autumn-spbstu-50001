package climate

import (
	"errors"
	"fmt"
)

const (
	defaultMaxTemp = 30
	defaultMinTemp = 15
)

type climate struct {
	maxTemp      int
	minTemp      int
	temperatures []int
}

var (
	ErrIncorrectSign = errors.New("incorrect sign")
	ErrInvalidInput  = errors.New("signs and vals length differ")
)

func CalculateClimateInRoom(signs []string, vals []int) ([]int, error) {
	if len(signs) != len(vals) {
		return nil, ErrInvalidInput
	}

	climateInRoom := climate{
		maxTemp:      defaultMaxTemp,
		minTemp:      defaultMinTemp,
		temperatures: make([]int, len(vals)),
	}

	for index := range signs {
		err := updateClimate(&climateInRoom, index, vals[index], signs[index])
		if err != nil {
			return nil, fmt.Errorf("calculateClimateInRoom: %w", err)
		}
	}

	return climateInRoom.temperatures, nil
}

func updateClimate(currClimate *climate, index int, preferredTemp int, sign string) error {
	switch sign {
	case ">=":
		currClimate.minTemp = max(currClimate.minTemp, preferredTemp)
	case "<=":
		currClimate.maxTemp = min(currClimate.maxTemp, preferredTemp)
	default:
		return ErrIncorrectSign
	}

	if currClimate.minTemp > currClimate.maxTemp {
		currClimate.temperatures[index] = -1
	} else {
		currClimate.temperatures[index] = currClimate.minTemp
	}

	return nil
}
