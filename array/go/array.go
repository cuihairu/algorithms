package array

import "errors"

var (
	OutOfRange = errors.New("out of index range")
	FullArray  = errors.New("full array")
)

type Array[T any] struct {
	data []T
	len  uint
}

func NewArray[T any](capacity uint) *Array[T] {
	return &Array[T]{
		data: make([]T, capacity, capacity),
		len:  0,
	}
}

func (a *Array[T]) Len() uint {
	return a.len
}

// 获取当前数组定义的容量
func (a *Array[T]) Cap() uint {
	return uint(cap(a.data))
}

func (a *Array[T]) isIndexOutOfRange(index uint) bool {
	if index >= a.Cap() {
		return true
	} else {
		return false
	}
}

func (a *Array[T]) Find(index uint) (T, error) {
	if a.isIndexOutOfRange(index) {
		var v T
		return v, OutOfRange
	}
	return a.data[index], nil
}

// 插入值到索引index上
func (a *Array[T]) Insert(index uint, v T) error {
	if a.Len() == a.Cap() {
		return FullArray
	}
	if index != a.len && a.isIndexOutOfRange(index) {
		return OutOfRange
	}
	for i := a.len; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.len++
	return nil
}

// 追加尾部
func (a *Array[T]) Append(v T) error {
	return a.Insert(a.len, v)
}


func (a *Array[T]) Pop(index uint) (T, error) {
	if a.isIndexOutOfRange(index) {
		var v T
		return v, OutOfRange
	}
	v := a.data[index]
	for i := index; i < a.len-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.len--
	return v, nil
}

// 删除
func (a *Array[T]) Delete(index uint) error {
	_, err := a.Pop(index)
	return err
}
