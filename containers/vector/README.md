# Vector Package

The `vector` package provides a dynamic array data structure similar to a vector in C++. It allows for efficient management of a sequence of elements, with operations to access, modify, and traverse the data.

## Types

### `Vector`

```go
type Vector struct {
    data []interface{}
}
```

`Vector` is the main structure that represents a dynamic array. It holds a slice of `interface{}` elements, allowing it to store values of any type.

### `VectorIterator`

```go
type VectorIterator struct {
    vector *Vector
    index  int
}
```

`VectorIterator` is used to traverse the elements of a `Vector`. It maintains a reference to the `Vector` and the current index within the data slice.

## Functions

### `New`

```go
func New() *Vector
```

Creates and returns a new, empty `Vector`.

### `NewVectorWithSize`

```go
func NewVectorWithSize(size int, value int) *Vector
```

Creates a new `Vector` with a specified initial `size`. Each element is initialized to the provided `value`.

### `NewVectorWithValues`

```go
func NewVectorWithValues(values ...interface{}) *Vector
```

Creates a new `Vector` initialized with the provided `values`.

### `PushBack`

```go
func (v *Vector) PushBack(value interface{})
```

Appends a new `value` to the end of the `Vector`.

### `At`

```go
func (v *Vector) At(index int) (interface{}, error)
```

Returns the element at the specified `index`. Returns an error if the index is out of range.

### `PopBack`

```go
func (v *Vector) PopBack() (interface{}, error)
```

Removes and returns the last element in the `Vector`. Returns an error if the vector is empty.

### `Size`

```go
func (v *Vector) Size() int
```

Returns the number of elements currently in the `Vector`.

### `Empty`

```go
func (v *Vector) Empty() bool
```

Returns `true` if the `Vector` is empty, otherwise `false`.

### `Reverse`

```go
func (v *Vector) Reverse()
```

Reverses the order of elements in the `Vector`.

### `Swap`

```go
func (v *Vector) Swap(i, j int) error
```

Swaps the elements at indices `i` and `j`. Returns an error if either index is out of range.

### `Back`

```go
func (v *Vector) Back() (interface{}, error)
```

Returns the last element in the `Vector`. Returns an error if the vector is empty.

### `Front`

```go
func (v *Vector) Front() (interface{}, error)
```

Returns the first element in the `Vector`. Returns an error if the vector is empty.

### `Clear`

```go
func (v *Vector) Clear()
```

Removes all elements from the `Vector`.

### `Erase`

```go
func (v *Vector) Erase(index int) error
```

Removes the element at the specified `index`. Returns an error if the index is out of range.

### `Insert`

```go
func (v *Vector) Insert(index int, value interface{}) error
```

Inserts a `value` at the specified `index`. Returns an error if the index is out of range.

### `Assign`

```go
func (v *Vector) Assign(values ...interface{})
```

Replaces the contents of the `Vector` with the provided `values`.

### `AssignWithSize`

```go
func (v *Vector) AssignWithSize(size int, value interface{})
```

Replaces the contents of the `Vector` with `size` elements, each initialized to the provided `value`.

### `Capacity`

```go
func (v *Vector) Capacity() int
```

Returns the capacity of the `Vector`'s underlying slice.

## Iteration Functions

### `Begin`

```go
func (v *Vector) Begin() *VectorIterator
```

Returns an iterator pointing to the first element of the `Vector`.

### `End`

```go
func (v *Vector) End() *VectorIterator
```

Returns an iterator pointing to one past the last element of the `Vector`.

### `VectorIterator.Next`

```go
func (vi *VectorIterator) Next()
```

Advances the iterator to the next element.

### `VectorIterator.HasNext`

```go
func (vi *VectorIterator) HasNext() bool
```

Returns `true` if there are more elements to iterate over, otherwise `false`.

### `VectorIterator.Value`

```go
func (vi *VectorIterator) Value() (interface{}, error)
```

Returns the value of the current element. Returns an error if the iterator is out of range.
