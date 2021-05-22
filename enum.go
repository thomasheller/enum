package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
)

func enumerate(s *bufio.Scanner, w io.Writer, left, none bool, separator string) error {
	if left && none {
		return errors.New("Invalid options: left and none cannot be combined (see -help)")
	}

	var lines []string

	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("Error reading from stdin: %s", err)
	}

	max := intLen(len(lines)) + len(separator)

	for i, line := range lines {
		n := i + 1

		prefix := fmt.Sprintf("%d%s", n, separator)

		var padded string

		if none {
			padded = prefix
		} else if left {
			padded = padRight(prefix, max)
		} else {
			padded = padLeft(prefix, max)
		}

		fmt.Fprintf(w, "%s%s\n", padded, line)
	}

	return nil
}

func intLen(i int) int {
	return int(math.Log10(float64(i))) + 1
}

func padLeft(s string, max int) string {
	return pad(s, max) + s
}

func padRight(s string, max int) string {
	return s + pad(s, max)
}

func pad(s string, max int) string {
	return strings.Repeat(" ", max-len(s))
}
