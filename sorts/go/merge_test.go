package sorts

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	t.Log("合并排序")
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	t.Logf("排序之前:%+v", arr)
	MergeSort(arr)
	t.Logf("排序之后:%+v", arr)
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("not sorted: %+v", arr)
	}
}

// 回归:merge 在取左侧元素时忘了 i++,左<=右即死循环越界 panic。
// 非断言烟测曾掩盖此缺陷,此处对升序/降序/重复/随机/边界全部断言。
func TestMergeSortCorrectness(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4, 5},          // 升序(触发原 panic)
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
		MergeSort(got)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("MergeSort(%v) = %v, want %v", in, got, want)
		}
	}
}
