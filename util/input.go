package util

import (
	"io"
	"strings"
)

func ReadLines(reader io.Reader) []string {
	// It's OK to read it all at once instead of streaming line-by-line.
	// Advent of Code input is small and bounded.
	input, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	if len(input) == 0 {
		return []string{}
	}

	contents := strings.ReplaceAll(string(input), "\r\n", "\n")
	contents = strings.TrimSuffix(contents, "\n")

	return strings.Split(contents, "\n")
}
