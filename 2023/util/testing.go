package util

import (
	"testing"
)

type Tests[T Result] []struct {
	Name     string
	Input    string
	Expected T
}

func (tests Tests[T]) Run(t *testing.T, solution Solution[T]) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := sanitize(tt.Input)
			actual, err := solution(input)

			if err != nil {
				t.Errorf("%v", err)
			}

			if actual != tt.Expected {
				t.Errorf("actual: %v, expected: %v", actual, tt.Expected)
			}
		})
	}
}
