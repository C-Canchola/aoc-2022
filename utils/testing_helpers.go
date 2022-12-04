package utils

import (
	"testing"
)

// Equal checks the deep equality
func Equal[T comparable](t *testing.T, expected T, actual T) {
	if expected == actual {
		return
	}

	t.Error("expected", expected, "actual", actual)
}
