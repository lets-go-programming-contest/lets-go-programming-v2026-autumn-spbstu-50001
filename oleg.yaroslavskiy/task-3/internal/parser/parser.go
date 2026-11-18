package parser

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
	filePermission = 0o644
	dirPermission  = 0o755
	jsonIndent     = "    "
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type OutputRate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertXMLToJSON(inputFile, outputFile string) error {
	rawData, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	var curs ValCurs

	decoder := xml.NewDecoder(bytes.NewReader(rawData))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&curs); err != nil {
		return fmt.Errorf("decode xml: %w", err)
	}

	rates := make([]OutputRate, 0, len(curs.Valutes))

	for _, v := range curs.Valutes {
		numCodeStr := strings.TrimSpace(v.NumCode)
		valueStr := strings.ReplaceAll(strings.TrimSpace(v.Value), ",", ".")
		charCode := strings.TrimSpace(v.CharCode)

		if numCodeStr == "" || valueStr == "" || charCode == "" {
			continue
		}

		numCode, err := strconv.Atoi(numCodeStr)
		if err != nil {
			return fmt.Errorf("parse num code %q: %w", v.NumCode, err)
		}

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return fmt.Errorf("parse value %q: %w", v.Value, err)
		}

		rates = append(rates, OutputRate{
			NumCode:  numCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Value > rates[j].Value
	})

	if err := saveJSON(outputFile, rates); err != nil {
		return err
	}

	return nil
}

func saveJSON(path string, rates []OutputRate) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermission); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(rates, "", jsonIndent)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(path, jsonData, filePermission); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
