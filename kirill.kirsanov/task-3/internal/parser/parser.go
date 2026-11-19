package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"github.com/kirillkirsanov2007/task-3/internal/dto"
)

func ReadXML(path string) (dto.ValCurs, error) {
	var valCurs dto.ValCurs

	data, err := os.ReadFile(path)
	if err != nil {
		return valCurs, fmt.Errorf("failed to read input file: %w", err)
	}

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return input, nil
	}

	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("failed to decode XML: %w", err)
	}

	return valCurs, nil
}
