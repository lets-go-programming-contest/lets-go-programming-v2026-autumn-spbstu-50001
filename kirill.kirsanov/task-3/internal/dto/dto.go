package dto

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type FloatValue float64

func (f *FloatValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var stringValue string
	if err := d.DecodeElement(&stringValue, &start); err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	stringValue = strings.ReplaceAll(stringValue, ",", ".")

	val, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float: %w", err)
	}

	*f = FloatValue(val)

	return nil
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatValue `json:"value"     xml:"Value"`
}
