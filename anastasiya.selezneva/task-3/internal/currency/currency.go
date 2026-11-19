package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"

	"github.com/seleznao/task-3/internal/decimal"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    decimal.Float `json:"value"     xml:"Value"`
}

type valCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func DecodeXML(data []byte) ([]Currency, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return curs.Currencies, nil
}

func SortByValue(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func EncodeJSON(currencies []Currency) ([]byte, error) {
	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("encode JSON: %w", err)
	}

	return data, nil
}
