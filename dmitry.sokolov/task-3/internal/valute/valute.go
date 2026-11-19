package valute

import (
	"cmp"
	"encoding/xml"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type CommaFloat float64

func (f *CommaFloat) UnmarshalText(text []byte) error {
	trimmed := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)
	if trimmed == "" {
		*f = 0

		return nil
	}

	val, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return fmt.Errorf("UnmarshalText CommaFloat: %w", err)
	}

	*f = CommaFloat(val)

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID        string     `json:"-"         xml:"ID,attr"`
	NumCode   int        `json:"num_code"  xml:"NumCode"`
	CharCode  string     `json:"char_code" xml:"CharCode"`
	Nominal   int        `json:"-"         xml:"Nominal"`
	Name      string     `json:"-"         xml:"Name"`
	Value     CommaFloat `json:"value"     xml:"Value"`
	VunitRate CommaFloat `json:"-"         xml:"VunitRate"`
}

func (valCurs *ValCurs) SortByValue() {
	slices.SortFunc(valCurs.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
