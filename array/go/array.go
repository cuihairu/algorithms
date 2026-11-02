package array

import (
	"errors"
	"fmt"
	"strings"
)

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
	// 元素读写只对已占用区间 [0, len) 合法;容量仅决定可扩展上限
	return index >= a.Len()
}

func (a *Array[T]) Find(index uint) (T, error) {
	if a.isIndexOutOfRange(index) {
		var v T
		return v, OutOfRange
	}
	return a.data[index], nil
}

// 插入值到索引index上,合法区间 [0, len](等于 len 即尾插)
func (a *Array[T]) Insert(index uint, v T) error {
	if index > a.Len() {
		return OutOfRange
	}
	if a.Len() == a.Cap() {
		return FullArray
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

// 方便打印
func (a *Array[T]) String() string {
	l := make([]string, a.len, a.len)
	for i := uint(0); i < a.len; i++ {
		l[i] = fmt.Sprintf("%v", a.data[i])
	}
	return fmt.Sprintf("[%s]", strings.Join(l, ","))
}
