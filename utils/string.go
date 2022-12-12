package utils

import "strings"

// StirngSplitter is a struct which provides helper functions for accessing elements of a split string
// also allows negative indexing like python
type StringSplitter struct {
	Splt []string
}

func NewStringSplitter(s string, delim string) StringSplitter {
	return StringSplitter{strings.Split(s, delim)}
}

func (s StringSplitter) GetItem(idx int) string {
	if idx < 0 {
		return s.Splt[len(s.Splt)+idx]
	}
	return s.Splt[idx]
}
