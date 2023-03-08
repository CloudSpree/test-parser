package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FileName string
}

func NewFromFlags() (Config, error) {
	cfg := Config{}

	// specify flags here
	fileNamePtr := flag.String("file", "", "path to file with results")
	flag.Parse()

	// perform basic sanity checks and assign
	if *fileNamePtr == "" {
		return cfg, fmt.Errorf("file can't be empty")
	}
	cfg.FileName = *fileNamePtr

	return cfg, nil
}
