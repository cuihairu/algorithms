package sorts

import "testing"

func TestCountingSort(t *testing.T) {
	t.Log("测试基数排序")
	arr := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0}
	t.Logf("排序之前:%+v", arr)
	CountingSort(arr)
	t.Logf("排序之后:%+v", arr)
}
