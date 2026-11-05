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

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func LoadXML(path string) ([]Currency, error) {
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

	currencies := make([]Currency, 0, len(rates.Valutes))

	for _, item := range rates.Valutes {
		rawNumCode := strings.TrimSpace(item.NumCode)

		var numCode int

		if rawNumCode != "" {
			if numCode, err = strconv.Atoi(rawNumCode); err != nil {
				return nil, fmt.Errorf("parse num code %q: %w", rawNumCode, err)
			}
		}

		value, err := parseValue(item.Value)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
}

func SortByValueDescending(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func SaveJSON(path string, currencies []Currency) (err error) {
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
