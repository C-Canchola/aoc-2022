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
func MapFn[T any, K any](a []T, fn func(T) K) []K {
	mapped := make([]K, 0, len(a))
	for _, el := range a {
		mapped = append(mapped, fn(el))
	}
	return mapped
}

// Reduce applies a reduction to a slice by taking a function which acts upon each successive element
// and returns/update the reduced value
func Reduce[T any, K any](a []T, initial K, fn func(T, K) K) K {
	reducedValue := initial
	for _, el := range a {
		reducedValue = fn(el, reducedValue)
	}
	return reducedValue
}

// UniqueMap converts any ordered slice to a map with each element found with a true value
func UniqueMap[T comparable](a []T) map[T]bool {
	m := make(map[T]bool)
	for _, el := range a {
		m[el] = true
	}
	return m
}

// MapKeys returns a slice of only the keys of a map
func MapKeys[T comparable, K any](m map[T]K) []T {
	a := make([]T, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

// Intersect is a utility function which returns the unique elements which appear in both slices
func Intersect[T constraints.Ordered](as ...[]T) []T {
	if len(as) == 0 {
		return []T{}
	}

	uniqueMap := Reduce(as, nil, func(a []T, m map[T]bool) map[T]bool {
		if m == nil {
			return UniqueMap(a)
		}

		newM := map[T]bool{}
		for _, el := range a {
			if !m[el] {
				continue
			}
			newM[el] = true
		}
		return newM
	})
	return MapKeys(uniqueMap)
}
