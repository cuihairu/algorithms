package sorts

import "golang.org/x/exp/constraints"

/*
插入排序
@param array 数组
*/
func InsertionSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}
