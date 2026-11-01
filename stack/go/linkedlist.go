package stack

type node struct {
	next    *node
	element any
}

type LinkedListStack struct {
	top  *node
	size int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

func (this *LinkedListStack) IsEmpty() bool {
	return this.top == nil
}

func (this *LinkedListStack) Push(v any) {
	node := &node{element: v, next: this.top}
	this.top = node
	this.size++
}

func (this *LinkedListStack) Pop() any {
	if this.IsEmpty() {
		return nil
	}
	node := this.top
	this.top = node.next
	this.size--
	return node.element
}

func (this *LinkedListStack) Top() any {
	if this.IsEmpty() {
		return nil
	}
	return this.top.element
}

func (this *LinkedListStack) Size() int {
	return this.size
}

func (this *LinkedListStack) Flush() {
	this.top = nil
	this.size = 0
}

// 编译期断言:LinkedListStack 实现 Stack 接口
var _ Stack = (*LinkedListStack)(nil)
