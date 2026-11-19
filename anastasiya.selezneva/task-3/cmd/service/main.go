package main

import (
	"fmt"
	"os"

	"github.com/seleznao/task-3/internal/config"
	"github.com/seleznao/task-3/internal/currency"
	"github.com/seleznao/task-3/internal/storage"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.Parse(os.Args[1:])
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	input, err := storage.Read(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	currencies, err := currency.DecodeXML(input)
	if err != nil {
		return fmt.Errorf("decode XML: %w", err)
	}

	currency.SortByValue(currencies)

	output, err := currency.EncodeJSON(currencies)
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := storage.Write(cfg.OutputFile, output); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
