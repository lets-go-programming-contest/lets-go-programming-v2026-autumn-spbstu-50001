package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var ErrEmptyConfigField = errors.New("config field cannot be empty")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, fmt.Errorf("input-file is empty: %w", ErrEmptyConfigField)
	}

	if cfg.OutputFile == "" {
		return nil, fmt.Errorf("output-file is empty: %w", ErrEmptyConfigField)
	}

	return &cfg, nil
}
