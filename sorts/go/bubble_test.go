package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Log("测试冒泡排序")
	arr := []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, 0}
	t.Logf("排序之前:%+v", arr)
	BubbleSort(arr)
	t.Logf("排序之后:%+v", arr)
}
