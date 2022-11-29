package sorts

import "golang.org/x/exp/constraints"

/*
插入排序
@param array 数组
*/
func InsertionSort[T constraints.Ordered](array []T) {
	n := len(array)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > value {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = value
	}
}
