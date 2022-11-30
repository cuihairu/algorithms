package sorts

import (
	"golang.org/x/exp/constraints"
)

func merge[T constraints.Ordered](array []T, start int, mid int, end int) {
	tmpArr := make([]T, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if array[i] <= array[j] {
			tmpArr[k] = array[i]
		} else {
			tmpArr[k] = array[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tmpArr[k] = array[i]
		k++
	}
	copy(array[start:end+1], tmpArr)
}

func mergeSort[T constraints.Ordered](arr []T, start int, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
}

func MergeSort[T constraints.Ordered](array []T) {
	n := len(array)
	if n <= 1 {
		return
	}
	mergeSort(array, 0, n-1)
}
