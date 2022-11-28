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
	v, err := a.Pop(0)
	if err != nil {
		t.FailNow()
	}
	t.Logf("%v", v)
}
