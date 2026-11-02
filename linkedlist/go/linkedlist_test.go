package linkedlist

import "testing"

func newList(vals ...any) *LinkedList {
	l := NewLinkedList()
	for _, v := range vals {
		l.AddLast(v)
	}
	return l
}

// elements 按序取出链表全部元素(内部测试可直接访问节点)
func elements(l *LinkedList) []any {
	var out []any
	for n := l.head.next; n != nil; n = n.next {
		out = append(out, n.element)
	}
	return out
}

func sameSeq(a []any, vals ...any) bool {
	if len(a) != len(vals) {
		return false
	}
	for i := range a {
		if a[i] != vals[i] {
			return false
		}
	}
	return true
}

func TestLinkedListEmpty(t *testing.T) {
	l := NewLinkedList()
	if !l.IsEmpty() {
		t.Fatal("new list should be empty")
	}
	if l.Size() != 0 {
		t.Fatalf("size = %d, want 0", l.Size())
	}
	if l.GetFirst() != nil {
		t.Fatal("GetFirst on empty should return nil")
	}
	if l.GetLast() != nil {
		t.Fatal("GetLast on empty should return nil")
	}
	if l.RemoveFirst() != nil {
		t.Fatal("RemoveFirst on empty should return nil")
	}
	if l.RemoveLast() != nil {
		t.Fatal("RemoveLast on empty should return nil")
	}
	if l.FindMiddleNode() != nil {
		t.Fatal("FindMiddleNode on empty should return nil")
	}
	if l.HasCycle() {
		t.Fatal("empty list should have no cycle")
	}
}

func TestLinkedListAddFirstLast(t *testing.T) {
	l := NewLinkedList()
	l.AddFirst(2) // [2]
	l.AddFirst(1) // [1,2]
	l.AddLast(3)  // [1,2,3]

	if l.Size() != 3 {
		t.Fatalf("size = %d, want 3", l.Size())
	}
	if l.GetFirst() != 1 {
		t.Fatalf("GetFirst = %v, want 1", l.GetFirst())
	}
	if l.GetLast() != 3 {
		t.Fatalf("GetLast = %v, want 3", l.GetLast())
	}
	if !sameSeq(elements(l), 1, 2, 3) {
		t.Fatalf("elements = %v, want [1 2 3]", elements(l))
	}
}

func TestLinkedListRemoveFirstLast(t *testing.T) {
	l := newList(1, 2, 3)
	if v := l.RemoveFirst(); v != 1 {
		t.Fatalf("RemoveFirst = %v, want 1", v)
	}
	if l.Size() != 2 {
		t.Fatalf("size = %d, want 2", l.Size())
	}
	if v := l.RemoveLast(); v != 3 {
		t.Fatalf("RemoveLast = %v, want 3", v)
	}
	if v := l.RemoveLast(); v != 2 {
		t.Fatalf("RemoveLast = %v, want 2", v)
	}
	if !l.IsEmpty() {
		t.Fatal("list should be empty after removing all")
	}
	if l.RemoveLast() != nil {
		t.Fatal("RemoveLast on empty should return nil")
	}
}

func TestLinkedListContains(t *testing.T) {
	l := newList(1, 2, 3)
	if !l.Contains(1) {
		t.Fatal("Contains should find head element")
	}
	if !l.Contains(2) {
		t.Fatal("Contains should find middle element")
	}
	if !l.Contains(3) {
		t.Fatal("Contains should find tail element")
	}
	if l.Contains(4) {
		t.Fatal("Contains should not find absent element")
	}
	if NewLinkedList().Contains(1) {
		t.Fatal("empty list contains nothing")
	}
}

func TestLinkedListInsert(t *testing.T) {
	l := newList(1, 3)
	if err := l.Insert(2, 2); err != nil { // [1,2,3]
		t.Fatalf("Insert(2,2) returned error: %v", err)
	}
	if !sameSeq(elements(l), 1, 2, 3) {
		t.Fatalf("elements = %v, want [1 2 3]", elements(l))
	}

	l2 := newList(1, 2)
	if err := l2.Insert(0, 9); err != ErrIndexOutOfRange {
		t.Fatalf("Insert(0,..) want ErrIndexOutOfRange, got %v", err)
	}
	if l2.Size() != 2 {
		t.Fatalf("size after Insert(0,..) = %d, want 2", l2.Size())
	}

	l3 := newList(1, 2)
	if err := l3.Insert(3, 3); err != nil { // index==size+1: 追加到尾部
		t.Fatalf("Insert(size+1) returned error: %v", err)
	}
	if l3.Size() != 3 || l3.GetLast() != 3 {
		t.Fatalf("Insert at len should append, got size=%d last=%v", l3.Size(), l3.GetLast())
	}

	l4 := newList(1, 2)
	if err := l4.Insert(5, 9); err != ErrIndexOutOfRange {
		t.Fatalf("out-of-range Insert want ErrIndexOutOfRange, got %v", err)
	}
	if l4.Size() != 2 || l4.GetLast() != 2 {
		t.Fatalf("out-of-range Insert should not modify list, got size=%d last=%v", l4.Size(), l4.GetLast())
	}
}

