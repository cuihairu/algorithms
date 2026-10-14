package tree

import (
	"golang.org/x/exp/constraints"
)

// BST 二叉搜索树,不含重复元素
type BST[T constraints.Ordered] struct {
	root *bstNode[T]
	size int
}

type bstNode[T constraints.Ordered] struct {
	left  *bstNode[T]
	right *bstNode[T]
	val   T
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

func (this *BST[T]) Size() int {
	return this.size
}

func (this *BST[T]) IsEmpty() bool {
	return this.size == 0
}

// Insert 插入元素,重复元素不动作并返回 false
func (this *BST[T]) Insert(v T) bool {
	var inserted bool
	this.root, inserted = insertNode(this.root, v)
	if inserted {
		this.size++
	}
	return inserted
}

func insertNode[T constraints.Ordered](node *bstNode[T], v T) (*bstNode[T], bool) {
	if node == nil {
		return &bstNode[T]{val: v}, true
	}
	if v < node.val {
		var inserted bool
		node.left, inserted = insertNode(node.left, v)
		return node, inserted
	}
	if v > node.val {
		var inserted bool
		node.right, inserted = insertNode(node.right, v)
		return node, inserted
	}
	return node, false // 已存在
}

// Search 判断元素是否存在
func (this *BST[T]) Search(v T) bool {
	cur := this.root
	for cur != nil {
		if v < cur.val {
			cur = cur.left
		} else if v > cur.val {
			cur = cur.right
		} else {
			return true
		}
	}
	return false
}

// Remove 删除元素,不存在返回 false
func (this *BST[T]) Remove(v T) bool {
	var removed bool
	this.root, removed = removeNode(this.root, v)
	if removed {
		this.size--
	}
	return removed
}

func removeNode[T constraints.Ordered](node *bstNode[T], v T) (*bstNode[T], bool) {
	if node == nil {
		return nil, false
	}
	if v < node.val {
		var removed bool
		node.left, removed = removeNode(node.left, v)
		return node, removed
	}
	if v > node.val {
		var removed bool
		node.right, removed = removeNode(node.right, v)
		return node, removed
	}
	// 命中:按子节点数分三种情况
	if node.left == nil {
		return node.right, true
	}
	if node.right == nil {
		return node.left, true
	}
	// 双子节点:用右子树最小节点替换
	succ := node.right
	for succ.left != nil {
		succ = succ.left
	}
	node.val = succ.val
	node.right, _ = removeNode(node.right, succ.val)
	return node, true
}

// Min 最小值,空树返回 false
func (this *BST[T]) Min() (T, bool) {
	if this.root == nil {
		var zero T
		return zero, false
	}
	cur := this.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.val, true
}

// Max 最大值,空树返回 false
func (this *BST[T]) Max() (T, bool) {
	if this.root == nil {
		var zero T
		return zero, false
	}
	cur := this.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.val, true
}

// InOrder 中序遍历,返回升序切片
func (this *BST[T]) InOrder() []T {
	out := make([]T, 0, this.size)
	var walk func(node *bstNode[T])
	walk = func(node *bstNode[T]) {
		if node == nil {
			return
		}
		walk(node.left)
		out = append(out, node.val)
		walk(node.right)
	}
	walk(this.root)
	return out
}

// Height 返回树高(节点数),空树为 0
func (this *BST[T]) Height() int {
	h, _ := checkBalance(this.root)
	return h
}

// IsBalanced 判断是否平衡(每个节点左右子树高度差不超过 1)
func (this *BST[T]) IsBalanced() bool {
	_, ok := checkBalance(this.root)
	return ok
}

// checkBalance 一次遍历同时返回子树高与是否平衡,避免 Height 被反复重算
func checkBalance[T constraints.Ordered](node *bstNode[T]) (height int, balanced bool) {
	if node == nil {
		return 0, true
	}
	lh, lok := checkBalance(node.left)
	rh, rok := checkBalance(node.right)
	diff := lh - rh
	if diff < 0 {
		diff = -diff
	}
	if lh >= rh {
		return lh + 1, lok && rok && diff <= 1
	}
	return rh + 1, lok && rok && diff <= 1
}
