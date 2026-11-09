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

type config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return config{}, fmt.Errorf("read config file %q: %w", path, err)
	}

	var cfg config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return config{}, fmt.Errorf("parse config file %q: %w", path, err)
	}

	if cfg.InputFile == "" {
		return config{}, errInputFileEmpty
	}

	if cfg.OutputFile == "" {
		return config{}, errOutputFileEmpty
	}

	return cfg, nil
}
