package util

import (
	"strings"
)

// Tests if a string is empty.
// Spaces and tabs at the beginning and end are trimmed.
func IsEmpty(s string) bool {
	return strings.Trim(s, " \t") == ""
}
