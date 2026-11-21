package main

import (
	"fmt"
	"os"

	"github.com/x1tinter/task-3/internal/fileio"
	"github.com/x1tinter/task-3/internal/rates"
	"github.com/x1tinter/task-3/internal/settings"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		panic(err)
	}
}

func run(args []string) error {
	configPath, err := settings.ConfigPath(args)
	if err != nil {
		return fmt.Errorf("read command line: %w", err)
	}

	config, err := fileio.Read(configPath, settings.Decode)
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	quotes, err := fileio.Read(config.InputFile, rates.Decode)
	if err != nil {
		return fmt.Errorf("load exchange rates: %w", err)
	}

	rates.SortByValue(quotes)

	if err := fileio.WriteJSON(config.OutputFile, quotes); err != nil {
		return fmt.Errorf("save exchange rates: %w", err)
	}

	return nil
}
