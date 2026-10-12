package queue

import "testing"

// 编译期守卫:LinkedListQueue 必须实现 Queue 接口
var _ Queue = (*LinkedListQueue)(nil)

func TestLinkedListQueueEmpty(t *testing.T) {
	q := NewLinkedListQueue()
	if !q.IsEmpty() {
		t.Fatal("new queue should be empty")
	}
	if q.Size() != 0 {
		t.Fatalf("new queue size = %d, want 0", q.Size())
	}
	if q.Dequeue() != nil {
		t.Fatal("dequeue on empty queue should return nil")
	}
	if q.Front() != nil {
		t.Fatal("front on empty queue should return nil")
	}
	if q.String() != "[]" {
		t.Fatalf("empty queue string = %q, want \"[]\"", q.String())
	}
}

func TestLinkedListQueueFIFO(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("three")

	if q.Size() != 3 {
		t.Fatalf("size = %d, want 3", q.Size())
	}
	if q.IsEmpty() {
		t.Fatal("queue with 3 elements should not be empty")
	}
	if q.Front() != 1 {
		t.Fatalf("front = %v, want 1", q.Front())
	}
	if q.String() != "[1<2<three]" {
		t.Fatalf("string = %q, want \"[1<2<three]\"", q.String())
	}

	if v := q.Dequeue(); v != 1 {
		t.Fatalf("dequeue = %v, want 1 (FIFO)", v)
	}
	if v := q.Dequeue(); v != 2 {
		t.Fatalf("dequeue = %v, want 2", v)
	}
	if q.Size() != 1 {
		t.Fatalf("size = %d, want 1", q.Size())
	}
	if v := q.Dequeue(); v != "three" {
		t.Fatalf("dequeue = %v, want three", v)
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after dequeuing all")
	}
	if q.Dequeue() != nil {
		t.Fatal("dequeue on drained queue should return nil")
	}
}

func TestLinkedListQueueFlushReuse(t *testing.T) {
	q := NewLinkedListQueue()
	q.Enqueue(9)
	q.Enqueue(8)
	q.Flush()
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after flush")
	}
	if q.Size() != 0 {
		t.Fatalf("size after flush = %d, want 0", q.Size())
	}
	if q.Front() != nil {
		t.Fatal("front after flush should return nil")
	}

	q.Enqueue(7)
	if q.Size() != 1 {
		t.Fatalf("size after flush+enqueue = %d, want 1", q.Size())
	}
	if v := q.Dequeue(); v != 7 {
		t.Fatalf("dequeue after flush+enqueue = %v, want 7", v)
	}
}
