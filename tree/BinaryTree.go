package tree

import (
	"fmt"
	"strings"
)

type BinaryTree[T comparable] struct {
	Head *TreeNode[T]
	Size int
}

type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

// Inserts At the first empty place creating a full binary tree
func (b *BinaryTree[T]) Insert(value T) {

	newNode := &TreeNode[T]{
		Value: value,
	}

	if b.Head == nil {
		b.Head = newNode
		b.Size++
		return
	}

	// BFS
	queue := make([]*TreeNode[T], 0)
	queue = append(queue, b.Head)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil {
			node.Left = newNode
			b.Size++
			return
		} else {
			queue = append(queue, node.Left)
		}

		if node.Right == nil {
			node.Right = newNode
			b.Size++
			return
		} else {
			queue = append(queue, node.Right)
		}
	}
}

func (b *BinaryTree[T]) BFS() []T {
	if b.Head == nil {
		return nil
	}

	queue := make([]*TreeNode[T], 0)
	queue = append(queue, b.Head)
	output := make([]T, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		output = append(output, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return output
}

func (b *BinaryTree[T]) DFSPreOrder() []T {
	if b.Head == nil {
		return []T{}
	}
	output := []T{}
	b.dfsPreOrderHelper(b.Head, &output)
	return output
}

func (b *BinaryTree[T]) dfsPreOrderHelper(node *TreeNode[T], output *[]T) {
	if node == nil {
		return
	}
	*output = append(*output, node.Value)
	b.dfsPreOrderHelper(node.Left, output)
	b.dfsPreOrderHelper(node.Right, output)
}

func (b *BinaryTree[T]) DFSInOrder() []T {
	if b.Head == nil {
		return nil
	}
	output := []T{}
	b.dfsInorderHelper(b.Head, &output)
	return output
}

func (b *BinaryTree[T]) dfsInorderHelper(node *TreeNode[T], output *[]T) {
	if node == nil {
		return
	}
	b.dfsInorderHelper(node.Left, output)
	*output = append(*output, node.Value)
	b.dfsInorderHelper(node.Right, output)
}

func (b *BinaryTree[T]) DFSPostOrder() []T {
	if b.Head == nil {
		return nil
	}
	output := []T{}
	b.dfsPostorderHelper(b.Head, &output)
	return output
}

func (b *BinaryTree[T]) dfsPostorderHelper(node *TreeNode[T], output *[]T) {
	if node == nil {
		return
	}
	b.dfsPostorderHelper(node.Left, output)
	b.dfsPostorderHelper(node.Right, output)
	*output = append(*output, node.Value)
}

func (b *BinaryTree[T]) Search(value T) bool {
	if b.Head == nil {
		return false
	}

	queue := make([]*TreeNode[T], 0)
	queue = append(queue, b.Head)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Value == value {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}

func (b *BinaryTree[T]) Height() int {
	return b.getHeight(b.Head)
}

func (b *BinaryTree[T]) getHeight(node *TreeNode[T]) int {
	if node == nil {
		return -1
	}
	leftHeight := b.getHeight(node.Left)
	rightHeight := b.getHeight(node.Right)

	return 1 + max(leftHeight, rightHeight)
}

func (b *BinaryTree[T]) Depth(value T) int {
	if b.Head == nil {
		return -1
	}

	queue := []*TreeNode[T]{b.Head}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Value == value {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return -1 // Node not found
}

func (b *BinaryTree[T]) PrintTree() {
	if b.Head == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*TreeNode[T]{b.Head}

	for len(queue) > 0 {
		levelSize := len(queue)
		nodesAtLevel := []string{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			nodesAtLevel = append(nodesAtLevel, fmt.Sprintf("%v", node.Value))

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		fmt.Println(strings.Repeat(" ", 3*(len(queue)/2)) + strings.Join(nodesAtLevel, "   "))
	}
}
