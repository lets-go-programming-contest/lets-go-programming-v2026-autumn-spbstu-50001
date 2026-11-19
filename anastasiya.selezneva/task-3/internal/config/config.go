package config

import (
	"errors"
	"flag"

	"gopkg.in/yaml.v3"

	"github.com/seleznao/task-3/internal/storage"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Parse(args []string) (Config, error) {
	flags := flag.NewFlagSet("service", flag.ContinueOnError)
	path := flags.String("config", "", "path to the YAML config file")
	if err := flags.Parse(args); err != nil {
		return Config{}, err
	}
	if *path == "" {
		return Config{}, errors.New("flag -config is required")
	}

	data, err := storage.Read(*path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Config{}, errors.New("input-file and output-file must be set")
	}
	return cfg, nil
}
