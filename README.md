# DS-IN-GO

A Go module providing generic implementations of fundamental data structures. This package offers type-safe, efficient, and easy-to-use data structures not included in Go's standard library.

## Installation

```bash
go get github.com/mostafa-mahmood/DS-IN-GO
```

## LinkedList Usage

### Basic Operations

```go
package main

import (
    "github.com/mostafa-mahmood/DS-IN-GO/linkedlist"
)

func main() {
    // Create a new integer linked list
    list := linkedlist.NewLinkedList[int]()
    
    // Add elements
    list.InsertHead(1)
    list.InsertTail(2)
    list.InsertAt(1, 3)
    
    // Print the list
    list.PrintList() // Output: Head -> 1 -> 3 -> 2 -> nil
    
    // Convert to slice
    slice := list.ToSlice() // []int{1, 3, 2}
    
    // Create from slice
    newList := linkedlist.FromSlice([]int{4, 5, 6})
}
```

### Functional Operations

```go
// Using ForEach to modify elements
list.ForEach(func(val *int) {
    *val *= 2
})

// Using Filter to create a new list
evenNumbers := list.Filter(func(val int) bool {
    return val%2 == 0
})

// Using Iterator
iter := list.Iterator()
for val, ok := iter(); ok; val, ok = iter() {
    fmt.Printf("%v ", val)
}
```

### Constructor
- `NewLinkedList[T comparable]() *LinkedList[T]` - Creates a new empty linked list

### Insertion Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `InsertHead(value T)` | Inserts at the beginning | O(1) |
| `InsertTail(value T)` | Inserts at the end | O(1) |
| `InsertAt(index int, value T) error` | Inserts at specified position | O(n) |

### Deletion Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `DeleteHead()` | Removes first element | O(1) |
| `DeleteTail()` | Removes last element | O(1) |
| `DeleteByValue(value T)` | Removes first occurrence | O(n) |
| `DeleteAt(index int) error` | Removes element at index | O(n) |

### Query Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `Find(value T) bool` | Checks if value exists | O(n) |
| `FindIndex(value T) int` | Returns index of value | O(n) |
| `GetAt(index int) (T, error)` | Returns value at index | O(n) |
| `IsEmpty() bool` | Checks if list is empty | O(1) |
| `ListSize() int` | Returns list size | O(1) |

### Utility Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `Clear()` | Removes all elements | O(1) |
| `ToSlice() []T` | Converts to slice | O(n) |
| `ForEach(func(*T))` | Applies function to elements | O(n) |
| `Filter(func(T) bool) *LinkedList[T]` | Creates filtered list | O(n) |
| `Iterator() func() (T, bool)` | Returns iterator function | O(1) |
| `PrintList()` | Prints list contents | O(n) |

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

Future implementations planned:
- Stack
- Queue
- Tree
- Graph
- Hash Table
- Graph

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
