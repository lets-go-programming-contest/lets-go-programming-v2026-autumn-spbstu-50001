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

	defer data.Close()

	inOutFiles, err = config.GetIOFiles(data)

	return inOutFiles, err
}

func openAndReadValCurs(filename string) (valute.ValCurs, error) {
	var valCurs valute.ValCurs

	inFile, err := os.Open(filename)
	if err != nil {
		return valCurs, fmt.Errorf("openAndReadValCurs: %w", err)
	}

	defer inFile.Close()

	valCurs, err = config.ReadValCurs(inFile)

	return valCurs, err
}

func writeValCurs(filename string, valCurs *valute.ValCurs) error {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0o755)

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
