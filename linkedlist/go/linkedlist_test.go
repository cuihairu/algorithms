package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()
	if ll == nil {
		t.Fatal("NewLinkedList returned nil")
	}
	if !ll.IsEmpty() {
		t.Fatalf("new list should be empty, size = %d", ll.Size())
	}
}
