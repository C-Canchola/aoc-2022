package utils

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// ReadLinesFromReader scans all the lines from a reader and returns them as an array of string for each line
func ReadLinesFromReader(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic("error reading lines from reader " + scanner.Err().Error())
	}

	return lines
}

// ReadLinesFromFile returns lines from a file path
func ReadLinesFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return ReadLinesFromReader(f)
}

// Must is a wrapper which causes a panic in case of an error. Bad practice in normal circumstances but good for this
func Must[T any](fn func() (T, error)) T {
	v, err := fn()
	if err != nil {
		panic(err)
	}
	return v
}

// ParseIntMust is a helper function which panics if a string is not parseable as an integer.
// Reduces unecessary error handling as this call signifies an unparseable should never happen
func ParseIntMust(s string) int {
	return Must(func() (int, error) {
		return strconv.Atoi(s)
	})
}
