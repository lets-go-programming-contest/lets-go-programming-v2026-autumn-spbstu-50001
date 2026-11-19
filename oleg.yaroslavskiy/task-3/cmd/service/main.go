package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yarsoleg/task-3/internal/config"
	"github.com/yarsoleg/task-3/internal/parser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := parser.ConvertXMLToJSON(cfg.InputFile, cfg.OutputFile); err != nil {
		log.Fatalf("Error converting data: %v", err)
	}

	fmt.Println("Conversion completed successfully!")
}
