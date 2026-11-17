package main

import (
	"errors"
	"flag"

	"github.com/Alehandro-de-bah/task-3/internal/config"
	"github.com/Alehandro-de-bah/task-3/internal/currency"
)

const (
	configFlagName  = "config"
	configFlagUsage = "path to YAML configuration file"
)

var errConfigPathRequired = errors.New("config path is required")

func parseFlags() (string, error) {
	configPath := flag.String(configFlagName, "", configFlagUsage)
	flag.Parse()

	if *configPath == "" {
		return "", errConfigPathRequired
	}

	return *configPath, nil
}

func main() {
	configPath, err := parseFlags()
	if err != nil {
		panic(err)
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.LoadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(currencies)

	if err := currency.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
