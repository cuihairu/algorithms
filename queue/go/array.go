// Package queue 提供队列数据结构。
//
// 非线程安全:ArrayQueue/LinkedListQueue 不含同步机制,并发访问需调用方自行加锁。
package queue

import (
	"fmt"
	"strings"
)

// 编译期接口守卫:ArrayQueue 必须实现 Queue
var _ Queue = (*ArrayQueue)(nil)

type ArrayQueue struct {
	data []any
	head int // 队首游标
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		data: make([]any, 0, 32),
	}
}

func (this *ArrayQueue) IsEmpty() bool {
	return len(this.data) == this.head
}

func (this *ArrayQueue) Size() int {
	return len(this.data) - this.head
}

func (this *ArrayQueue) Enqueue(v any) {
	// 队列为空时游标已重置,直接追加即可
	this.data = append(this.data, v)
}

func (this *ArrayQueue) Dequeue() any {
	if this.IsEmpty() {
		return nil
	}
	v := this.data[this.head]
	this.head++
	// 清空后回收切片,避免游标单调增长
	if this.IsEmpty() {
		this.data = this.data[:0]
		this.head = 0
	}
	return v
}

func (this *ArrayQueue) Front() any {
	if this.IsEmpty() {
		return nil
	}
	return this.data[this.head]
}

func (this *ArrayQueue) Flush() {
	this.data = this.data[:0]
	this.head = 0
}

func (this *ArrayQueue) String() string {
	if this.IsEmpty() {
		return "[]"
	}
	ss := make([]string, 0, this.Size())
	for i := this.head; i < len(this.data); i++ {
		ss = append(ss, fmt.Sprintf("%v", this.data[i]))
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, "<"))
}
