package rates

import (
	"cmp"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"golang.org/x/text/encoding/htmlindex"
)

var (
	errAfterRoot       = errors.New("unexpected data after ValCurs")
	errIncompleteQuote = errors.New("missing Value in Valute")
)

type Quote struct {
	Number int      `json:"num_code"  xml:"NumCode"`
	Code   string   `json:"char_code" xml:"CharCode"`
	Value  *decimal `json:"value"     xml:"Value"`
}

type dailyRates struct {
	XMLName xml.Name `xml:"ValCurs"`
	Quotes  []Quote  `xml:"Valute"`
}

func (quote *Quote) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	type fields Quote

	if err := decoder.DecodeElement((*fields)(quote), &start); err != nil {
		return fmt.Errorf("decode Valute: %w", err)
	}

	if quote.Value == nil {
		return errIncompleteQuote
	}

	return nil
}

func Decode(reader io.Reader) ([]Quote, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charsetReader

	var document dailyRates

	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode ValCurs: %w", err)
	}

	if err := checkEnd(decoder); err != nil {
		return nil, err
	}

	if document.Quotes == nil {
		return []Quote{}, nil
	}

	return document.Quotes, nil
}

func SortByValue(quotes []Quote) {
	slices.SortStableFunc(quotes, func(left, right Quote) int {
		return cmp.Compare(*right.Value, *left.Value)
	})
}

func charsetReader(label string, reader io.Reader) (io.Reader, error) {
	encoding, err := htmlindex.Get(label)
	if err != nil {
		return nil, fmt.Errorf("select XML encoding %q: %w", label, err)
	}

	return encoding.NewDecoder().Reader(reader), nil
}

func checkEnd(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("read end of XML: %w", err)
		}

		switch tail := token.(type) {
		case xml.Comment, xml.ProcInst:
			continue
		case xml.CharData:
			if strings.TrimSpace(string(tail)) == "" {
				continue
			}
		}

		return errAfterRoot
	}
}
