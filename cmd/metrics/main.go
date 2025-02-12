package main

import (
	"fmt"
	"log"

	"github.com/CloudSpree/test-parser/internal/process"
	"github.com/CloudSpree/test-parser/pkg/config"
	"github.com/CloudSpree/test-parser/pkg/test"
)

func main() {
	// get the config file
	cfg, err := config.NewFromFlags()
	if err != nil {
		log.Fatalf("coud not get the config: %s", err)
	}

	// open test results
	test, err := test.NewFromFile(cfg.FileName)
	if err != nil {
		log.Fatal("could not read test results")
	}

	// get metrics in prometheus format
	metrics, err := process.ProcessResult(test, cfg.Format)
	if err != nil {
		log.Fatal("could not get the metrics")
	}

	fmt.Print(metrics)
}
