package sorts

import (
	"golang.org/x/exp/constraints"
)

func getMinAndMax[T constraints.Integer](arr []T) (T, T) {
	max := arr[0]
	min := max
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return min, max
}

func BucketSort[T constraints.Integer](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	min, max := getMinAndMax(arr)
	nBuckets := int(max-min)/n + 1
	bucketSize := int(max-min)/nBuckets + 1
	buckets := make([][]T, nBuckets) // 二维切片

	for i := 0; i < n; i++ {
		v := arr[i]
		index := int(v-min) / bucketSize           // 桶序号按桶宽计算
		buckets[index] = append(buckets[index], v) // 加入对应的桶中
	}

	tmpPos := 0 // 标记数组位置
	for i := 0; i < nBuckets; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			QuickSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
}

// 桶排序简单实现
// func BucketSortSimple(source []int) {
// 	if len(source) <= 1 {
// 		return
// 	}
// 	array := make([]int, getMax(source)+1)
// 	for i := 0; i < len(source); i++ {
// 		array[source[i]]++
// 	}
// 	fmt.Println(array)
// 	c := make([]int, 0)
// 	for i := 0; i < len(array); i++ {
// 		for array[i] != 0 {
// 			c = append(c, i)
// 			array[i]--
// 		}
// 	}
// 	copy(source, c)

// }