// 回归:Insert 边界覆盖 index==1 头插与空链表合法插入(size+1==1),且越界不破坏链表。
func TestLinkedListInsertBoundaries(t *testing.T) {
	// index==1:头插
	l := newList(2, 3)
	if err := l.Insert(1, 1); err != nil {
		t.Fatalf("Insert(1,1) returned error: %v", err)
	}
	if !sameSeq(elements(l), 1, 2, 3) {
		t.Fatalf("after head insert = %v, want [1 2 3]", elements(l))
	}

	// 空链表合法插入:size+1==1,等价于头插/尾插
	empty := NewLinkedList()
	if err := empty.Insert(1, 42); err != nil {
		t.Fatalf("Insert on empty list returned error: %v", err)
	}
	if !sameSeq(elements(empty), 42) {
		t.Fatalf("empty list after insert = %v, want [42]", elements(empty))
	}

	// 空链表越界:index==0 与 index>1 均应报错且不修改
	for _, idx := range []int{0, 2} {
		e := NewLinkedList()
		if err := e.Insert(idx, 7); err != ErrIndexOutOfRange {
			t.Fatalf("Insert(%d,..) on empty want ErrIndexOutOfRange, got %v", idx, err)
		}
		if e.Size() != 0 {
			t.Fatalf("empty list Insert(%d,..) should not modify, size=%d", idx, e.Size())
		}
	}
}

func TestLinkedListReverse(t *testing.T) {
	l := newList(1, 2, 3)
	l.Reverse()
	if !sameSeq(elements(l), 3, 2, 1) {
		t.Fatalf("after reverse = %v, want [3 2 1]", elements(l))
	}

	s := newList(7)
	s.Reverse()
	if !sameSeq(elements(s), 7) {
		t.Fatalf("single reverse = %v, want [7]", elements(s))
	}

	e := NewLinkedList()
	e.Reverse()
	if !e.IsEmpty() {
		t.Fatal("empty reverse should stay empty")
	}
}

func TestLinkedListHasCycle(t *testing.T) {
	l := newList(1, 2, 3)
	if l.HasCycle() {
		t.Fatal("acyclic list reported cycle")
	}

	s := newList(1)
	if s.HasCycle() {
		t.Fatal("single node acyclic list reported cycle")
	}

	// 手工构造环:1->2->3->2
	c := &LinkedList{head: &ListNode{}}
	n1 := &ListNode{element: 1}
	n2 := &ListNode{element: 2}
	n3 := &ListNode{element: 3}
	n1.next, n2.next, n3.next = n2, n3, n2
	c.head.next = n1
	if !c.HasCycle() {
		t.Fatal("cycle not detected")
	}
}

func TestLinkedListDeleteByIndex(t *testing.T) {
	l := newList(1, 2, 3)
	l.DeleteByIndex(2) // [1,3]
	if !sameSeq(elements(l), 1, 3) {
		t.Fatalf("elements = %v, want [1 3]", elements(l))
	}
	l.DeleteByIndex(0) // index<=0: no-op
	if l.Size() != 2 {
		t.Fatalf("size after DeleteByIndex(0) = %d, want 2", l.Size())
	}
	l.DeleteByIndex(9) // 越界: no-op
	if l.Size() != 2 {
		t.Fatalf("size after out-of-range delete = %d, want 2", l.Size())
	}
	l.DeleteByIndex(1) // [3]
	if !sameSeq(elements(l), 3) {
		t.Fatalf("elements = %v, want [3]", elements(l))
	}
	l.DeleteByIndex(1) // []
	if !l.IsEmpty() {
		t.Fatal("list should be empty")
	}
}

func TestLinkedListDeleteByBottomIndex(t *testing.T) {
	l := newList(1, 2, 3)
	l.DeleteByBottomIndex(1) // 删倒数第1个(3) -> [1,2]
	if !sameSeq(elements(l), 1, 2) {
		t.Fatalf("elements = %v, want [1 2]", elements(l))
	}
	l.DeleteByBottomIndex(2) // 删倒数第2个(1) -> [2]
	if !sameSeq(elements(l), 2) {
		t.Fatalf("elements = %v, want [2]", elements(l))
	}
	l.DeleteByBottomIndex(5) // 越界: no-op
	if l.Size() != 1 {
		t.Fatalf("size after out-of-range delete = %d, want 1", l.Size())
	}
}

func TestLinkedListFindMiddleNode(t *testing.T) {
	if m := newList(1, 2, 3).FindMiddleNode(); m == nil || m.element != 2 {
		t.Fatalf("middle of [1,2,3] = %v, want 2", m)
	}
	// 偶数长度取前一个中点
	if m := newList(1, 2, 3, 4).FindMiddleNode(); m == nil || m.element != 2 {
		t.Fatalf("middle of [1,2,3,4] = %v, want 2", m)
	}
	if m := newList(7).FindMiddleNode(); m == nil || m.element != 7 {
		t.Fatalf("middle of [7] = %v, want 7", m)
	}
}

func TestLinkedListString(t *testing.T) {
	l := newList(1, 2, 3)
	if got := l.String(); got != "[ 1->2->3 ]" {
		t.Fatalf("String = %q, want \"[ 1->2->3 ]\"", got)
	}
	if got := NewLinkedList().String(); got != "[  ]" {
		t.Fatalf("empty String = %q, want \"[  ]\"", got)
	}
}
