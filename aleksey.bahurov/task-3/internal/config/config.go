package config

import (
	"errors"
)

var (
	errInputFileEmpty  = errors.New("config field 'input-file' is empty")
	errOutputFileEmpty = errors.New("config field 'output-file' is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (c Config) Validate() error {
	if c.InputFile == "" {
		return errInputFileEmpty
	}

	if c.OutputFile == "" {
		return errOutputFileEmpty
	}

	return nil
}
