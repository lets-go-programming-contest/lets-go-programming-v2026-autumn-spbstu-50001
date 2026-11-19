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

func openAndReadIOFiles(filename string) (inOutFiles config.IOFiles, err error) {
	data, err := os.Open(filename)
	if err != nil {
		return inOutFiles, fmt.Errorf("openAndReadIOFiles: %w", err)
	}

	defer func() {
		if closeErr := data.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("openAndReadIOFiles: close file: %w", closeErr)
		}
	}()

	inOutFiles, err = config.GetIOFiles(data)
	if err != nil {
		return inOutFiles, fmt.Errorf("openAndReadIOFiles: %w", err)
	}

	return inOutFiles, nil
}

func openAndReadValCurs(filename string) (valCurs valute.ValCurs, err error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: %w", err)
	}

	defer func() {
		if closeErr := inFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("openAndReadValCurs: close file: %w", closeErr)
		}
	}()

	valCurs, err = config.ReadValCurs(inFile)
	if err != nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: %w", err)
	}

	return valCurs, nil
}

func writeValCurs(filename string, valCurs *valute.ValCurs) (err error) {
	dir := filepath.Dir(filename)

	err = os.MkdirAll(dir, dirPerm)
	if err != nil {
		return fmt.Errorf("writeValCurs: %w", err)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("writeValCurs: %w", err)
	}

	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("writeValCurs: close file: %w", closeErr)
		}
	}()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(valCurs.Valutes)
	if err != nil {
		return fmt.Errorf("writeValCurs: %w", err)
	}

	return err
}
