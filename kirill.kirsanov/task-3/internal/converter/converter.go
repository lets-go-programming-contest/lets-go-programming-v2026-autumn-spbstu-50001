package converter

import (
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
	filePerm = 0o600
	dirPerm  = 0o755
)

type FloatValue float64

func (f *FloatValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return fmt.Errorf("failed to decode XML element: %w", err)
	}

	v = strings.ReplaceAll(v, ",", ".")

	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return fmt.Errorf("failed to parse float: %w", err)
	}

	*f = FloatValue(val)

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    FloatValue `json:"value"     xml:"Value"`
}

func ProcessData(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	var valCurs ValCurs

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return fmt.Errorf("failed to decode XML: %w", err)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	output := make([]Valute, 0, len(valCurs.Valutes))
	output = append(output, valCurs.Valutes...)

	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(outputFile, jsonData, filePerm); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
