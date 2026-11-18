package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	floatBitSize       = 64
	jsonIndent         = "  "
	windows1251Charset = "windows-1251"
)

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

type optionalInt int

func (oi *optionalInt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode int element: %w", err)
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		*oi = 0

		return nil
	}

	val, err := strconv.Atoi(trimmed)
	if err != nil {
		return fmt.Errorf("parse int value %q: %w", raw, err)
	}

	*oi = optionalInt(val)

	return nil
}

type Currency struct {
	NumCode  optionalInt `json:"num_code"  xml:"NumCode"`
	CharCode string      `json:"char_code" xml:"CharCode"`
	Value    commaFloat  `json:"value"     xml:"Value"`
}

type valCurs struct {
	Valutes []Currency `xml:"Valute"`
}

func DecodeXML(r io.Reader) ([]Currency, error) {
	decoder := xml.NewDecoder(r)
	decoder.CharsetReader = xmlCharsetReader

	var rates valCurs
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML file: %w", err)
	}

	return rates.Valutes, nil
}

func SortByValueDescending(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func EncodeJSON(w io.Writer, data any) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}

func xmlCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, windows1251Charset) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}
