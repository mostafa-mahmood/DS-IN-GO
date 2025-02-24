package stack

import (
	"fmt"
	"strings"
)

type Stack[T comparable] struct {
	Top  *Node[T]
	Size int
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// Constructor
func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// Methods

// adds element to the top of the stack
func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  s.Top,
	}
	s.Size++
	s.Top = newNode
}

// returns and remove element from the top of the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("Stack is Empty")
	}
	data := s.Top.Value
	s.Top = s.Top.Next
	s.Size--
	return data, nil
}

// returns the element on the top of the stack
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("Stack is Empty")
	}
	return s.Top.Value, nil
}

func (s *Stack[T]) String() string {
	if s.IsEmpty() {
		return "[]"
	}
	var result strings.Builder
	result.WriteString("[")
	current := s.Top
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

func (s *Stack[T]) Contains(value T) bool {
	current := s.Top
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (s *Stack[T]) ToSlice() []T {
	result := make([]T, 0, s.Size)
	current := s.Top
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	// Reverse to match natural stack order (bottom to top)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func (s *Stack[T]) FromSlice(items []T) {
	s.Clear()
	// Push in reverse to maintain original order
	for i := len(items) - 1; i >= 0; i-- {
		s.Push(items[i])
	}
}

func (s *Stack[T]) ForEach(fn func(T)) {
	current := s.Top
	for current != nil {
		fn(current.Value)
		current = current.Next
	}
}

// return the number of elements in the stack
func (s *Stack[T]) ListSize() int {
	return s.Size
}

// deletes all the elmenets in the stack
func (s *Stack[T]) Clear() {
	s.Top = nil
	s.Size = 0
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Top == nil
}
