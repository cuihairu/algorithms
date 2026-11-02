// Package heap 提供最小堆数据结构。
//
// 非线程安全:Heap 不含同步机制,并发访问需调用方自行加锁。
package heap

import (
	"golang.org/x/exp/constraints"
)

// Heap 最小堆,数组实现
type Heap[T constraints.Ordered] struct {
	data []T
}

func NewHeap[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{}
}

// NewHeapFrom 由既有切片建堆(自底向上堆化,O(n))
func NewHeapFrom[T constraints.Ordered](vals []T) *Heap[T] {
	h := &Heap[T]{data: make([]T, len(vals))}
	copy(h.data, vals)
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}

func (this *Heap[T]) Size() int {
	return len(this.data)
}

func (this *Heap[T]) IsEmpty() bool {
	return len(this.data) == 0
}

// Peek 取堆顶(最小值),空堆返回 false
func (this *Heap[T]) Peek() (T, bool) {
	if this.IsEmpty() {
		var zero T
		return zero, false
	}
	return this.data[0], true
}

func (this *Heap[T]) Push(v T) {
	this.data = append(this.data, v)
	this.siftUp(len(this.data) - 1)
}

// Pop 弹出堆顶(最小值),空堆返回 false
func (this *Heap[T]) Pop() (T, bool) {
	if this.IsEmpty() {
		var zero T
		return zero, false
	}
	top := this.data[0]
	last := len(this.data) - 1
	this.data[0] = this.data[last]
	this.data = this.data[:last]
	if last > 0 {
		this.siftDown(0)
	}
	return top, true
}

func (this *Heap[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if this.data[i] >= this.data[parent] {
			break
		}
		this.data[i], this.data[parent] = this.data[parent], this.data[i]
		i = parent
	}
}

func (this *Heap[T]) siftDown(i int) {
	n := len(this.data)
	for {
		left, right := 2*i+1, 2*i+2
		smallest := i
		if left < n && this.data[left] < this.data[smallest] {
			smallest = left
		}
		if right < n && this.data[right] < this.data[smallest] {
			smallest = right
		}
		if smallest == i {
			return
		}
		this.data[i], this.data[smallest] = this.data[smallest], this.data[i]
		i = smallest
	}
}
