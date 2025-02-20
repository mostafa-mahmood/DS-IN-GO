# DS-IN-GO

DS-IN-GO is a Go module that provides implementations of common data structures. Since Go does not include built-in data structures like linked lists, stacks, and trees, this module aims to bridge that gap with efficient and easy-to-use implementations.

## Features

- **Singly Linked List**
- More data structures coming soon (Stacks, Queues, Trees, Graphs, etc.)

## Installation

To install DS-IN-GO, run:

```sh
go get github.com/mostafa-mahmood/DS-IN-GO
```

Then, import the package in your Go project:

```go
import "github.com/mostafa-mahmood/DS-IN-GO/linkedlist"
```

## Usage

### Linked List

```go
package main

import (
	"github.com/mostafa-mahmood/DS-IN-GO/linkedlist"
)

func main() {
	ll := linkedlist.NewLinkedList[int]()
}
```

### Linked List Methods

| Method | Description |
|--------|------------|
| `InsertHead(value T)` | Inserts a value at the head of the list |
| `InsertTail(value T)` | Inserts a value at the tail of the list |
| `InsertAt(index int, value T)` | Inserts a value at a specific index |
| `DeleteHead()` | Deletes the head node |
| `DeleteTail()` | Deletes the tail node |
| `DeleteByValue(value T)` | Deletes the first occurrence of a value |
| `DeleteAt(index int)` | Deletes a node at a specific index |
| `Find(value T) bool` | Checks if a value exists in the list |
| `FindIndex(value T) int` | Returns the index of a value |
| `GetAt(index int) (T, error)` | Gets the value at a specific index |
| `IsEmpty() bool` | Checks if the list is empty |
| `ListSize() int` | Returns the number of elements in the list |
| `Clear()` | Removes all elements from the list |
| `ToSlice() []T` | Converts the list to a Go slice |
| `FromSlice(values []T)` | Creates a list from a slice |
| `ForEach(func(*T))` | Applies a function to each element in the list |
| `Filter(func(T) bool) *LinkedList[T]` | Returns a new list with elements that satisfy a condition |
| `Iterator() func() (T, bool)` | Returns an iterator function to traverse the list |
| `PrintList()` | Prints the linked list |

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request with improvements or new data structures.

## License

This project is licensed under the MIT License.


