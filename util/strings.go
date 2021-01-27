package util

import "strings"

func IsWhitespace(s string) bool {
	return strings.TrimSpace(s) == ""
}
