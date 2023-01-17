package gogis_test

import (
	"fmt"
	"regexp"
	"strings"
)

func matcher(expectedSQL, actualSQL string) error {
	space := regexp.MustCompile(`\s+`)

	expected := strings.Trim(space.ReplaceAllString(expectedSQL, " "), "; \t\n")
	actual := strings.Trim(space.ReplaceAllString(actualSQL, " "), "; \t\n")

	if expected != actual {
		return fmt.Errorf("\n** EXPECTED ** %s\n** ACTUAL   ** %s", expected, actual)
	}

	return nil
}
