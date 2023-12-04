package util

import "strings"

func Lines(input string) []string {
	return strings.Split(strings.TrimRight(input, "\n"), "\n")
}
