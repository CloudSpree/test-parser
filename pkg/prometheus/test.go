package prometheus

import (
	"fmt"

	"github.com/CloudSpree/test-parser/pkg/test"

	"github.com/gosimple/slug"
)

// TestDurationFromTest generates single prometheus record for test
func TestDurationFromTest(test test.Test, suite string, baseHostname string, suiteStart string) string {
	// slugify names
	testName := slug.MakeLang(test.Name, "en")
	suiteName := slug.MakeLang(suite, "en")

	return fmt.Sprintf(
		`test_duration{suite="%s",name="%s",state="%s",hostname="%s",suite_start="%s"} %d`,
		suiteName,
		testName,
		test.State,
		baseHostname,
		suiteStart,
		test.Duration,
	)
}
