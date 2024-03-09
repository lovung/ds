package queue

// Link: https://leetcode.com/problems/design-circular-queue/
type CircularQueue[T any] struct {
	buffer []T
	size   int

	head, tail int
	len        int
}

func NewCircularQueue[T any](k int) Queue[T] {
	return &CircularQueue[T]{
		buffer: make([]T, k),
		size:   k,
		head:   0,
		tail:   -1,
		len:    0,
	}
}

func (this *CircularQueue[T]) EnQueue(value T) bool {
	if this.IsFull() {
		return false
	}
	this.tail++
	if this.tail == this.size {
		this.tail = 0
	}
	this.buffer[this.tail] = value
	this.len++
	return true
}

func (this *CircularQueue[T]) DeQueue() (val T, ok bool) {
	if this.IsEmpty() {
		return val, false
	}
	val = this.buffer[this.head]
	this.head++
	if this.head == this.size {
		this.head = 0
	}
	this.len--
	return val, true
}

func (this *CircularQueue[T]) Front() (val T) {
	if this.IsEmpty() {
		return val
	}
	return this.buffer[this.head]
}

func (this *CircularQueue[T]) Rear() (val T) {
	if this.IsEmpty() {
		return val
	}
	return this.buffer[this.tail]
}

func (this *CircularQueue[T]) Len() int {
	return this.len
}

func (this *CircularQueue[T]) IsEmpty() bool {
	return this.len == 0
}

func (this *CircularQueue[T]) IsFull() bool {
	return this.len == this.size
}
