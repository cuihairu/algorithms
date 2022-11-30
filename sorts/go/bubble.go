package sorts

import (
	"golang.org/x/exp/constraints"
)

/*
冒泡排序
@param array 代表数组
https://visualgo.net/zh/sorting?slide=1-3
*/
func BubbleSort[T constraints.Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		// 提前退出标志
		swap := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		// 如果没有数据交换，那么说明数组已经是排序状态可以提前退出
		if !swap {
			break
		}
	}

}
