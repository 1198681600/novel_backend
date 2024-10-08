package utils

import (
	"sort"
)

func SliceFilterFunc[S ~[]E, E any](s S, f func(E) bool) S {
	result := make(S, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func StableSort[T any](sources []*T, f func(a, b *T) bool) []*T {
	sort.Slice(sources, func(i, j int) bool {
		return f(sources[i], sources[j])
	})
	return sources
}
