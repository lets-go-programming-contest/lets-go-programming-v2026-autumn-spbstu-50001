package converter

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type FloatValue float64

func (f *FloatValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	v = strings.ReplaceAll(v, ",", ".")
	val, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*f = FloatValue(val)

	return nil
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int        `xml:"NumCode" json:"num_code"`
	CharCode string     `xml:"CharCode" json:"char_code"`
	Value    FloatValue `xml:"Value" json:"value"`
}

func ProcessData(inputFile, outputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
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
		return err
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	output := make([]Valute, 0, len(valCurs.Valutes))
	output = append(output, valCurs.Valutes...)

	dir := filepath.Dir(outputFile)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, jsonData, 0o644)
}
