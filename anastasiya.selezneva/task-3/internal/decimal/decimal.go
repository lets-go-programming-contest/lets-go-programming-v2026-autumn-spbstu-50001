package decimal

import (
	"fmt"
	"strconv"
	"strings"
)

type Float float64

func (f *Float) UnmarshalText(text []byte) error {
	s := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parse decimal value %q: %w", s, err)
	}

	*f = Float(v)

	return nil
}
