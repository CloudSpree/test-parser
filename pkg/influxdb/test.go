package influxdb

import (
	"fmt"

	"github.com/CloudSpree/test-parser/pkg/test"

	"github.com/gosimple/slug"
)

func TestDurationFromTest(test test.Test, suite string, baseHostname string, suiteStart string, timestamp int64) string {
	testName := slug.MakeLang(test.Name, "en")
	suiteName := slug.MakeLang(suite, "en")

	return fmt.Sprintf(
		`test_duration,suite=%s,name=%s,state=%s,hostname=%s,suite_start=%s, value=%d %d`,
		suiteName,
		testName,
		test.State,
		baseHostname,
		suiteStart,
		test.Duration,
		timestamp,
	)
}
