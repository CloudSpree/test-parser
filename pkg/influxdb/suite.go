package influxdb

import (
	"fmt"

	"github.com/CloudSpree/test-parser/pkg/test"

	"github.com/gosimple/slug"
)

func SuiteDurationFromSuite(suite test.TestSuite, baseHostname string, timestamp int64) string {
	// slugify names
	suiteName := slug.MakeLang(suite.Name, "en")

	return fmt.Sprintf(
		`suite_duration,suite=%s,hostname=%s,suite_start=%s value=%d %d`,
		checkEmpty(suiteName),
		checkEmpty(baseHostname),
		checkEmpty(suite.Start),
		suite.Duration,
		timestamp,
	)
}
