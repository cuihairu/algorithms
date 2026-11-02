package sorts

import (
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Log("测试冒泡排序")
	arr := []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, 0}
	t.Logf("排序之前:%+v", arr)
	BubbleSort(arr)
	t.Logf("排序之后:%+v", arr)
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("bubble sort failed: %+v", arr)
	}
}

// 回归:仅 t.Log 烟测曾掩盖 MergeSort P0,此处对升序/降序/重复/边界全部断言。
func TestBubbleSortCorrectness(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4, 5},          // 升序(触发提前退出)
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
		BubbleSort(got)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("BubbleSort(%v) = %v, want %v", in, got, want)
		}
	}
}
