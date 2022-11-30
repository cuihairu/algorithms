package sorts

import "testing"

func TestBucketSort(t *testing.T) {
	t.Log("测试桶排序")
	arr := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0}
	t.Logf("排序之前:%+v", arr)
	BucketSort(arr)
	t.Logf("排序之后:%+v", arr)
}
