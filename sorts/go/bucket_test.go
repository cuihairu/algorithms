package sorts

import (
	"sort"
	"testing"
)

func TestBucketSort(t *testing.T) {
	t.Log("测试桶排序")
	arr := []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0}
	t.Logf("排序之前:%+v", arr)
	BucketSort(arr)
	t.Logf("排序之后:%+v", arr)
	if !sort.IntsAreSorted(arr) {
		t.Fatalf("bucket sort failed: %+v", arr)
	}
}

func TestBucketSortSparse(t *testing.T) {
	arr := []int{0, 100}
	BucketSort(arr)
	if !sort.IntsAreSorted(arr) || arr[0] != 0 || arr[1] != 100 {
		t.Fatalf("bucket sort failed: %+v", arr)
	}
}
