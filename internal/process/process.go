package process

import (
	"fmt"
	"github.com/CloudSpree/test-parser/pkg/influxdb"
	"net/url"
	"strings"
	"time"

	"github.com/CloudSpree/test-parser/pkg/prometheus"
	"github.com/CloudSpree/test-parser/pkg/test"
)

// ProcessResult generates result metrics in prometheus format
func ProcessResult(result test.TestResult, format string) (string, error) {
	if format == "prometheus" {
		return processPrometheus(result)
	}

	if format == "influxdb" {
		return processInfluxDb(result)
	}

	return "", fmt.Errorf("unknown format: %s", format)
}

func processPrometheus(result test.TestResult) (string, error) {
	// get the base domain
	url, err := url.Parse(result.BaseURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	// iterate through suites
	suites := []string{}
	tests := []string{}
	hooks := []string{}

	for _, s := range result.Suites {
		suites = append(suites, processSuitePrometheus(s, url.Hostname()))
		tests = append(tests, processSuiteTestsPrometheus(s, url.Hostname()))
		hooks = append(hooks, processSuiteHooksPromeheus(s, url.Hostname()))
	}

	// ad empty line at the end of each block
	suites = append(suites, "")
	tests = append(tests, "")
	hooks = append(hooks, "")

	// decorate metrics with types
	suites = append([]string{"# TYPE suite_duration gauge"}, suites...)
	tests = append([]string{"# TYPE test_duration gauge"}, tests...)
	hooks = append([]string{"# TYPE hook_duration gauge"}, hooks...)

	return strings.Join(suites, "\n") + strings.Join(tests, "\n") + strings.Join(hooks, "\n"), nil
}

// processSuite generates text metrics for the whole suite
func processSuitePrometheus(suite test.TestSuite, baseHostname string) string {
	suiteDuration := []string{}
	suiteDuration = append(suiteDuration, prometheus.SuiteDurationFromSuite(suite, baseHostname))
	return strings.Join(suiteDuration, "\n")
}

// processSuiteTests generates text metrics for suite's tests
func processSuiteTestsPrometheus(suite test.TestSuite, baseHostname string) string {
	testDurations := []string{}
	for _, t := range suite.Tests {
		testDurations = append(testDurations, prometheus.TestDurationFromTest(t, suite.Name, baseHostname, suite.Start))
	}
	return strings.Join(testDurations, "\n")
}

// processSuiteHooks generates text metrics for suite's hooks
func processSuiteHooksPromeheus(suite test.TestSuite, baseHostname string) string {
	hookDurations := []string{}
	for _, h := range suite.Hooks {
		hookDurations = append(hookDurations, prometheus.HookDurationFromHook(h, suite.Name, baseHostname, suite.Start))
	}
	return strings.Join(hookDurations, "\n")
}

func processInfluxDb(result test.TestResult) (string, error) {
	// get the base domain
	url, err := url.Parse(result.BaseURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	timestamp := time.Now().Unix()

	// iterate through suites
	suites := []string{}
	tests := []string{}
	hooks := []string{}

	for _, s := range result.Suites {
		suites = append(suites, processSuiteInfluxDb(s, url.Hostname(), timestamp))
		tests = append(tests, processSuiteTestsInfluxDb(s, url.Hostname(), timestamp))
		hooks = append(hooks, processSuiteHooksInfluxDb(s, url.Hostname(), timestamp))
	}

	return strings.Join(suites, "\n\n") + strings.Join(tests, "\n\n") + strings.Join(hooks, "\n"), nil
}

func processSuiteInfluxDb(suite test.TestSuite, baseHostname string, timestamp int64) string {
	suiteDuration := []string{}
	suiteDuration = append(suiteDuration, influxdb.SuiteDurationFromSuite(suite, baseHostname, timestamp))
	return strings.Join(suiteDuration, "\n")
}

func processSuiteTestsInfluxDb(suite test.TestSuite, baseHostname string, timestamp int64) string {
	testDurations := []string{}
	for _, t := range suite.Tests {
		testDurations = append(testDurations, influxdb.TestDurationFromTest(t, suite.Name, baseHostname, suite.Start, timestamp))
	}
	return strings.Join(testDurations, "\n")
}

func processSuiteHooksInfluxDb(suite test.TestSuite, baseHostname string, timestamp int64) string {
	hookDurations := []string{}
	for _, h := range suite.Hooks {
		hookDurations = append(hookDurations, influxdb.HookDurationFromHook(h, suite.Name, baseHostname, suite.Start, timestamp))
	}
	return strings.Join(hookDurations, "\n")
}
