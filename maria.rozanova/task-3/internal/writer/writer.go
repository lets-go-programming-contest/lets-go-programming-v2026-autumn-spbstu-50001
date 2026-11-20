package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mrzls/task-3/internal/model"
)

const (
	dirPerm    = 0o750
	filePerm   = 0o600
	jsonIndent = "  "
)

func WriteJSON(path string, rates []model.Rate) error {
	if rates == nil {
		rates = []model.Rate{}
	}

	data, err := json.MarshalIndent(rates, "", jsonIndent)
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	data = append(data, '\n')

	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create directory for %q: %w", path, err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write output file %q: %w", path, err)
	}

	return nil
}
