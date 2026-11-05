package climate

import (
	"errors"
)

type Climate struct {
	MaxTemp int
	MinTemp int
}

var (
	ErrIncorrectSign = errors.New("incorrect sign")
	ErrInvalidTemp   = errors.New("invalid temperature")
)

func UpdateClimate(currClimate *Climate, sign string, preferredTemp int) error {
	if currClimate.MinTemp == -1 {
		return ErrInvalidTemp
	}

	switch sign {
	case ">=":
		currClimate.MinTemp = max(currClimate.MinTemp, preferredTemp)
	case "<=":
		currClimate.MaxTemp = min(currClimate.MaxTemp, preferredTemp)
	default:
		return ErrIncorrectSign
	}

	if currClimate.MinTemp > currClimate.MaxTemp {
		return ErrInvalidTemp
	}

	return nil
}
