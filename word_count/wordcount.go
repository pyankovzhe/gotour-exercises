package wordcount

import (
	"strings"
)

func Count(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}

	return m
}
