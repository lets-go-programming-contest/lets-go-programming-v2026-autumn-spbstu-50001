package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mrzls/task-3/internal/model"
	"github.com/mrzls/task-3/internal/reader"
	"github.com/mrzls/task-3/internal/settings"
	"github.com/mrzls/task-3/internal/writer"
)

var errNoConfigFlag = errors.New("flag -config is required")

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	configPath := flag.String("config", "", "path to the YAML config file")
	flag.Parse()

	if *configPath == "" {
		return errNoConfigFlag
	}

	cfg, err := settings.Read(*configPath)
	if err != nil {
		return fmt.Errorf("read settings: %w", err)
	}

	rates, err := reader.ReadXML(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("read xml: %w", err)
	}

	model.SortByValueDesc(rates)

	return fmt.Errorf("write json: %w", writer.WriteJSON(cfg.OutputFile, rates))
}
