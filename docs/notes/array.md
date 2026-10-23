---
title: Array
icon: config
---

# Array

> **English** | [中文](/zh/notes/array)

`array/go` wraps Go's native slice in a generic array type with bounds checks.

## Location

- Source: `array/go/array.go`
- Tests: `array/go/array_test.go` (4 test functions)

## Interface

```go
type Array[T any] struct{ /* ... */ }
```

| Method | Description |
| --- | --- |
| `NewArray[T any](capacity uint) *Array[T]` | create with capacity |
| `Len() uint` | number of elements |
| `Cap() uint` | capacity |
| `Find(index uint) (T, error)` | read element at index, out of range returns error |
| `Insert(index uint, v T) error` | insert at index |
| `Append(v T) error` | append at tail |
| `Pop(index uint) (T, error)` | remove and return element at index |
| `Delete(index uint) error` | remove element at index |
| `String() string` | print |

## Usage

```go
a := array.NewArray[int](10)
_ = a.Append(1)
_ = a.Append(2)
v, err := a.Pop(0) // v == 1
```

## Test status

```bash
cd array/go && go test ./...
```

All 4 tests pass.
