package sorts

import (
	"golang.org/x/exp/constraints"
)

/*
冒泡排序
@param array 代表数组
https://visualgo.net/zh/sorting?slide=1-3
*/
func BubbleSort[T constraints.Ordered](array []T) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		// 提前退出标志
		swap := false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swap = true
			}
		}
		// 如果没有数据交换，那么说明数组已经是排序状态可以提前退出
		if !swap {
			break
		}
	}

}
