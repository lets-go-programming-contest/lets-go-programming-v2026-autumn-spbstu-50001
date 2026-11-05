package main

import (
	"flag"

	"github.com/Alehandro-de-bah/task-3/internal/config"
	"github.com/Alehandro-de-bah/task-3/internal/currency"
)

const (
	configFlagName  = "config"
	configFlagUsage = "path to YAML configuration file"
)

func main() {
	configPath := flag.String(configFlagName, "", configFlagUsage)
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	config, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.LoadXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDescending(currencies)

	if err := currency.SaveJSON(config.OutputFile, currencies); err != nil {
		panic(err)
	}
}
