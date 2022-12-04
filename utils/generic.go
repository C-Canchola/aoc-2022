package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// FindMax is a generic function which allows the passing of a slice of any type and some accessor to return the max value of the accessor
//
// Must have len >0 or else panic will occur
//
// Also returns FIRST FOUND position of max
func FindMax[T any, V constraints.Ordered](a []T, fn func(T) V) (V, int) {
	if len(a) == 0 {
		panic("must not attempt to find max of empty slice")
	}

	isFirst := true
	var curMax V
	var curMaxPos int

	for idx, el := range a {
		potentialMax := fn(el)

		if isFirst {
			curMax = potentialMax
			curMaxPos = idx
			isFirst = false
			continue
		}

		if !(potentialMax > curMax) {
			continue
		}

		curMax = potentialMax
		curMaxPos = idx
	}
	return curMax, curMaxPos
}

type Number interface {
	constraints.Integer | constraints.Float
}

// Sort is just a convenience function which calls the std lib sort function using a more intuitive interface
func Sort[T any, V constraints.Ordered](a []T, fn func(T) V) {
	sort.SliceStable(a, func(i, j int) bool {
		return fn(a[i]) < fn(a[j])
	})
}

// Sum is a generic sum function to any numeric slice
func Sum[T Number](a []T) T {
	var sum T

	for _, el := range a {
		sum += el
	}
	return sum
}

// Copy is a convenience function to return a caopy of the provided slice since I don't remember how to use built in copy
func Copy[T any](a []T) []T {
	copyA := make([]T, len(a))
	for i, el := range a {
		copyA[i] = el
	}
	return copyA
}

// allows the mapping of a slice to another slice type by providing a function to be applied to each element
func Map[T any, K any](a []T, fn func(T) K) []K {
	mapped := make([]K, 0, len(a))
	for _, el := range a {
		mapped = append(mapped, fn(el))
	}
	return mapped
}
