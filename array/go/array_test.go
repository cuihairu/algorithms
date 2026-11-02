package array

import (
	"testing"
)

func TestCapAndLen(t *testing.T) {
	a := NewArray[int](0)
	if a.Cap() != 0 {
		t.FailNow()
	}
	t.Logf("a: %s  len:%d", a, a.Len())
	b := NewArray[string](10)

	b.Append("1")
	t.Logf("b:%s len:%d", b, b.Len())
	b.Append("2")
	t.Logf("b:%s len:%d", b, b.Len())
	if b.Cap() != 10 {
		t.FailNow()
	}
}

func TestInsert(t *testing.T) {
	a := NewArray[int](10)
	for i := 0; i < int(a.Cap())-1; i++ {
		err := a.Append(i)
		if err != nil {
			t.Error(err)
		}
	}
	t.Logf("%s", a)
	err := a.Insert(1, 100)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", a)
}

func TestFind(t *testing.T) {
	a := NewArray[int](10)
	for i := 0; i < 10; i++ {
		a.Insert(uint(i), i)
	}
	v, err := a.Find(uint(6))
	if err != nil {
		t.Error(err)
	}
	if v != 6 {
		t.FailNow()
	}
}

func TestPop(t *testing.T) {
	a := NewArray[int](10)
	a.Append(1)
	v, err := a.Pop(0)
	if err != nil {
		t.FailNow()
	}
	t.Logf("%v", v)
}

// 回归:边界必须以 len 为准,不能用 cap(否则对未占用区间返回零值/写空洞)
func TestOutOfRangeOnLogicalLength(t *testing.T) {
	a := NewArray[int](10)
	a.Append(1)
	a.Append(2) // len=2, cap=10

	if _, err := a.Find(2); err != OutOfRange {
		t.Fatalf("Find(len) want OutOfRange, got %v", err)
	}
	if _, err := a.Find(5); err != OutOfRange {
		t.Fatalf("Find(5) want OutOfRange, got %v", err)
	}
	if err := a.Insert(3, 99); err != OutOfRange {
		t.Fatalf("Insert(3) want OutOfRange, got %v", err)
	}
	if _, err := a.Pop(2); err != OutOfRange {
		t.Fatalf("Pop(2) want OutOfRange, got %v", err)
	}

	empty := NewArray[int](10)
	if _, err := empty.Pop(0); err != OutOfRange {
		t.Fatalf("Pop on empty want OutOfRange, got %v", err)
	}
	if _, err := empty.Find(0); err != OutOfRange {
		t.Fatalf("Find on empty want OutOfRange, got %v", err)
	}

	// 合法插入后内容与长度完好
	if err := a.Insert(2, 99); err != nil {
		t.Fatalf("Insert(len) want ok, got %v", err)
	}
	if a.Len() != 3 || a.String() != "[1,2,99]" {
		t.Fatalf("after tail insert want [1,2,99] len 3, got %s len %d", a.String(), a.Len())
	}
}
