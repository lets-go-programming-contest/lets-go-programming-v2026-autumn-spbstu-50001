package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	errInputFileEmpty  = errors.New("config field 'input-file' is empty")
	errOutputFileEmpty = errors.New("config field 'output-file' is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config file %q: %w", path, err)
	}

	if cfg.InputFile == "" {
		return Config{}, errInputFileEmpty
	}

	if cfg.OutputFile == "" {
		return Config{}, errOutputFileEmpty
	}

	return cfg, nil
}
