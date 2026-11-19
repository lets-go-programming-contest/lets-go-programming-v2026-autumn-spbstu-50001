package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	directoryPermissions = 0o755
	filePermissions      = 0o600
)

func Read(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}

	return data, nil
}

func Write(path string, data []byte) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, directoryPermissions); err != nil {
		return fmt.Errorf("create directory %q: %w", dir, err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write file %q: %w", path, err)
	}

	return nil
}
