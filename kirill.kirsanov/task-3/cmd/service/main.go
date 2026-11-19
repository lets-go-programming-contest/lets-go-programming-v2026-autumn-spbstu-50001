package main

import (
	"flag"
	"fmt"

	"github.com/kirillkirsanov2007/task-3/internal/config"
	"github.com/kirillkirsanov2007/task-3/internal/converter"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	if err := converter.ProcessData(cfg.InputFile, cfg.OutputFile); err != nil {
		panic(fmt.Sprintf("failed to process data: %v", err))
	}

	fmt.Println("Success! Data saved to", cfg.OutputFile)
}
