package tree

import "testing"

func TestBSTEmpty(t *testing.T) {
	b := NewBST[int]()
	if !b.IsEmpty() {
		t.Fatal("new tree should be empty")
	}
	if b.Size() != 0 {
		t.Fatalf("size = %d, want 0", b.Size())
	}
	if b.Search(1) {
		t.Fatal("search on empty tree should be false")
	}
	if b.Remove(1) {
		t.Fatal("remove on empty tree should be false")
	}
	if _, ok := b.Min(); ok {
		t.Fatal("min on empty tree should not exist")
	}
	if _, ok := b.Max(); ok {
		t.Fatal("max on empty tree should not exist")
	}
	if len(b.InOrder()) != 0 {
		t.Fatal("inorder on empty tree should be empty")
	}
}

func TestBSTInsertSearch(t *testing.T) {
	b := NewBST[int]()
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		if !b.Insert(v) {
			t.Fatalf("insert %d should succeed", v)
		}
	}
	if b.Size() != 7 {
		t.Fatalf("size = %d, want 7", b.Size())
	}
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		if !b.Search(v) {
			t.Fatalf("search %d should be true", v)
		}
	}
	for _, v := range []int{0, 2, 6, 10} {
		if b.Search(v) {
			t.Fatalf("search %d should be false", v)
		}
	}
	// 重复插入
	if b.Insert(5) {
		t.Fatal("duplicate insert should return false")
	}
	if b.Size() != 7 {
		t.Fatalf("size after duplicate insert = %d, want 7", b.Size())
	}
}

func TestBSTMinMaxInOrder(t *testing.T) {
	b := NewBST[int]()
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		b.Insert(v)
	}
	if v, ok := b.Min(); !ok || v != 1 {
		t.Fatalf("min = %v,%v want 1,true", v, ok)
	}
	if v, ok := b.Max(); !ok || v != 9 {
		t.Fatalf("max = %v,%v want 9,true", v, ok)
	}
	got := b.InOrder()
	want := []int{1, 3, 4, 5, 7, 8, 9}
	if len(got) != len(want) {
		t.Fatalf("inorder len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("inorder = %v, want %v", got, want)
		}
	}
}

func TestBSTRemove(t *testing.T) {
	b := NewBST[int]()
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		b.Insert(v)
	}

	if b.Remove(10) {
		t.Fatal("remove absent element should return false")
	}
	if b.Size() != 7 {
		t.Fatalf("size = %d, want 7", b.Size())
	}

	// 删叶子节点 1
	if !b.Remove(1) {
		t.Fatal("remove leaf should return true")
	}
	if b.Search(1) {
		t.Fatal("1 should be gone")
	}

	// 删单子节点场景:插入 0 使其成为(原 1 位置附近)叶子,再删 4(叶子),构造单子节点:删 3 后 4 已在——直接构造:5,3,8 中删 3 前 3 有子节点 4
	// 4 是叶子;给 3 挂一个单子节点:插 2 后 3 的左子为 2,删 3(单子节点:2)
	if !b.Insert(2) {
		t.Fatal("insert 2 should succeed")
	}
	if !b.Remove(3) {
		t.Fatal("remove 3 should return true")
	}
	if b.Search(3) {
		t.Fatal("3 should be gone")
	}

	// 删双子节点:根 5(左子树含 2,4,右子树含 7,8,9)
	if !b.Remove(5) {
		t.Fatal("remove root(two children) should return true")
	}
	if b.Search(5) {
		t.Fatal("5 should be gone")
	}

	if b.Size() != 5 {
		t.Fatalf("size = %d, want 5", b.Size())
	}

	// 全部删空
	for _, v := range []int{2, 4, 7, 8, 9} {
		if !b.Remove(v) {
			t.Fatalf("remove %d should return true", v)
		}
	}
	if !b.IsEmpty() {
		t.Fatal("tree should be empty after removing all")
	}
	if got := b.InOrder(); len(got) != 0 {
		t.Fatalf("inorder after emptying = %v, want empty", got)
	}
}

func TestBSTRemoveRootVariants(t *testing.T) {
	// 单节点树删根
	b := NewBST[int]()
	b.Insert(42)
	if !b.Remove(42) {
		t.Fatal("remove single root should return true")
	}
	if !b.IsEmpty() {
		t.Fatal("tree should be empty")
	}

	// 只有右子树的根
	c := NewBST[int]()
	c.Insert(2)
	c.Insert(1)
	c.Insert(3)
	if !c.Remove(2) {
		t.Fatal("remove root should return true")
	}
	if got := c.InOrder(); len(got) != 2 || got[0] != 1 || got[1] != 3 {
		t.Fatalf("inorder = %v, want [1 3]", got)
	}
}

func TestBSTHeight(t *testing.T) {
	b := NewBST[int]()
	if b.Height() != 0 {
		t.Fatalf("height of empty tree = %d, want 0", b.Height())
	}

	b.Insert(5)
	if b.Height() != 1 {
		t.Fatalf("height of single node = %d, want 1", b.Height())
	}

	for _, v := range []int{3, 8, 1, 4, 7, 9} {
		b.Insert(v)
	}
	if b.Height() != 3 {
		t.Fatalf("height = %d, want 3", b.Height())
	}

	// 退化成链:1-2-3-4
	c := NewBST[int]()
	for _, v := range []int{1, 2, 3, 4} {
		c.Insert(v)
	}
	if c.Height() != 4 {
		t.Fatalf("height of chain = %d, want 4", c.Height())
	}

	// 删空后高度归零
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		b.Remove(v)
	}
	if b.Height() != 0 {
		t.Fatalf("height after emptying = %d, want 0", b.Height())
	}
}

func TestBSTIsBalanced(t *testing.T) {
	// 空树与单节点均平衡
	b := NewBST[int]()
	if !b.IsBalanced() {
		t.Fatal("empty tree should be balanced")
	}
	b.Insert(5)
	if !b.IsBalanced() {
		t.Fatal("single node tree should be balanced")
	}

	// 完全平衡树
	for _, v := range []int{3, 8, 1, 4, 7, 9} {
		b.Insert(v)
	}
	if !b.IsBalanced() {
		t.Fatal("complete tree should be balanced")
	}

	// 一侧加深一层仍平衡(高度差 1)
	b.Insert(0)
	if !b.IsBalanced() {
		t.Fatal("tree with height diff 1 should be balanced")
	}

	// 退化成链必不平衡
	c := NewBST[int]()
	for _, v := range []int{1, 2, 3} {
		c.Insert(v)
	}
	if c.IsBalanced() {
		t.Fatal("chain tree should not be balanced")
	}
}

func TestBSTStringValues(t *testing.T) {
	b := NewBST[string]()
	for _, v := range []string{"m", "a", "z"} {
		b.Insert(v)
	}
	if v, ok := b.Min(); !ok || v != "a" {
		t.Fatalf("min = %v,%v want a,true", v, ok)
	}
	got := b.InOrder()
	if len(got) != 3 || got[0] != "a" || got[1] != "m" || got[2] != "z" {
		t.Fatalf("inorder = %v, want [a m z]", got)
	}
}
