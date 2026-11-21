package fileio

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const directoryMode = 0o755

func Read[T any](path string, decode func(io.Reader) (T, error)) (T, error) {
	var empty T

	file, err := os.Open(path)
	if err != nil {
		return empty, fmt.Errorf("open %q: %w", path, err)
	}

	defer closeFile(file)

	data, err := decode(file)
	if err != nil {
		return empty, fmt.Errorf("read %q: %w", path, err)
	}

	return data, nil
}

func WriteJSON(path string, data any) error {
	if err := os.MkdirAll(filepath.Dir(path), directoryMode); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %q: %w", path, err)
	}

	defer closeFile(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("write JSON to %q: %w", path, err)
	}

	return nil
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(fmt.Errorf("close %q: %w", file.Name(), err))
	}
}
