package main

import (
	"errors"
	"flag"

	"github.com/Alehandro-de-bah/task-3/internal/config"
	"github.com/Alehandro-de-bah/task-3/internal/currency"
	"github.com/Alehandro-de-bah/task-3/internal/reader"
	"github.com/Alehandro-de-bah/task-3/internal/writer"
)

const (
	configFlagName  = "config"
	configFlagUsage = "path to YAML configuration file"
)

var errConfigPathRequired = errors.New("config path is required")

func main() {
	configPath, err := parseFlags()
	if err != nil {
		panic(err)
	}

	cfg, err := reader.FromYAMLFile[config.Config](configPath)
	if err != nil {
		panic(err)
	}

	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	rates, err := reader.FromXMLFile[currency.ValCurs](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(rates.Valutes)

	if err := writer.ToJSONFile(cfg.OutputFile, rates.Valutes); err != nil {
		panic(err)
	}
}

func parseFlags() (string, error) {
	configPath := flag.String(configFlagName, "", configFlagUsage)
	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathRequired
	}

	return *configPath, nil
}
