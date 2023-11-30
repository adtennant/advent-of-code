package util

import "strings"

func Lines(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}

func Sanitize(input string) string {
	var trimmed []string

	for _, line := range Lines(input) {
		trimmed = append(trimmed, strings.TrimSpace(line))
	}

	return strings.Join(trimmed, "\n")
}
