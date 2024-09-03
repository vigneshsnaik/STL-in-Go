# Stack Package

The `stack` package provides a simple, generic stack data structure for Go. A stack is a LIFO (Last In, First Out) collection, where elements are added and removed from the top of the stack.

## Types

### `Stack`

```go
type Stack struct {
    data []interface{}
}
```

`Stack` is the main structure representing a stack. It holds a slice of `interface{}` elements, allowing it to store values of any type.

## Functions

### `Push`

```go
func (s *Stack) Push(value interface{})
```

Adds a new `value` to the top of the stack.

### `Pop`

```go
func (s *Stack) Pop() (interface{}, error)
```

Removes and returns the top element from the stack. Returns an error if the stack is empty.

### `Size`

```go
func (s *Stack) Size() int
```

Returns the number of elements currently in the stack.

### `Empty`

```go
func (s *Stack) Empty() bool
```

Returns `true` if the stack is empty, otherwise `false`.

### `Top`

```go
func (s *Stack) Top() (interface{}, error)
```

Returns the top element of the stack without removing it. Returns an error if the stack is empty.

### `Clear`

```go
func (s *Stack) Clear()
```

Removes all elements from the stack.
