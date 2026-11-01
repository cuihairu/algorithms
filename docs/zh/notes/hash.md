---
title: 哈希表
icon: key
---

# 哈希表

> [English](/notes/hash) | **中文**

`hash/go` 模块实现字符串键哈希表，链地址法解决冲突。Go module 名为 `hashmap`，使导入路径避开标准库 `hash` 包；目录为 `hash`。

## 实现位置

- 源码：`hash/go/hashmap.go`
- 测试：`hash/go/hashmap_test.go`

## API

```go
func NewHashMap() *HashMap
func (this *HashMap) Size() int
func (this *HashMap) IsEmpty() bool
func (this *HashMap) Put(key string, value any) (old any, existed bool) // 已存在覆盖并返回旧值
func (this *HashMap) Get(key string) (any, bool)
func (this *HashMap) Contains(key string) bool
func (this *HashMap) Delete(key string) (old any, existed bool)
func (this *HashMap) Keys() []string // 顺序不定
```

## 实现要点

- 哈希函数：FNV-1a 32 位（偏移量 2166136261，质数 16777619），桶下标 `h & (len(buckets)-1)`，桶数为 2 的幂（初始 16）。
- **负载因子超过 1.0 时倍增扩容**：全部条目重哈希到新桶（`resize` 期间借 `size` 重新计数）。
- 每个桶是 `[]entry` 线性扫描；`Put` 命中已有键只覆盖不扩容计数，`Delete` 用切片拼接摘除。

## 复杂度

| 操作 | 平均 | 最坏（全碰撞） |
| --- | --- | --- |
| Put / Get / Delete / Contains | O(1) | O(n) |
| Keys | O(n) | O(n) |

## 测试状态

```bash
cd hash/go && go test ./...
```

当前 5 个测试全部通过（空表 / 基本读写 / 覆盖返回旧值 / 删除 / 扩容后仍可读写）。
