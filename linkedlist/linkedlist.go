// Package linkedlist provides a generic implementation of a singly linked list.
// The list is thread-unsafe and is designed for simplicity and ease of use.
package linkedlist

import (
	"fmt"
)

// LinkedList structure
type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

// Node structure
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList Constructor
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{Size: 0}
}

// Insert Node at head
func (l *LinkedList[T]) InsertHead(value T) {
	newNode := &Node[T]{
		Value: value,
		Next:  l.Head,
	}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
	l.Size++
}

// Insert Node at tail
func (l *LinkedList[T]) InsertTail(value T) {
	newNode := &Node[T]{
		Value: value,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

// Insert at specific position
func (l *LinkedList[T]) InsertAt(index int, value T) error {
	// invalid indices
	if index > l.Size {
		return fmt.Errorf("index %d is out of bounds (size: %d)", index, l.Size)
	}
	if index < 0 {
		return fmt.Errorf("index cannot be negative, got %d", index)
	}

	// Case first element
	if index == 0 {
		l.InsertHead(value)
		return nil
	}

	// Case last element
	if index == l.Size {
		l.InsertTail(value)
		return nil
	}

	// other
	newNode := &Node[T]{
		Value: value,
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	current.Next = newNode
	l.Size++

	return nil
}

// Delete Node at head
func (l *LinkedList[T]) DeleteHead() {
	// list is empty
	if l.Head == nil && l.Tail == nil {
		return
	}
	// list has only one node
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return
	}
	l.Head = l.Head.Next
	l.Size--
}

// Delete Node at tail
func (l *LinkedList[T]) DeleteTail() {
	// list is empty
	if l.Head == nil && l.Tail == nil {
		return
	}
	// list has only one node
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Size--
		return
	}

	current := l.Head
	for current.Next != l.Tail {
		current = current.Next
	}
	current.Next = nil
	l.Tail = current
	l.Size--
}

// Delete first occurance for a value
func (l *LinkedList[T]) DeleteByValue(value T) {
	if l.Head == nil {
		return
	}

	// Special case: deleting head node
	if l.Head.Value == value {
		l.Head = l.Head.Next
		if l.Head == nil { // List had only one node
			l.Tail = nil
		}
		l.Size--
		return
	}

	// For all other nodes
	current := l.Head
	for current.Next != nil && current.Next.Value != value {
		current = current.Next
	}

	// Value not found
	if current.Next == nil {
		return
	}

	// Delete the node
	current.Next = current.Next.Next
	if current.Next == nil { // We just deleted the last node
		l.Tail = current
	}
	l.Size--
}

// Delete at specific position
func (l *LinkedList[T]) DeleteAt(index int) error {
	// list is empty
	if l.Head == nil {
		return nil
	}
	// invalid indices
	if index >= l.Size { // Changed from > to >=
		return fmt.Errorf("index %d is out of bounds (size: %d)", index, l.Size)
	}
	if index < 0 {
		return fmt.Errorf("index cannot be negative, got %d", index)
	}

	// Case first element
	if index == 0 {
		l.DeleteHead()
		return nil
	}

	// Case last element
	if index == l.Size-1 { // Changed from l.Size to l.Size-1
		l.DeleteTail()
		return nil
	}

	current := l.Head
	counter := 0

	for counter < index-1 {
		current = current.Next
		counter++
	}
	current.Next = current.Next.Next
	l.Size--
	return nil
}

// Update first occurrence of a value
func (l *LinkedList[T]) Update(oldValue, newValue T) bool {
	// list is empty
	if l.Head == nil {
		return false
	}

	current := l.Head
	for current != nil {
		if current.Value == oldValue {
			current.Value = newValue
			return true
		}
		current = current.Next
	}
	return false
}

// Check if value exists
func (l *LinkedList[T]) Find(value T) bool {
	if l.Head == nil {
		return false
	}
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

// Return position of value
func (l *LinkedList[T]) FindIndex(value T) int {
	if l.Head == nil {
		return -1
	}
	index := 0
	current := l.Head

	for current != nil {
		if current.Value == value {
			return index
		}
		current = current.Next
		index++
	}
	return -1
}

// Get value at specific index
func (l *LinkedList[T]) GetAt(index int) (T, error) {
	var zeroValue T

	// list is empty
	if l.Head == nil {
		return zeroValue, fmt.Errorf("list is empty")
	}

	// invalid indices
	if index >= l.Size {
		return zeroValue, fmt.Errorf("index %d is out of bounds (size: %d)", index, l.Size)
	}
	if index < 0 {
		return zeroValue, fmt.Errorf("index cannot be negative, got %d", index)
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value, nil
}

// Check if list is empty
func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size == 0
}

// Check list size
func (l *LinkedList[T]) ListSize() int {
	return l.Size
}

// Remove all nodes
func (l *LinkedList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
	l.Size = 0
}

// Convert list to Go slice
func (l *LinkedList[T]) ToSlice() []T {
	slice := make([]T, 0, l.Size)
	current := l.Head
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

// Create list from slice
func FromSlice[T comparable](values []T) *LinkedList[T] {
	ll := NewLinkedList[T]()

	for _, value := range values {
		ll.InsertTail(value)
	}

	return ll
}

// Apply function to each element, modifying it in place
func (l *LinkedList[T]) ForEach(f func(*T)) {
	current := l.Head
	for current != nil {
		f(&current.Value)
		current = current.Next
	}
}

// Create new list with filtered elements
func (l *LinkedList[T]) Filter(f func(T) bool) *LinkedList[T] {
	ll := NewLinkedList[T]()
	if l.Head == nil {
		return ll
	}

	current := l.Head
	for current != nil {
		if f(current.Value) {
			ll.InsertTail(current.Value)
		}
		current = current.Next
	}

	return ll
}

func (l *LinkedList[T]) Iterator() func() (T, bool) {
	current := l.Head
	return func() (T, bool) {
		if current == nil {
			var zero T
			return zero, false
		}
		value := current.Value
		current = current.Next
		return value, true
	}
}

func (l *LinkedList[T]) PrintList() {
	current := l.Head
	fmt.Print("Head -> ")
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
