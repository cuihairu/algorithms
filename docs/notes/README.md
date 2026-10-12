---
title: 笔记
icon: note
---

# 算法笔记

本仓库的数据结构与排序实现，均为 Go 泛型实现，均带单元测试。

| 主题 | 模块 | 实现 |
| --- | --- | --- |
| [数组封装](array.md) | `array/go` | 泛型数组 + 越界检查 |
| [排序算法](sorts.md) | `sorts/go` | 冒泡/插入/选择/归并/快排/计数/桶 |
| [栈](stack.md) | `stack/go` | 数组栈 / 链表栈 |
| [单链表](linkedlist.md) | `linkedlist/go` | 头哨兵 + 读写锁 + 快慢指针 |

## 一键跑全部测试

四个模块已纳入 `go.work`，在各模块目录内执行：

```bash
cd <module>/go && go test ./...
```
