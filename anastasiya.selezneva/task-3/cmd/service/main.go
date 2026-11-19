package main

import (
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
		return err
	}

	input, err := storage.Read(cfg.InputFile)
	if err != nil {
		return err
	}

	currencies, err := currency.DecodeXML(input)
	if err != nil {
		return err
	}
	currency.SortByValue(currencies)

	output, err := currency.EncodeJSON(currencies)
	if err != nil {
		return err
	}
	return storage.Write(cfg.OutputFile, output)
}
