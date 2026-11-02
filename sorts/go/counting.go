package sorts

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func CountingSort[T constraints.Integer](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	max := arr[0]
	min := max
	for i := 1; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}

	// 值域过大时 max-min+1 溢出为 0,无法分配计数数组;退化为比较排序保证正确性
	if max-min+1 <= 0 {
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		return
	}

	count := make([]int, max-min+1)
	for i := range arr {
		count[arr[i]-min]++
	}
	for i := 1; i <= int(max-min); i++ {
		count[i] += count[i-1]
	}

	tmpArr := make([]T, n)
	for i := n - 1; i >= 0; i-- {
		index := count[arr[i]-min] - 1
		tmpArr[index] = arr[i]
		count[arr[i]-min]--
	}

	copy(arr, tmpArr)
}
