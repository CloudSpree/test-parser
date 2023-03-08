package prometheus

import (
	"fmt"

	"github.com/CloudSpree/test-parser/pkg/test"

	"github.com/gosimple/slug"
)

// SuiteDurationFromSuite generates single prometheus record for suite
func SuiteDurationFromSuite(suite test.TestSuite, baseHostname string) string {
	// slugify names
	suiteName := slug.MakeLang(suite.Name, "en")

	return fmt.Sprintf(
		`suite_duration{suite="%s",hostname="%s",suite_start="%s"} %d`,
		suiteName,
		baseHostname,
		suite.Start,
		suite.Duration,
	)
}
