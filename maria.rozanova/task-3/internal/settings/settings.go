package settings

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errInputFileEmpty  = errors.New("config field 'input-file' is empty")
	errOutputFileEmpty = errors.New("config field 'output-file' is empty")
)

type Settings struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Read(path string) (Settings, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	var cfg Settings
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Settings{}, fmt.Errorf("parse config file %q: %w", path, err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return Settings{}, errInputFileEmpty
	}

	if cfg.OutputFile == "" {
		return Settings{}, errOutputFileEmpty
	}

	return cfg, nil
}
