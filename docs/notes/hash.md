---
title: Hash Map
icon: key
---

# Hash Map

> **English** | [中文](/zh/notes/hash)

`hash/go` implements a string-key hash map that resolves collisions by chaining. The Go module is named `hashmap`, which keeps the import path free of the standard library's `hash` package; the directory is `hash`.

## Location

- Source: `hash/go/hashmap.go`
- Tests: `hash/go/hashmap_test.go`

## API

```go
func NewHashMap() *HashMap
func (this *HashMap) Size() int
func (this *HashMap) IsEmpty() bool
func (this *HashMap) Put(key string, value any) (old any, existed bool) // overwrites an existing key and returns the old value
func (this *HashMap) Get(key string) (any, bool)
func (this *HashMap) Contains(key string) bool
func (this *HashMap) Delete(key string) (old any, existed bool)
func (this *HashMap) Keys() []string // order not guaranteed
```

## Implementation notes

- Hash function: FNV-1a 32-bit (offset 2166136261, prime 16777619); the bucket index is `h & (len(buckets)-1)` and the bucket count is a power of two (initially 16).
- **Doubles in size when the load factor exceeds 1.0**: all entries are rehashed into the new buckets (`size` is recounted during `resize`).
- Each bucket is a `[]entry` scanned linearly; a hit on an existing key in `Put` only overwrites, without growing the count; `Delete` removes by slicing.

## Complexity

| Operation | Average | Worst (all collide) |
| --- | --- | --- |
| Put / Get / Delete / Contains | O(1) | O(n) |
| Keys | O(n) | O(n) |

## Test status

```bash
cd hash/go && go test ./...
```

All 5 tests pass (empty map / basic read-write / overwrite returns old value / delete / still readable after resize).
