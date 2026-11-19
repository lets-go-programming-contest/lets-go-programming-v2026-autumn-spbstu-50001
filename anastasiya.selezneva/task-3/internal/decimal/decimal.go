package decimal

import (
	"strconv"
	"strings"
)

type Float float64

func (f *Float) UnmarshalText(text []byte) error {
	s := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = Float(v)
	return nil
}
