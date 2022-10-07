package helpers

import "strings"

func ValidStrLen(s string, l int) bool {
	return len(s) > l
}

// min. valid email length = 7 (eg. a@bc.de)
func ValidEmail(s string) bool {
	return strings.Contains(s, "@") && len(s) > 6 && len(s) < 20
}
