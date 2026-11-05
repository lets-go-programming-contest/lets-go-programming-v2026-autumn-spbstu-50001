package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	inputFileField  = "input-file"
	outputFileField = "output-file"
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
		return Config{}, fmt.Errorf("config field %q is empty", inputFileField)
	}

	if cfg.OutputFile == "" {
		return Config{}, fmt.Errorf("config field %q is empty", outputFileField)
	}

	return cfg, nil
}
