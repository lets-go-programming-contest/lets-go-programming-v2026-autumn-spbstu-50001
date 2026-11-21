package settings

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	errConfigPath   = errors.New("provide a path with -config")
	errExtraArgs    = errors.New("unexpected positional arguments")
	errMultipleDocs = errors.New("expected one YAML document")
	errMissingField = errors.New("required configuration field is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ConfigPath(args []string) (string, error) {
	flags := flag.NewFlagSet("task-3", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	path := flags.String("config", "", "path to the YAML configuration")

	if err := flags.Parse(args); err != nil {
		return "", fmt.Errorf("parse arguments: %w", err)
	}

	if flags.NArg() != 0 {
		return "", errExtraArgs
	}

	if strings.TrimSpace(*path) == "" {
		return "", errConfigPath
	}

	return *path, nil
}

func Decode(reader io.Reader) (Config, error) {
	decoder := yaml.NewDecoder(reader)
	decoder.KnownFields(true)

	var config Config

	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("decode YAML: %w", err)
	}

	var extra yaml.Node

	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err != nil {
			return Config{}, fmt.Errorf("read end of YAML: %w", err)
		}

		return Config{}, errMultipleDocs
	}

	return config, config.validate()
}

func (config Config) validate() error {
	if strings.TrimSpace(config.InputFile) == "" {
		return fmt.Errorf("%w: input-file", errMissingField)
	}

	if strings.TrimSpace(config.OutputFile) == "" {
		return fmt.Errorf("%w: output-file", errMissingField)
	}

	return nil
}
