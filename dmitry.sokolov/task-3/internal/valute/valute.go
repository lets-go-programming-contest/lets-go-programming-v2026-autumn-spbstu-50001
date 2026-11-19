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
	s := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)
	val, err := strconv.ParseFloat(s, 64)
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
	ID        string     `xml:"ID,attr" json:"-"`
	NumCode   string     `xml:"NumCode" json:"num_code"`
	CharCode  string     `xml:"CharCode" json:"char_code"`
	Nominal   int        `xml:"Nominal" json:"-"`
	Name      string     `xml:"Name" json:"-"`
	Value     CommaFloat `xml:"Value" json:"value"`
	VunitRate CommaFloat `xml:"VunitRate" json:"-"`
}

func (vc *ValCurs) SortByValue() {
	slices.SortFunc(vc.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
