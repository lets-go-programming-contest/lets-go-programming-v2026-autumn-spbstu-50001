package main

import (
	"errors"
	"flag"
	"io"

	"github.com/Alehandro-de-bah/task-3/internal/config"
	"github.com/Alehandro-de-bah/task-3/internal/currency"
	"github.com/Alehandro-de-bah/task-3/internal/storage"
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

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := storage.LoadCurrenciesFromFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(currencies)

	if err := storage.SaveToFile(cfg.OutputFile, func(w io.Writer) error {
		return currency.EncodeJSON(w, currencies)
	}); err != nil {
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
