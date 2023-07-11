package gomod

import "regexp"

func GetSemanticVersion(line string) string {
	re := regexp.MustCompile(`\bv?(?:\d+\.\d+\.\d+(?:-\w+(?:\.\d+)?)?(?:-\w+(?:\.\d+)?)?)(?:\+[^\s]+)?\b`)
	match := re.FindStringSubmatch(line)
	if len(match) > 0 {
		return match[0]
	}
	return ""
}
