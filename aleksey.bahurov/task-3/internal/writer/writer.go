package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	outputDirectoryPermission = 0o755
	jsonIndent                = "  "
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToJSONFile[T any](path string, data T) error {
	if err := os.MkdirAll(filepath.Dir(path), outputDirectoryPermission); err != nil {
		return fmt.Errorf("create output directory for %q: %w", path, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}

	defer func() {
		panicIfErr(file.Close())
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", jsonIndent)

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode JSON to %q: %w", path, err)
	}

	return nil
}
