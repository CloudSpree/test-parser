package prometheus

import (
	"fmt"

	"github.com/CloudSpree/test-parser/pkg/test"

	"github.com/gosimple/slug"
)

// HookDurationFromHook generates single prometheus record for hook
func HookDurationFromHook(test test.Hook, suite string, baseHostname string, suiteStart string) string {
	// slugify names
	testTitle := slug.MakeLang(test.Title, "en")
	suiteName := slug.MakeLang(suite, "en")
	associatedTest := slug.MakeLang(test.AssociatedTest, "en")

	return fmt.Sprintf(
		`hook_duration{suite="%s",title="%s",state="%s",test="%s",hostname="%s",suite_start="%s"} %d`,
		suiteName,
		testTitle,
		test.State,
		associatedTest,
		baseHostname,
		suiteStart,
		test.Duration,
	)
}
