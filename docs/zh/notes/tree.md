---
title: 二叉搜索树
icon: tree
---

# 二叉搜索树

> [English](/notes/tree) | **中文**

`tree/go` 模块实现泛型二叉搜索树（BST），不含重复元素。

## 实现位置

- 源码：`tree/go/bst.go`
- 测试：`tree/go/bst_test.go`

## API

```go
func NewBST[T constraints.Ordered]() *BST[T]
func (this *BST[T]) Size() int
func (this *BST[T]) IsEmpty() bool
func (this *BST[T]) Insert(v T) bool   // 重复元素不动作并返回 false
func (this *BST[T]) Search(v T) bool
func (this *BST[T]) Remove(v T) bool   // 不存在返回 false
func (this *BST[T]) Min() (T, bool)    // 空树返回 false
func (this *BST[T]) Max() (T, bool)
func (this *BST[T]) InOrder() []T      // 返回升序切片
func (this *BST[T]) Height() int       // 节点数树高,空树为 0
func (this *BST[T]) IsBalanced() bool  // 每节点左右子树高度差 ≤ 1
```

## 实现要点

- 插入/删除为递归实现，返回值带「新子树根 + 是否生效」，由外层统一维护 `size`；查询为迭代（从根向下走）。
- 删除分三种情况：叶子/单子直接摘除，**双子节点用右子树最小节点（后继）替换**后转去右子树删后继。
- `InOrder` 用闭包递归中序遍历，容量按 `size` 预分配。
- `Height` 与 `IsBalanced` 共用一次递归：`checkBalance` 单趟同时返回子树高与平衡性，O(n) 一趟出结果，避免 `Height` 在平衡判定里被反复重算。

## 复杂度

| 操作 | 平均 | 最坏（退化成链） |
| --- | --- | --- |
| Insert / Search / Remove | O(log n) | O(n) |
| Min / Max | O(h) | O(n) |
| InOrder | O(n) | O(n) |
| Height / IsBalanced | O(n) | O(n) |

## 测试状态

```bash
cd tree/go && go test ./...
```

当前 6 个测试全部通过（空树 / 插入查重 / Min/Max/中序 / 三态删除 / 根节点三变体 / 字符串泛型）。
