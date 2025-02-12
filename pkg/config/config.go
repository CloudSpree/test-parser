package config

import (
	"flag"
	"fmt"
)

type Config struct {
	FileName string
	Format   string
}

func NewFromFlags() (Config, error) {
	cfg := Config{}

	// specify flags here
	fileNamePtr := flag.String("file", "", "path to file with results")
	formatPtr := flag.String("format", "prometheus", "exported format, supported values: prometheus, influxdb")
	flag.Parse()

	// perform basic sanity checks and assign
	if *fileNamePtr == "" {
		return cfg, fmt.Errorf("file can't be empty")
	}
	cfg.FileName = *fileNamePtr
	cfg.Format = *formatPtr

	return cfg, nil
}
