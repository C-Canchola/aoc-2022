package utils

import (
	"reflect"
	"testing"
)

// Equal checks the deep equality
func Equal[T any](t *testing.T, expected T, actual T) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	t.Error("expected", expected, "actual", actual)
}
