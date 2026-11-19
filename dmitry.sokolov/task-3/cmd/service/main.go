package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dmitrys6/task-3/internal/config"
	"github.com/dmitrys6/task-3/internal/valute"
)

const dirPerm = 0o755

func main() {
	configPath := flag.String("config", "config.yml", "config file path")
	flag.Parse()

	if strings.TrimSpace(*configPath) == "" {
		panic("no config file path")
	}

	inOutFiles, err := openAndReadIOFiles(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := openAndReadValCurs(inOutFiles.InFile)
	if err != nil {
		panic(err)
	}

	valCurs.SortByValue()

	err = writeValCurs(inOutFiles.OutFile, &valCurs)
	if err != nil {
		panic(err)
	}
}

func openAndReadIOFiles(filename string) (config.IOFiles, error) {
	var inOutFiles config.IOFiles

	data, err := os.Open(filename)
	if err != nil {
		return inOutFiles, fmt.Errorf("openAndReadIOFiles: %w", err)
	}

	inOutFiles, readErr := config.GetIOFiles(data)

	closeErr := data.Close()
	if closeErr != nil && readErr == nil {
		return inOutFiles, fmt.Errorf("openAndReadIOFiles: close file: %w", closeErr)
	}

	if readErr != nil {
		return inOutFiles, fmt.Errorf("openAndReadIOFiles: %w", readErr)
	}

	return inOutFiles, nil
}

func openAndReadValCurs(filename string) (valute.ValCurs, error) {
	var valCurs valute.ValCurs

	inFile, err := os.Open(filename)
	if err != nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: %w", err)
	}

	valCurs, readErr := config.ReadValCurs(inFile)

	closeErr := inFile.Close()
	if closeErr != nil && readErr == nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: close file: %w", closeErr)
	}

	if readErr != nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: %w", readErr)
	}

	return valCurs, nil
}

func writeValCurs(filename string, valCurs *valute.ValCurs) error {
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, dirPerm)
	if err != nil {
		return fmt.Errorf("writeValCurs: %w", err)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("writeValCurs: %w", err)
	}

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "    ")

	encodeErr := encoder.Encode(valCurs.Valutes)

	closeErr := outFile.Close()
	if closeErr != nil && encodeErr == nil {
		return fmt.Errorf("writeValCurs: close file: %w", closeErr)
	}

	if encodeErr != nil {
		return fmt.Errorf("writeValCurs: %w", encodeErr)
	}

	return nil
}
