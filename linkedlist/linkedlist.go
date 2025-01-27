package linkedlist

import "fmt"

// LinkedList represents a singly linked list.
type LinkedList struct {
	head *Node
	tail *Node
}

// Node represents an element in the linked list.
type Node struct {
	Value int
	Next  *Node
}

// New creates and returns an empty linked list.
func New() *LinkedList {
	return &LinkedList{}
}

// InsertHead inserts a new node at the head of the list.
func (l *LinkedList) InsertHead(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.Next = l.head
		l.head = newNode
	}
}

// InsertTail inserts a new node at the tail of the list.
func (l *LinkedList) InsertTail(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
}

// Print prints the list elements.
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
