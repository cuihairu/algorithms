package linkedlist

import (
	"fmt"
	"strings"
	"sync"
)

type ListNode struct {
	next    *ListNode
	element any
}

type LinkedList struct {
	sync.RWMutex
	head *ListNode // 哨兵
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: &ListNode{},
	}
}

// 获取第一个元素
func (this *LinkedList) GetFirst() any {
	this.RLock()
	defer this.RUnlock()
	if this.isEmpty() {
		return nil
	}
	return this.head.next.element
}

// 在头部添加元素
func (this *LinkedList) AddFirst(element any) {
	this.Lock()
	defer this.Unlock()
	if this.head == nil {
		this.head = &ListNode{}
	}
	node := &ListNode{
		next:    this.head.next,
		element: element,
	}
	this.head.next = node
}

func (this *LinkedList) GetLast() any {
	this.RLock()
	defer this.Unlock()
	if this.isEmpty() {
		return nil
	}
	tail := this.head
	for tail.next != nil {
		tail = tail.next
	}
	return tail.element
}

func (this *LinkedList) AddLast(element any) {
	this.Lock()
	defer this.Unlock()
	if this.head == nil {
		this.head = &ListNode{}
	}
	tail := this.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = &ListNode{
		element: element,
	}
}

// 打印链表
func (this *LinkedList) String() string {
	this.RLock()
	defer this.RUnlock()
	ll := make([]string, 0)
	cur := this.head.next
	for cur != nil {
		ll = append(ll, fmt.Sprintf("%+v", cur.element))
	}
	return fmt.Sprintf("[ %s ]", strings.Join(ll, "->"))
}

// 判断链表是否是空链表
func (this *LinkedList) isEmpty() bool {
	if this.head == nil || this.head.next == nil {
		return true
	}
	return false
}

// 是否是空链表
func (this *LinkedList) IsEmpty() bool {
	this.RLock()
	defer this.Unlock()
	return this.isEmpty()
}

// 反转链表
func (this *LinkedList) Reverse() {
	this.Lock()
	defer this.Unlock()

	if this.isEmpty() || this.head.next.next == nil {
		return
	}
	var pre *ListNode
	cur := this.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

// 删除第一个节点
func (this *LinkedList) RemoveFirst() any {
	this.Lock()
	defer this.Unlock()
	if this.isEmpty() {
		return nil
	}
	ret := this.head.next.element
	this.head.next = this.head.next.next
	return ret
}

// 删除尾部元素
func (this *LinkedList) RemoveLast() any {
	this.Lock()
	defer this.Unlock()
	if this.isEmpty() {
		return nil
	}
	if this.isEmpty() {
		return nil
	}
	tmp := this.head
	for tmp.next != nil && tmp.next.next != nil {
		tmp = tmp.next
	}
	ret := tmp.next.element
	tmp.next = nil
	return ret
}

// 是否包含元素
func (this *LinkedList) Contains(element any) bool {
	this.RLock()
	defer this.RUnlock()
	if this.isEmpty() {
		return false
	}
	next := this.head.next
	for next.next != nil {
		if next.element == element {
			return true
		}
		next = next.next
	}
	return false
}

func (this *LinkedList) Insert(index int, element any) {

}

// 判断单链表是否有环
func (this *LinkedList) HasCycle() bool {
	this.RLock()
	defer this.RUnlock()
	if this.isEmpty() {
		return false
	}
	slow := this.head
	fast := this.head
	for slow != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// 获取链表的长度
func (this *LinkedList) Size() int {
	if this.isEmpty() {
		return 0
	}
	num := 0
	next := this.head
	for next != nil {
		num++
		next = next.next
	}
	return num
}

// 根据索引删除节点
func (this *LinkedList) DeleteByIndex(index int) {

}

// 删除倒数的第N个节点
func (this *LinkedList) DeleteByBottomIndex(index int) {
	if index <= 0 || this.isEmpty() {
		return
	}
	fast := this.head
	for i := 1; i <= index && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}
	slow := this.head
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

// 获取中间节点
func (this *LinkedList) FindMiddleNode() *ListNode {
	if this.isEmpty() {
		return nil
	}
	if this.head.next.next == nil {
		return this.head.next
	}
	slow, fast := this.head, this.head
	for slow != nil && fast != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
