package queue

import (
	"fmt"
	"strings"
)

type Queue[T comparable] struct {
	Front *Node[T]
	Rear  *Node[T]
	Size  int
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// Constructor
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{Size: 0}
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  nil,
	}

	if q.IsEmpty() {
		q.Front = newNode
		q.Rear = newNode
	} else {
		q.Rear.Next = newNode
		q.Rear = newNode
	}
	q.Size++
}

func (q *Queue[T]) Dequeue() (T, error) {
	// queue has no elements
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("Queue Is Empty")
	}

	value := q.Front.Value
	q.Front = q.Front.Next
	q.Size--

	if q.Front == nil {
		q.Rear = nil
	}

	return value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Front == nil
}

func (q *Queue[T]) Clear() {
	q.Front = nil
	q.Rear = nil
	q.Size = 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("Queue Is Empty")
	} else {
		return q.Front.Value, nil
	}
}

func (q *Queue[T]) Contains(value T) bool {
	if q.IsEmpty() {
		return false
	}

	current := q.Front
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (q *Queue[T]) String() string {
	if q.IsEmpty() {
		return "[]"
	}
	var result strings.Builder
	result.WriteString("[")
	current := q.Front
	for current != nil {
		result.WriteString(fmt.Sprintf("%v", current.Value))
		current = current.Next
		if current != nil {
			result.WriteString(", ")
		}
	}
	result.WriteString("]")
	return result.String()
}

func (q *Queue[T]) ForEach(fn func(T)) {
	current := q.Front
	for current != nil {
		fn(current.Value)
		current = current.Next
	}
}
