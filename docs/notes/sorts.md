---
title: Sorting
icon: order
---

# Sorting

> **English** | [中文](/zh/notes/sorts)

`sorts/go` implements 7 sorting algorithms, all generic (`golang.org/x/exp/constraints`).

## Location

- Source: `sorts/go/*.go`
- Tests: `sorts/go/*_test.go`

## Overview

| Algorithm | Function | Average complexity | Stable | Notes |
| --- | --- | --- | --- | --- |
| Bubble | `BubbleSort[T]` | O(n²) | stable | |
| Insertion | `InsertionSort[T]` | O(n²) | stable | |
| Selection | `SelectionSort[T]` | O(n²) | unstable | |
| Merge | `MergeSort[T]` | O(n log n) | stable | |
| Quick | `QuickSort[T]` | O(n log n) | unstable | reused inside bucket sort |
| Counting | `CountingSort[T]` | O(n+k) | stable | supports negatives (indexed by min offset) |
| Bucket | `BucketSort[T]` | O(n+k) | stable | quick sort inside buckets |

## Historical pitfalls (fixed, regression cases live in tests)

- **Bucket sort**: the bucket index must be divided by the **bucket width**, not the **bucket count**, otherwise all-negative input (e.g. `[-10..0]`) panics on an out-of-range index; the final copy-back must scan all `nBuckets` buckets (sparse input like `[0,100]` drops elements).
- **Counting sort**: the count array must be indexed with a `min` offset, otherwise negative numbers go out of bounds directly (this bug was masked by the bucket sort panic).

## Test status

```bash
cd sorts/go && go test ./...
```

All 7 tests pass (including all-negative and sparse input cases).
