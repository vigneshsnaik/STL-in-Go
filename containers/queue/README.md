# Queue Package

The `queue` package provides a simple implementation of a queue data structure. A queue is a FIFO (First In, First Out) collection, where elements are added to the back and removed from the front.

## Types

### `Queue`

```go
type Queue struct {
    data []interface{}
}
```

`Queue` is the main structure representing a queue. It holds a slice of `interface{}` elements, allowing it to store values of any type.

## Functions

### `Enqueue`

```go
func (q *Queue) Enqueue(value interface{})
```

Adds a new `value` to the end of the queue.

### `Dequeue`

```go
func (q *Queue) Dequeue() (interface{}, error)
```

Removes and returns the front element from the queue. Returns an error if the queue is empty.

### `Size`

```go
func (q *Queue) Size() int
```

Returns the number of elements currently in the queue.

### `Empty`

```go
func (q *Queue) Empty() bool
```

Returns `true` if the queue is empty, otherwise `false`.

### `Front`

```go
func (q *Queue) Front() (interface{}, error)
```

Returns the front element of the queue without removing it. Returns an error if the queue is empty.

### `Clear`

```go
func (q *Queue) Clear()
```

Removes all elements from the queue.
