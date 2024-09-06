# Set Package

The `set` package provides a simple implementation of a set data structure. A set is a collection of unique elements, and this implementation allows for efficient insertion, deletion, and lookup operations.

## Types

### `Set`

```go
type Set struct {
    data map[interface{}]bool
}
```

`Set` is the main structure representing a set. It holds a map with keys as the elements of the set and values as `bool`, which always store `true` to ensure uniqueness of elements.

### `SetIterator`

```go
type SetIterator struct {
    set   *Set
    index int
}
```

`SetIterator` is used to iterate over the elements of a `Set`. It maintains a reference to the `Set` and an index to track the current position within the set.

## Functions

### `New`

```go
func New() *Set
```

Creates and returns a new, empty `Set`.

### `Insert`

```go
func (s *Set) Insert(value interface{})
func Insert(s *Set, value interface{})
```

Adds a `value` to the `Set`. If the value already exists, it does nothing. Both a method receiver version and a standalone function version are provided.

### `Contains`

```go
func Contains(s *Set, value interface{}) bool
```

Checks if the `Set` contains the specified `value`. Returns `true` if the value exists in the set, otherwise `false`.

### `Erase`

```go
func (s *Set) Erase(value interface{})
func Erase(s *Set, value interface{})
```

Removes the specified `value` from the `Set`. If the value does not exist, nothing happens. Both a method receiver version and a standalone function version are provided.

### `Size`

```go
func (s *Set) Size() int
func Size(s *Set) int
```

Returns the number of elements currently in the `Set`. Both a method receiver version and a standalone function version are provided.

### `Empty`

```go
func (s *Set) Empty() bool
```

Returns `true` if the `Set` is empty, otherwise `false`.

### `Clear`

```go
func (s *Set) Clear()
```

Removes all elements from the `Set`.

## Iteration Functions

### `Begin`

```go
func (s *Set) Begin() *SetIterator
```

Returns an iterator pointing to the first element in the `Set`.

### `End`

```go
func (s *Set) End() *SetIterator
```

Returns an iterator pointing to one past the last element in the `Set`.

### `SetIterator.Next`

```go
func (si *SetIterator) Next()
```

Advances the iterator to the next element.

### `SetIterator.HasNext`

```go
func (si *SetIterator) HasNext() bool
```

Returns `true` if there are more elements to iterate over, otherwise `false`.
