package sqlerror

import (
	"regexp"
)

func ExtractFKViolationColumn(err string) string{
	regex := regexp.MustCompile("FOREIGN KEY \\(`([^`]+)`\\) REFERENCES")
	matches := regex.FindStringSubmatch(err)

	if len(matches) == 2{
		return matches[1]
	}

	return ""
}