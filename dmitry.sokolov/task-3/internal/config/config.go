package config

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/dmitrys6/task-3/internal/valute"
	"golang.org/x/net/html/charset"
	"gopkg.in/yaml.v3"
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
		return inOutFiles, fmt.Errorf("GetIOFiles: input-file and output-file must not be empty")
	}

	return inOutFiles, nil
}

func ReadValCurs(inFile io.Reader) (valute.ValCurs, error) {
	decoder := xml.NewDecoder(inFile)
	decoder.CharsetReader = charset.NewReaderLabel

	var vc valute.ValCurs
	err := decoder.Decode(&vc)
	if err != nil {
		return vc, fmt.Errorf("ReadValCurs: %w", err)
	}

	if len(vc.Valutes) == 0 {
		return vc, fmt.Errorf("ReadValCurs: empty Valute list")
	}

	for i, v := range vc.Valutes {
		if strings.TrimSpace(v.NumCode) == "" || strings.TrimSpace(v.CharCode) == "" || v.Value <= 0 {
			return vc, fmt.Errorf("ReadValCurs: invalid Valute fields at index %d", i)
		}
	}

	return vc, nil
}
