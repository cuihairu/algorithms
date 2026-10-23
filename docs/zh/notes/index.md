---
title: 笔记
icon: note
---

# 算法笔记

> [English](/notes/) | **中文**

本仓库的数据结构与排序实现均为 Go 实现（`tree`/`heap` 使用泛型），均带单元测试。

| 主题 | 模块 | 实现 |
| --- | --- | --- |
| [数组封装](array.md) | `array/go` | 泛型数组 + 越界检查 |
| [排序算法](sorts.md) | `sorts/go` | 冒泡/插入/选择/归并/快排/计数/桶 |
| [栈](stack.md) | `stack/go` | 数组栈 / 链表栈 |
| [单链表](linkedlist.md) | `linkedlist/go` | 头哨兵 + 读写锁 + 快慢指针 |
| [队列](queue.md) | `queue/go` | 数组队列（头游标）/ 链表队列 |
| [二叉搜索树](tree.md) | `tree/go` | 泛型 BST：插入/查/三态删除/最值/中序/高度/平衡 |
| [二叉堆](heap.md) | `heap/go` | 泛型最小堆：上浮/下沉/O(n) 建堆 |
| [哈希表](hash.md) | `hash/go` | 链地址 + FNV-1a + 负载因子倍增扩容 |
| [图](graph.md) | `graph/go` | 邻接表有向/无向图 + BFS/DFS/最短路 |

## 一键跑全部测试

九个模块已纳入 `go.work`，在各模块目录内执行：

```bash
cd <module>/go && go test ./...
```
