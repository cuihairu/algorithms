package queue

import (
	"fmt"
	"strings"
)

type node struct {
	next    *node
	element any
}

type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (this *LinkedListQueue) IsEmpty() bool {
	return this.head == nil
}

func (this *LinkedListQueue) Size() int {
	return this.size
}

func (this *LinkedListQueue) Enqueue(v any) {
	n := &node{element: v}
	if this.IsEmpty() {
		this.head = n
	} else {
		this.tail.next = n
	}
	this.tail = n
	this.size++
}

func (this *LinkedListQueue) Dequeue() any {
	if this.IsEmpty() {
		return nil
	}
	n := this.head
	this.head = n.next
	if this.head == nil {
		this.tail = nil
	}
	this.size--
	return n.element
}

func (this *LinkedListQueue) Front() any {
	if this.IsEmpty() {
		return nil
	}
	return this.head.element
}

func (this *LinkedListQueue) Flush() {
	this.head = nil
	this.tail = nil
	this.size = 0
}

func (this *LinkedListQueue) String() string {
	if this.IsEmpty() {
		return "[]"
	}
	ss := make([]string, 0, this.size)
	for cur := this.head; cur != nil; cur = cur.next {
		ss = append(ss, fmt.Sprintf("%v", cur.element))
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, "<"))
}
