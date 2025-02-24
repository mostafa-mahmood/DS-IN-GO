# DS-IN-GO

A Go module providing generic implementations of fundamental data structures. This package offers type-safe, efficient, and easy-to-use data structures not included in Go's standard library.

## Installation

```bash
go get github.com/mostafa-mahmood/DS-IN-GO
```

```GO
import github.com/mostafa-mahmood/DS-IN-GO/linkedlist
// replace linkedlist with the needed package name
```

## linkedlist

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

## Stack

### Constructor
- `NewStack[T comparable]() *Stack[T]` - Creates a new empty stack

### Core Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `Push(value T)` | Adds an element to the top | O(1) |
| `Pop() (T, error)` | Removes and returns the top element | O(1) |
| `Peek() (T, error)` | Returns the top element without removing | O(1) |
| `IsEmpty() bool` | Checks if stack is empty | O(1) |
| `ListSize() int` | Returns the number of elements | O(1) |
| `Clear()` | Removes all elements from the stack | O(1) |

### Utility Operations
| Method | Description | Time Complexity |
|--------|-------------|-----------------|
| `String() string` | Returns string representation of stack | O(n) |
| `Contains(value T) bool` | Checks if value exists in stack | O(n) |
| `ToSlice() []T` | Converts stack to slice (bottom to top order) | O(n) |
| `FromSlice(items []T)` | Creates stack from slice | O(n) |
| `ForEach(fn func(T))` | Applies function to each element | O(n) |

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

Future implementations planned:
- Queue
- Tree
- Graph
- Hash Table

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.