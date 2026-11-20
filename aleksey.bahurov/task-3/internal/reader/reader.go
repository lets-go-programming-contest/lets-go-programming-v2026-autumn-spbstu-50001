package reader

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"gopkg.in/yaml.v3"
)

const windows1251Charset = "windows-1251"

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FromYAMLFile[T any](path string) (T, error) {
	var result T

	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("open YAML file %q: %w", path, err)
	}

	defer func() {
		panicIfErr(file.Close())
	}()

	if err := yaml.NewDecoder(file).Decode(&result); err != nil {
		return result, fmt.Errorf("decode YAML file %q: %w", path, err)
	}

	return result, nil
}

func FromXMLFile[T any](path string) (T, error) {
	var result T

	file, err := os.Open(path)
	if err != nil {
		return result, fmt.Errorf("open XML file %q: %w", path, err)
	}

	defer func() {
		panicIfErr(file.Close())
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = xmlCharsetReader

	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("decode XML file %q: %w", path, err)
	}

	return result, nil
}

func xmlCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(charset, windows1251Charset) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}
