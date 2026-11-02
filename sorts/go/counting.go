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

	// 值域超过阈值时退化为稳定比较排序,避免 make 分配数 TB 级计数数组导致 OOM/挂起;
	// 同时规避窄整型 max-min+1 提前回绕为 0 的性能降级。
	const maxCountingRange = 1 << 20
	span := uint64(max) - uint64(min)
	if span >= maxCountingRange {
		sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
		return
	}

	count := make([]int, span+1)
	for i := range arr {
		count[int(uint64(arr[i])-uint64(min))]++ // #nosec G115 -- 索引 < span < 1<<20,无实际溢出
	}
	for i := 1; i <= int(span); i++ {
		count[i] += count[i-1]
	}

	tmpArr := make([]T, n)
	for i := n - 1; i >= 0; i-- {
		index := count[int(uint64(arr[i])-uint64(min))] - 1 // #nosec G115 -- 索引 < span < 1<<20,无实际溢出
		tmpArr[index] = arr[i]
		count[int(uint64(arr[i])-uint64(min))]-- // #nosec G115 -- 索引 < span < 1<<20,无实际溢出
	}

	copy(arr, tmpArr)
}
