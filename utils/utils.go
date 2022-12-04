package utils

import (
	"bufio"
	"io"
	"os"
)

// Must is a utility function to wrap a function which returns a value and an error to panic if a non nil error is returned
//
// This removes need for unecessary error checking as the code in this project is for trivial riddles.
func Must[T any](fn func() (T, error)) T {
	x, err := fn()
	if err != nil {
		panic(err)
	}
	return x
}

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
