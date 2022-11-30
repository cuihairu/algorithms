package sorts

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func getMax[T constraints.Integer](arr []T) T {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func BucketSort[T constraints.Integer](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	max := getMax(arr)
	buckets := make([][]T, n) // 二维切片

	index := 0
	for i := 0; i < n; i++ {
		v := int(arr[i])
		index = v * (n - 1) / int(max)                  // 桶序号 bug
		buckets[index] = append(buckets[index], arr[i]) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < n; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

// 桶排序简单实现
func BucketSortSimple(source []int) {
	if len(source) <= 1 {
		return
	}
	array := make([]int, getMax(source)+1)
	for i := 0; i < len(source); i++ {
		array[source[i]]++
	}
	fmt.Println(array)
	c := make([]int, 0)
	for i := 0; i < len(array); i++ {
		for array[i] != 0 {
			c = append(c, i)
			array[i]--
		}
	}
	copy(source, c)

}
