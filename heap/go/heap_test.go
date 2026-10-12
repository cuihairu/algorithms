package heap

import "testing"

func TestHeapEmpty(t *testing.T) {
	h := NewHeap[int]()
	if !h.IsEmpty() {
		t.Fatal("new heap should be empty")
	}
	if h.Size() != 0 {
		t.Fatalf("size = %d, want 0", h.Size())
	}
	if _, ok := h.Peek(); ok {
		t.Fatal("peek on empty heap should not exist")
	}
	if _, ok := h.Pop(); ok {
		t.Fatal("pop on empty heap should not exist")
	}
}

func TestHeapPushPopOrder(t *testing.T) {
	h := NewHeap[int]()
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9, 2, 6} {
		h.Push(v)
	}
	if h.Size() != 9 {
		t.Fatalf("size = %d, want 9", h.Size())
	}
	if v, ok := h.Peek(); !ok || v != 1 {
		t.Fatalf("peek = %v,%v want 1,true", v, ok)
	}

	// 弹出顺序必须升序
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, w := range want {
		v, ok := h.Pop()
		if !ok || v != w {
			t.Fatalf("pop #%d = %v,%v want %v,true", i, v, ok, w)
		}
	}
	if !h.IsEmpty() {
		t.Fatal("heap should be empty after popping all")
	}
}

func TestHeapInterleaved(t *testing.T) {
	h := NewHeap[int]()
	h.Push(10)
	h.Push(5)
	if v, _ := h.Pop(); v != 5 {
		t.Fatalf("pop = %v, want 5", v)
	}
	h.Push(1)
	h.Push(20)
	if v, _ := h.Peek(); v != 1 {
		t.Fatalf("peek = %v, want 1", v)
	}
	if v, _ := h.Pop(); v != 1 {
		t.Fatalf("pop = %v, want 1", v)
	}
	if v, _ := h.Pop(); v != 10 {
		t.Fatalf("pop = %v, want 10", v)
	}
	if v, _ := h.Pop(); v != 20 {
		t.Fatalf("pop = %v, want 20", v)
	}
	if !h.IsEmpty() {
		t.Fatal("heap should be empty")
	}
}

func TestHeapFromSlice(t *testing.T) {
	h := NewHeapFrom([]int{9, 4, 7, 1, 8, 2})
	if h.Size() != 6 {
		t.Fatalf("size = %d, want 6", h.Size())
	}
	want := []int{1, 2, 4, 7, 8, 9}
	for i, w := range want {
		v, ok := h.Pop()
		if !ok || v != w {
			t.Fatalf("pop #%d = %v,%v want %v,true", i, v, ok, w)
		}
	}
}

func TestHeapFromEmptySlice(t *testing.T) {
	h := NewHeapFrom([]int{})
	if !h.IsEmpty() {
		t.Fatal("heap from empty slice should be empty")
	}
}

func TestHeapStrings(t *testing.T) {
	h := NewHeap[string]()
	for _, v := range []string{"pear", "apple", "zebra"} {
		h.Push(v)
	}
	if v, _ := h.Pop(); v != "apple" {
		t.Fatalf("pop = %v, want apple", v)
	}
	if v, _ := h.Pop(); v != "pear" {
		t.Fatalf("pop = %v, want pear", v)
	}
	if v, _ := h.Pop(); v != "zebra" {
		t.Fatalf("pop = %v, want zebra", v)
	}
}
