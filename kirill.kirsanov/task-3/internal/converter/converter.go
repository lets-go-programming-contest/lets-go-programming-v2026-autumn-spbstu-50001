package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/kirillkirsanov2007/task-3/internal/parser"
)

const (
	filePerm = 0o600
	dirPerm  = 0o755
)

func ProcessData(inputFile, outputFile string) error {
	valCurs, err := parser.ReadXML(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read XML: %w", err)
	}

	sortData(&valCurs)

	if err := writeJSON(outputFile, valCurs.Valutes); err != nil {
		return err
	}

	return nil
}

func sortData(valCurs *parser.ValCurs) {
	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})
}

func writeJSON(path string, valutes []parser.Valute) error {
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
