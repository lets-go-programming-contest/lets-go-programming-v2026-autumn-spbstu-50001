package storage

import (
	"os"
	"path/filepath"
)

func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func Write(path string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
