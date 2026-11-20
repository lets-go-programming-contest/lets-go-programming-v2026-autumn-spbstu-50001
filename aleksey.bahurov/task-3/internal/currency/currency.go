package currency

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const floatBitSize = 64

type commaFloat float64

func (cf *commaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode float element: %w", err)
	}

	normalized := strings.ReplaceAll(strings.TrimSpace(raw), ",", ".")

	val, err := strconv.ParseFloat(normalized, floatBitSize)
	if err != nil {
		return fmt.Errorf("parse float value %q: %w", raw, err)
	}

	*cf = commaFloat(val)

	return nil
}

type Currency struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    commaFloat `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Valutes []Currency `xml:"Valute"`
}

func SortByValueDescending(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
