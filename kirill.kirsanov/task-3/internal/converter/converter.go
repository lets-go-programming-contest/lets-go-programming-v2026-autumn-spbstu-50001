package converter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"golang.org/x/text/encoding/charmap"
	"github.com/kirillkirsanov2007/task-3/internal/dto"
)

const (
	filePerm = 0o600
	dirPerm  = 0o755
)

func ProcessData(inputFile, outputFile string) error {
	valCurs, err := readXML(inputFile)
	if err != nil {
		return err
	}

	sortData(&valCurs)

	if err := writeJSON(outputFile, valCurs.Valutes); err != nil {
		return err
	}

	return nil
}

func readXML(path string) (dto.ValCurs, error) {
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

func sortData(valCurs *dto.ValCurs) {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})
}

func writeJSON(path string, valutes []dto.Valute) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	jsonData, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(path, jsonData, filePerm); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}
