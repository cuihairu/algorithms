---
title: 排序算法
icon: order
---

# 排序算法

> [English](sorts.md) | **中文**

`sorts/go` 模块实现 7 种排序，全部为泛型（`golang.org/x/exp/constraints`）。

## 实现位置

- 源码：`sorts/go/*.go`
- 测试：`sorts/go/*_test.go`

## 一览

| 算法 | 函数 | 平均复杂度 | 稳定性 | 备注 |
| --- | --- | --- | --- | --- |
| 冒泡 | `BubbleSort[T]` | O(n²) | 稳定 | |
| 插入 | `InsertionSort[T]` | O(n²) | 稳定 | |
| 选择 | `SelectionSort[T]` | O(n²) | 不稳定 | |
| 归并 | `MergeSort[T]` | O(n log n) | 稳定 | |
| 快排 | `QuickSort[T]` | O(n log n) | 不稳定 | 桶排序内部复用 |
| 计数 | `CountingSort[T]` | O(n+k) | 稳定 | 支持负数（按 min 偏移） |
| 桶 | `BucketSort[T]` | O(n+k) | 稳定 | 桶内快排 |

## 历史坑位（已修复，回归用例在测试里）

- **桶排序**：桶序号要用「桶宽」整除而不是「桶数」，否则全负输入（如 `[-10..0]`）时下标越界 panic；末尾回拷必须扫满 `nBuckets` 个桶（`[0,100]` 这类稀疏输入会丢元素）。
- **计数排序**：计数数组必须按 `min` 偏移下标，否则负数直接越界（此 bug 曾被桶排序的 panic 掩盖）。

## 测试状态

```bash
cd sorts/go && go test ./...
```

当前 7 个测试全部通过（含全负输入与稀疏输入用例）。
