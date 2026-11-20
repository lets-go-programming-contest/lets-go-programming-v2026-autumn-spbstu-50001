package model

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const floatBitSize = 64

type Decimal float64

func (d *Decimal) UnmarshalText(text []byte) error {
	normalized := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)

	value, err := strconv.ParseFloat(normalized, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse decimal %q: %w", text, err)
	}

	*d = Decimal(value)

	return nil
}

type Rate struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value"     xml:"Value"`
}

func SortByValueDesc(rates []Rate) {
	sort.SliceStable(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})
}
