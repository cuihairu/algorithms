package sorts

import "golang.org/x/exp/constraints"

/*
选择排序
@param arr 数组
*/
func SelectionSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
