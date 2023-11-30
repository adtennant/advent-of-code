package util

import (
	"testing"
)

type Tests []struct {
	Name     string
	Input    string
	Expected int
}

type Solution func(string) int

func (tests Tests) Run(t *testing.T, solution Solution) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := Sanitize(tt.Input)

			if actual := solution(input); actual != tt.Expected {
				t.Errorf("actual: %v, expected: %v", actual, tt.Expected)
			}
		})
	}
}
