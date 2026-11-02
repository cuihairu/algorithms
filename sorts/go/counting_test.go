package sorts

import (
	"reflect"
	"sort"
	"testing"
)

func TestCountingSort(t *testing.T) {
	t.Log("测试计数排序")
	arr := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0}
	t.Logf("排序之前:%+v", arr)
	CountingSort(arr)
	t.Logf("排序之后:%+v", arr)
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("counting sort failed: %+v", arr)
	}
}

// 回归:仅 t.Log 烟测曾掩盖 MergeSort P0,此处对升序/降序/重复/边界全部断言。
func TestCountingSortCorrectness(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4, 5},          // 升序
		{5, 4, 3, 2, 1},          // 降序
		{3, 1, 4, 1, 5, 9, 2, 6}, // 随机含重复
		{2, 2, 2, 2},             // 全等
		{1},                      // 单元素
		{},                       // 空
		{0, -5, 3, -1, 2},        // 含负数
	}
	for _, in := range cases {
		got := append([]int(nil), in...)
		want := append([]int(nil), in...)
		CountingSort(got)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("CountingSort(%v) = %v, want %v", in, got, want)
		}
	}
}
