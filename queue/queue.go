package queue

type Queue[T any] interface {
	EnQueue(value T) bool
	DeQueue() (T, bool)
	Front() T
	Rear() T
	IsEmpty() bool
	IsFull() bool
	Len() int
}

type SimpleQueue[T any] []T

func NewSimpleQueue[T any]() Queue[T] {
	q := SimpleQueue[T](make([]T, 0))
	return &q
}

func (q *SimpleQueue[T]) EnQueue(v T) bool {
	*q = append(*q, v)
	return true
}

func (q *SimpleQueue[T]) DeQueue() (v T, ok bool) {
	if q.IsEmpty() {
		return v, false
	}
	x := (*q)[0]
	*q = (*q)[1:]
	return x, true
}

func (q *SimpleQueue[T]) Front() (v T) {
	if q.IsEmpty() {
		return v
	}
	return (*q)[0]
}

func (q *SimpleQueue[T]) Rear() (v T) {
	if q.IsEmpty() {
		return v
	}
	return (*q)[len(*q)-1]
}

func (q *SimpleQueue[T]) Len() int {
	return len(*q)
}

func (q *SimpleQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *SimpleQueue[T]) IsFull() bool {
	return false
}
