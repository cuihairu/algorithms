package stack

import "testing"

// 编译期守卫:LinkedListStack 必须实现 Stack 接口
var _ Stack = (*LinkedListStack)(nil)

func TestLinkedListStackEmpty(t *testing.T) {
	s := NewLinkedListStack()
	if !s.IsEmpty() {
		t.Fatal("new stack should be empty")
	}
	if s.Size() != 0 {
		t.Fatalf("new stack size = %d, want 0", s.Size())
	}
	if s.Pop() != nil {
		t.Fatal("pop on empty stack should return nil")
	}
	if s.Top() != nil {
		t.Fatal("top on empty stack should return nil")
	}
}

func TestLinkedListStackPushPop(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(1)
	s.Push(2)
	s.Push("three")

	if s.Size() != 3 {
		t.Fatalf("size = %d, want 3", s.Size())
	}
	if s.IsEmpty() {
		t.Fatal("stack with 3 elements should not be empty")
	}
	if s.Top() != "three" {
		t.Fatalf("top = %v, want three", s.Top())
	}

	if v := s.Pop(); v != "three" {
		t.Fatalf("pop = %v, want three (LIFO)", v)
	}
	if v := s.Pop(); v != 2 {
		t.Fatalf("pop = %v, want 2", v)
	}
	if s.Size() != 1 {
		t.Fatalf("size = %d, want 1", s.Size())
	}
	if v := s.Pop(); v != 1 {
		t.Fatalf("pop = %v, want 1", v)
	}
	if !s.IsEmpty() {
		t.Fatal("stack should be empty after popping all")
	}
	if s.Pop() != nil {
		t.Fatal("pop on drained stack should return nil")
	}
}

func TestLinkedListStackFlushReuse(t *testing.T) {
	s := NewLinkedListStack()
	s.Push(9)
	s.Push(8)
	s.Flush()
	if !s.IsEmpty() {
		t.Fatal("stack should be empty after flush")
	}
	if s.Size() != 0 {
		t.Fatalf("size after flush = %d, want 0", s.Size())
	}
	if s.Top() != nil {
		t.Fatal("top after flush should return nil")
	}

	s.Push(7)
	if s.Size() != 1 {
		t.Fatalf("size after flush+push = %d, want 1", s.Size())
	}
	if v := s.Pop(); v != 7 {
		t.Fatalf("pop after flush+push = %v, want 7", v)
	}
}
