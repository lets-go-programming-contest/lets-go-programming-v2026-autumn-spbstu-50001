package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatValue `json:"value"     xml:"Value"`
}

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

func ReadXML(path string) (ValCurs, error) {
	var valCurs ValCurs

	data, err := os.ReadFile(path)
	if err != nil {
		return valCurs, fmt.Errorf("failed to read input file: %w", err)
	}

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("failed to decode XML: %w", err)
	}

	return valCurs, nil
}
