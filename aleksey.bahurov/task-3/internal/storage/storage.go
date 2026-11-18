package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Alehandro-de-bah/task-3/internal/currency"
)

const outputDirectoryPermission = 0o755

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(fmt.Errorf("close file %q: %w", file.Name(), err))
	}
}

func LoadCurrenciesFromFile(path string) ([]currency.Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file %q: %w", path, err)
	}

	defer closeFile(file)

	currencies, err := currency.DecodeXML(file)
	if err != nil {
		return nil, fmt.Errorf("decode currencies from %q: %w", path, err)
	}

	return currencies, nil
}

func SaveToFile(path string, writeFunction func(io.Writer) error) error {
	if err := os.MkdirAll(filepath.Dir(path), outputDirectoryPermission); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	defer closeFile(file)

	return writeFunction(file)
}
