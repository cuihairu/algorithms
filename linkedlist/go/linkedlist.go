package linkedlist

import (
	"fmt"
	"strings"
)

type ListNode struct {
	next  *ListNode
	value any
}

type LinkedList struct {
	head *ListNode // 哨兵
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: &ListNode{},
	}
}

// 打印链表
func (l *LinkedList) String() string {
	ll := make([]string, 0)
	cur := l.head.next
	for cur != nil {
		ll = append(ll, fmt.Sprintf("%+v", cur.value))
	}
	return fmt.Sprintf("[ %s ]", strings.Join(ll, "->"))
}

// 反转链表
func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}

}
