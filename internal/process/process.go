package process

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/CloudSpree/test-parser/pkg/prometheus"
	"github.com/CloudSpree/test-parser/pkg/test"
)

// ProcessResult generates result metrics in prometheus format
func ProcessResult(result test.TestResult) (string, error) {
	// get the base domain
	url, err := url.Parse(result.BaseURL)
	if err != nil {
		return "", fmt.Errorf("could not parse provided URL: %s", err)
	}

	// iterate through suites
	resultMetrics := []string{}
	for _, s := range result.Suites {
		resultMetrics = append(resultMetrics, processSuite(s, url.Hostname()))
		resultMetrics = append(resultMetrics, processSuiteTests(s, url.Hostname()))
		resultMetrics = append(resultMetrics, processSuiteHooks(s, url.Hostname()))
	}

	return strings.Join(resultMetrics, "\n"), nil
}

// processSuite generates text metrics for the whole suite
func processSuite(suite test.TestSuite, baseHostname string) string {
	suiteDuration := []string{
		"# TYPE suite_duration gauge",
	}
	suiteDuration = append(suiteDuration, prometheus.SuiteDurationFromSuite(suite, baseHostname))
	suiteDuration = append(suiteDuration, "")
	return strings.Join(suiteDuration, "\n")
}

// processSuiteTests generates text metrics for suite's tests
func processSuiteTests(suite test.TestSuite, baseHostname string) string {
	testDurations := []string{
		"# TYPE test_duration gauge",
	}
	for _, t := range suite.Tests {
		testDurations = append(testDurations, prometheus.TestDurationFromTest(t, suite.Name, baseHostname, suite.Start))
	}
	testDurations = append(testDurations, "")

	return strings.Join(testDurations, "\n")
}

// processSuiteHooks generates text metrics for suite's hooks
func processSuiteHooks(suite test.TestSuite, baseHostname string) string {
	hookDurations := []string{
		"# TYPE hook_duration gauge",
	}
	for _, h := range suite.Hooks {
		hookDurations = append(hookDurations, prometheus.HookDurationFromHook(h, suite.Name, baseHostname, suite.Start))
	}
	hookDurations = append(hookDurations, "")

	return strings.Join(hookDurations, "\n")
}
