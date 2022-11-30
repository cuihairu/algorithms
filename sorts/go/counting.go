package sorts

import "golang.org/x/exp/constraints"

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

	count := make([]int, max-min+1)
	for i := range arr {
		count[arr[i]]++
	}
	for i := 1; i <= int(max-min); i++ {
		count[i] += count[i-1]
	}

	tmpArr := make([]T, n)
	for i := n - 1; i >= 0; i-- {
		index := count[arr[i]] - 1
		tmpArr[index] = arr[i]
		count[arr[i]]--
	}

	copy(arr, tmpArr)
}
