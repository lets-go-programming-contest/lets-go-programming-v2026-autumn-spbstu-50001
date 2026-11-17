package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	floatBitSize              = 64
	outputDirectoryPermission = 0o755
	jsonIndent                = "  "
	windows1251Charset        = "windows-1251"
)

type commaFloat float64

func (cf *commaFloat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
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
		return err
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

type currency struct {
	NumCode  optionalInt `xml:"NumCode" json:"num_code"`
	CharCode string      `xml:"CharCode" json:"char_code"`
	Value    commaFloat  `xml:"Value" json:"value"`
}

type valCurs struct {
	Valutes []currency `xml:"Valute"`
}

func LoadXML(path string) ([]currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = xmlCharsetReader

	var rates valCurs
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML file %q: %w", path, err)
	}

	return rates.Valutes, nil
}

func SortByValueDescending(currencies []currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func SaveJSON(path string, currencies []currency) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), outputDirectoryPermission); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close output file %q: %w", path, closeErr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode JSON to %q: %w", path, err)
	}

	return nil
}

func parseValue(strValue string) (float64, error) {
	normalizedValue := strings.ReplaceAll(strings.TrimSpace(strValue), ",", ".")

	value, err := strconv.ParseFloat(normalizedValue, floatBitSize)
	if err != nil {
		return 0, fmt.Errorf("parse value %q: %w", strValue, err)
	}

	return value, nil
}

func xmlCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, windows1251Charset) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}
