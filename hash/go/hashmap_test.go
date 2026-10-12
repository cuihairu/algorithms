package hash

import (
	"strconv"
	"testing"
)

func TestHashMapEmpty(t *testing.T) {
	m := NewHashMap()
	if !m.IsEmpty() {
		t.Fatal("new map should be empty")
	}
	if m.Size() != 0 {
		t.Fatalf("size = %d, want 0", m.Size())
	}
	if _, ok := m.Get("missing"); ok {
		t.Fatal("get on empty map should not exist")
	}
	if m.Contains("missing") {
		t.Fatal("contains on empty map should be false")
	}
	if _, ok := m.Delete("missing"); ok {
		t.Fatal("delete on empty map should not exist")
	}
	if len(m.Keys()) != 0 {
		t.Fatal("keys on empty map should be empty")
	}
}

func TestHashMapPutGet(t *testing.T) {
	m := NewHashMap()
	m.Put("a", 1)
	m.Put("b", 2)
	m.Put("c", 3)

	if m.Size() != 3 {
		t.Fatalf("size = %d, want 3", m.Size())
	}
	for k, want := range map[string]int{"a": 1, "b": 2, "c": 3} {
		if v, ok := m.Get(k); !ok || v != want {
			t.Fatalf("get %q = %v,%v want %v,true", k, v, ok, want)
		}
		if !m.Contains(k) {
			t.Fatalf("contains %q should be true", k)
		}
	}
	if _, ok := m.Get("d"); ok {
		t.Fatal("get absent key should not exist")
	}

	if len(m.Keys()) != 3 {
		t.Fatalf("keys len = %d, want 3", len(m.Keys()))
	}
}

func TestHashMapOverwrite(t *testing.T) {
	m := NewHashMap()
	m.Put("k", "old")
	old, existed := m.Put("k", "new")
	if !existed || old != "old" {
		t.Fatalf("overwrite old = %v,%v want old,true", old, existed)
	}
	if v, _ := m.Get("k"); v != "new" {
		t.Fatalf("get = %v, want new", v)
	}
	if m.Size() != 1 {
		t.Fatalf("size = %d, want 1", m.Size())
	}
}

func TestHashMapDelete(t *testing.T) {
	m := NewHashMap()
	m.Put("x", 10)
	m.Put("y", 20)

	old, existed := m.Delete("x")
	if !existed || old != 10 {
		t.Fatalf("delete = %v,%v want 10,true", old, existed)
	}
	if m.Contains("x") {
		t.Fatal("x should be gone")
	}
	if m.Size() != 1 {
		t.Fatalf("size = %d, want 1", m.Size())
	}
	// 删除后其余键仍可读
	if v, _ := m.Get("y"); v != 20 {
		t.Fatalf("get y = %v, want 20", v)
	}
}

func TestHashMapResize(t *testing.T) {
	m := NewHashMap()
	const n = 500
	for i := 0; i < n; i++ {
		m.Put(keys[i], i)
	}
	if m.Size() != n {
		t.Fatalf("size = %d, want %d", m.Size(), n)
	}
	// 扩容后全部可读且值正确
	for i := 0; i < n; i++ {
		if v, ok := m.Get(keys[i]); !ok || v != i {
			t.Fatalf("get %q = %v,%v want %d,true", keys[i], v, ok, i)
		}
	}
	// 抽样删除一半,再验证剩余
	for i := 0; i < n; i += 2 {
		if _, ok := m.Delete(keys[i]); !ok {
			t.Fatalf("delete %q should exist", keys[i])
		}
	}
	if m.Size() != n/2 {
		t.Fatalf("size = %d, want %d", m.Size(), n/2)
	}
	for i := 1; i < n; i += 2 {
		if v, ok := m.Get(keys[i]); !ok || v != i {
			t.Fatalf("get %q = %v,%v want %d,true", keys[i], v, ok, i)
		}
	}
}

// 500 个互异键,覆盖扩容路径
var keys = buildKeys()

func buildKeys() []string {
	out := make([]string, 0, 500)
	for i := 0; i < 500; i++ {
		out = append(out, "key-"+strconv.Itoa(i))
	}
	return out
}
