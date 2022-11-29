package sorts

import (
	"golang.org/x/exp/constraints"
)

func merge[T constraints.Ordered](array []T, start int, mid int, end int) {

}

func mergeSort[T constraints.Ordered](array []T, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(array, start, mid)
	mergeSort(array, mid+1, end)
}

func MergeSort[T constraints.Ordered](array []T) {
	n := len(array)
	if n <= 1 {
		return
	}
	mergeSort(array, 0, n-1)
}
