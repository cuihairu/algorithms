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
