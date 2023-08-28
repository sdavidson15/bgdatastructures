package datastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func mapItems[T any](mp map[string]T) []T {
	lst := make([]T, 0, len(mp))
	for _, e := range mp {
		lst = append(lst, e)
	}
	return lst
}

func strPtrToStr(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return fmt.Sprintf("%v", ptr)
}
