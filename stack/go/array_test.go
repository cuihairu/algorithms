package stack

import "testing"

// 编译期守卫:ArrayStack 必须实现 Stack 接口
var _ Stack = (*ArrayStack)(nil)

func TestArrayStackEmpty(t *testing.T) {
	s := NewArrayStack()
	if !s.IsEmpty() {
		t.Fatalf("new stack should be empty")
	}
	if s.Size() != 0 {
		t.Fatalf("new stack size = %d, want 0", s.Size())
	}
	if s.Pop() != nil {
		t.Fatalf("pop on empty stack = %v, want nil", s.Pop())
	}
	if s.Top() != nil {
		t.Fatalf("top on empty stack = %v, want nil", s.Top())
	}
	if s.String() != "[]" {
		t.Fatalf("empty stack string = %q, want \"[]\"", s.String())
	}
}

func TestArrayStackPushPop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Size() != 3 {
		t.Fatalf("size = %d, want 3", s.Size())
	}
	if s.IsEmpty() {
		t.Fatal("stack with 3 elements should not be empty")
	}
	if s.Top() != 3 {
		t.Fatalf("top = %v, want 3", s.Top())
	}
	if s.String() != "[1<2<3]" {
		t.Fatalf("string = %q, want \"[1<2<3]\"", s.String())
	}

	if v := s.Pop(); v != 3 {
		t.Fatalf("pop = %v, want 3 (LIFO)", v)
	}
	if v := s.Top(); v != 2 {
		t.Fatalf("top after pop = %v, want 2", v)
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

func TestArrayStackFlushReuse(t *testing.T) {
	s := NewArrayStack()
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
		t.Fatalf("top after flush = %v, want nil", s.Top())
	}

	// 复用:flush 后再 push,逻辑大小必须正确
	s.Push(7)
	if s.Size() != 1 {
		t.Fatalf("size after flush+push = %d, want 1", s.Size())
	}
	if v := s.Pop(); v != 7 {
		t.Fatalf("pop after flush+push = %v, want 7", v)
	}
}
