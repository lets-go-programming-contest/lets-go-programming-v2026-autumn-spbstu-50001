package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"sort"

	"golang.org/x/net/html/charset"

	"github.com/seleznao/task-3/internal/decimal"
)

type Currency struct {
	NumCode  int           `xml:"NumCode" json:"num_code"`
	CharCode string        `xml:"CharCode" json:"char_code"`
	Value    decimal.Float `xml:"Value" json:"value"`
}

type valCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func DecodeXML(data []byte) ([]Currency, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var curs valCurs
	if err := decoder.Decode(&curs); err != nil {
		return nil, err
	}
	return curs.Currencies, nil
}

func SortByValue(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func EncodeJSON(currencies []Currency) ([]byte, error) {
	return json.MarshalIndent(currencies, "", "    ")
}
