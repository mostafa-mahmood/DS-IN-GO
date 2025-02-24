package stack

import "fmt"

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
