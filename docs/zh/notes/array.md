---
title: 数组封装
icon: config
---

# 数组封装

> [English](/notes/array) | **中文**

`array/go` 模块对 Go 原生切片做了一层泛型数组封装，带越界检查。

## 实现位置

- 源码：`array/go/array.go`
- 测试：`array/go/array_test.go`（4 个测试函数）

## 接口

```go
type Array[T any] struct{ /* ... */ }
```

| 方法 | 说明 |
| --- | --- |
| `NewArray[T any](capacity uint) *Array[T]` | 按容量创建 |
| `Len() uint` | 元素个数 |
| `Cap() uint` | 容量 |
| `Find(index uint) (T, error)` | 读取下标元素，越界返回错误 |
| `Insert(index uint, v T) error` | 在下标处插入 |
| `Append(v T) error` | 尾部追加 |
| `Pop(index uint) (T, error)` | 删除并返回下标元素 |
| `Delete(index uint) error` | 删除下标元素 |
| `String() string` | 打印 |

## 用法示例

```go
a := array.NewArray[int](10)
_ = a.Append(1)
_ = a.Append(2)
v, err := a.Pop(0) // v == 1
```

## 测试状态

```bash
cd array/go && go test ./...
```

当前 4 个测试全部通过。
