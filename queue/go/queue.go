package queue

type Queue interface {
	Enqueue(v any)
	Dequeue() any
	IsEmpty() bool
	Size() int
	Front() any
	Flush()
}
