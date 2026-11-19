package reader

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mrzls/task-3/internal/model"
	"golang.org/x/text/encoding/charmap"
)

const windows1251 = "windows-1251"

var errUnsupportedCharset = errors.New("unsupported XML charset")

func ReadXML(path string) ([]model.Rate, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	var doc struct {
		Rates []model.Rate `xml:"Valute"`
	}

	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode XML file %q: %w", path, err)
	}

	return doc.Rates, nil
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	if !strings.EqualFold(label, windows1251) {
		return nil, fmt.Errorf("%w: %s", errUnsupportedCharset, label)
	}

	return charmap.Windows1251.NewDecoder().Reader(input), nil
}
