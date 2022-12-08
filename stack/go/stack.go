package stack

type Stack interface {
	Push(v any)
	Pop() any
	IsEmpty() bool
	Size() int
	Top() any
	Flush()
}
