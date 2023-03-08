package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Test struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Start    string `json:"start"`
	End      string `json:"end"`
	State    string `json:"state"`
}

type Hook struct {
	Title           string `json:"title"`
	Duration        int    `json:"duration"`
	Start           string `json:"start"`
	End             string `json:"end"`
	State           string `json:"state"`
	AssociatedSuite string `json:"associatedSuite"`
	AssociatedTest  string `json:"associatedTest"`
}

type TestSuite struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Tests    []Test `json:"tests"`
	Hooks    []Hook `json:"hooks"`
}

type TestResult struct {
	Start   string      `json:"start"`
	End     string      `json:"end"`
	BaseURL string      `json:"baseUrl"`
	Suites  []TestSuite `json:"suites"`
}

func NewFromFile(file string) (TestResult, error) {
	var test TestResult

	jsonFile, err := os.Open(file)
	if err != nil {
		return test, fmt.Errorf("could not open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return test, fmt.Errorf("could not read file: %s", err)
	}

	err = json.Unmarshal(byteValue, &test)
	if err != nil {
		return test, fmt.Errorf("could not unmarshal struct: %s", err)
	}

	return test, nil
}
