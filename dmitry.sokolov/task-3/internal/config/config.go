package config

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/dmitrys6/task-3/internal/valute"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyIOFiles   = errors.New("input-file and output-file must not be empty")
	ErrEmptyValueList = errors.New("empty Valute list")
	ErrInvalidFields  = errors.New("invalid Valute fields")
)

type IOFiles struct {
	InFile  string `yaml:"input-file"`
	OutFile string `yaml:"output-file"`
}

func GetIOFiles(configFile io.Reader) (IOFiles, error) {
	var inOutFiles IOFiles

	dec := yaml.NewDecoder(configFile)
	dec.KnownFields(true)

	err := dec.Decode(&inOutFiles)
	if err != nil {
		return inOutFiles, fmt.Errorf("GetIOFiles: invalid yaml signature: %w", err)
	}

	if strings.TrimSpace(inOutFiles.InFile) == "" || strings.TrimSpace(inOutFiles.OutFile) == "" {
		return inOutFiles, fmt.Errorf("GetIOFiles: %w", ErrEmptyIOFiles)
	}

	return inOutFiles, nil
}

func ReadValCurs(inFile io.Reader) (valute.ValCurs, error) {
	decoder := xml.NewDecoder(inFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var valCurs valute.ValCurs

	err := decoder.Decode(&valCurs)
	if err != nil {
		return valCurs, fmt.Errorf("ReadValCurs: %w", err)
	}

	if len(valCurs.Valutes) == 0 {
		return valCurs, fmt.Errorf("ReadValCurs: %w", ErrEmptyValueList)
	}

	for i, v := range valCurs.Valutes {
		if v.NumCode <= 0 || strings.TrimSpace(v.CharCode) == "" || v.Value < 0 {
			return valCurs, fmt.Errorf("ReadValCurs: %w at index %d", ErrInvalidFields, i)
		}
	}

	return valCurs, nil
}
