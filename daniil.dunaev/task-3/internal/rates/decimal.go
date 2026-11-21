package rates

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var errInvalidDecimal = errors.New("invalid decimal number")

type decimal float64

func (number *decimal) UnmarshalText(data []byte) error {
	source := strings.TrimSpace(string(data))

	value, err := strconv.ParseFloat(strings.ReplaceAll(source, ",", "."), 64)
	if err != nil {
		return fmt.Errorf("parse decimal %q: %w", source, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("%w: %q", errInvalidDecimal, source)
	}

	*number = decimal(value)

	return nil
}
