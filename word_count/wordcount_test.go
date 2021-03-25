package wordcount

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	var validTestCases = []struct {
		name   string
		input  string
		output map[string]int
	}{
		{
			name:  "without repeating words",
			input: "I learn Go",
			output: map[string]int{
				"I":     1,
				"learn": 1,
				"Go":    1,
			},
		},
		{
			name:  "with repeating words",
			input: "I learn Go and I like Go so much",
			output: map[string]int{
				"I":     2,
				"learn": 1,
				"Go":    2,
				"and":   1,
				"like":  1,
				"so":    1,
				"much":  1,
			},
		},
		{
			name:   "with empty string",
			input:  "",
			output: map[string]int{},
		},
	}

	for _, tc := range validTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Count(tc.input)
			if !mapEqual(got, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, got)
			}
		})
	}
}

func mapEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
