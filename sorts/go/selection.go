package sorts

import "golang.org/x/exp/constraints"

/*
选择排序
@param array 数组
*/
func SelectionSort[T constraints.Ordered](array []T) {
	n := len(array)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		// 交换
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
