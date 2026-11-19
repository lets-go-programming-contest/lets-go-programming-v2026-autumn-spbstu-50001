package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInputFile  = errors.New("config: 'input-file' cannot be empty")
	ErrEmptyOutputFile = errors.New("config: 'output-file' cannot be empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string) (*Config, error) {
	rawData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(rawData, &cfg); err != nil {
		return nil, fmt.Errorf("parse yaml: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, ErrEmptyInputFile
	}

	if cfg.OutputFile == "" {
		return nil, ErrEmptyOutputFile
	}

	return &cfg, nil
}