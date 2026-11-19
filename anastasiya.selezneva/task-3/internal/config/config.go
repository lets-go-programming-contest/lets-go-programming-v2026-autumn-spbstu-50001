package config

import (
	"errors"
	"flag"
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/seleznao/task-3/internal/storage"
)

var (
	errConfigRequired = errors.New("flag -config is required")
	errEmptyPaths     = errors.New("input-file and output-file must be set")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Parse(args []string) (Config, error) {
	flags := flag.NewFlagSet("service", flag.ContinueOnError)

	path := flags.String("config", "", "path to the YAML config file")

	if err := flags.Parse(args); err != nil {
		return Config{}, fmt.Errorf("parse flags: %w", err)
	}

	if *path == "" {
		return Config{}, errConfigRequired
	}

	data, err := storage.Read(*path)
	if err != nil {
		return Config{}, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse YAML config: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, errEmptyPaths
	}

	return cfg, nil
}
