package lists

import (
	"container/list"
)

// NewListFromValues return the linked list with the given array of values
func NewListFromValues[T any](values []T) *list.List {
	if len(values) == 0 {
		return nil
	}
	l := list.New()
	for i := range values {
		l.PushFront(values[i])
	}
	return l
}
