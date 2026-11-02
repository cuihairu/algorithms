package queue

import "testing"

// 编译期守卫:ArrayQueue 必须实现 Queue 接口
var _ Queue = (*ArrayQueue)(nil)

func TestArrayQueueEmpty(t *testing.T) {
	q := NewArrayQueue()
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

func TestArrayQueueFIFO(t *testing.T) {
	q := NewArrayQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Fatalf("size = %d, want 3", q.Size())
	}
	if q.IsEmpty() {
		t.Fatal("queue with 3 elements should not be empty")
	}
	if q.Front() != 1 {
		t.Fatalf("front = %v, want 1", q.Front())
	}
	if q.String() != "[1<2<3]" {
		t.Fatalf("string = %q, want \"[1<2<3]\"", q.String())
	}

	if v := q.Dequeue(); v != 1 {
		t.Fatalf("dequeue = %v, want 1 (FIFO)", v)
	}
	if v := q.Front(); v != 2 {
		t.Fatalf("front after dequeue = %v, want 2", v)
	}
	if v := q.Dequeue(); v != 2 {
		t.Fatalf("dequeue = %v, want 2", v)
	}
	if q.Size() != 1 {
		t.Fatalf("size = %d, want 1", q.Size())
	}
	if v := q.Dequeue(); v != 3 {
		t.Fatalf("dequeue = %v, want 3", v)
	}
	if !q.IsEmpty() {
		t.Fatal("queue should be empty after dequeuing all")
	}
	if q.Dequeue() != nil {
		t.Fatal("dequeue on drained queue should return nil")
	}
}

func TestArrayQueueFlushReuse(t *testing.T) {
	q := NewArrayQueue()
	q.Enqueue(9)
	q.Enqueue(8)
	// 出队一部分,让队首游标前移,再验证 Flush 复用
	q.Dequeue()
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

// 回归:稳态下 Enqueue 恒定 append 使底层切片容量随总操作数增长而非随当前大小;此处验证容量有界。
func TestArrayQueueBoundedMemory(t *testing.T) {
	q := NewArrayQueue()
	for i := 0; i < 16; i++ {
		q.Enqueue(i)
	}
	// 稳态:每轮入队1出队1,保持 16 个元素,共 10000 轮
	for i := 0; i < 10000; i++ {
		q.Enqueue(i)
		q.Dequeue()
	}
	if cap(q.data) > 64 {
		t.Fatalf("capacity grew to %d, want bounded near initial 32", cap(q.data))
	}
	if q.Size() != 16 {
		t.Fatalf("size = %d, want 16", q.Size())
	}
	// 压实后出队仍应保持 FIFO 顺序且元素完整;稳态末尾 16 个元素为 9984..9999
	for i := 0; i < 16; i++ {
		v, _ := q.Dequeue().(int)
		want := 9984 + i
		if v != want {
			t.Fatalf("dequeue #%d = %d, want %d (FIFO order broken after compaction)", i, v, want)
		}
	}
}
