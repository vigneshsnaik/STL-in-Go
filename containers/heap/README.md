# Heap Package

The `heap` package provides a basic implementation of a binary max-heap, which is a complete binary tree where the value of each node is greater than or equal to the values of its children. This structure ensures efficient retrieval and removal of the maximum value.

## Types

### `Heap`

```go
type Heap struct {
    data []int
}
```

`Heap` represents a binary max-heap. It uses a slice of integers (`[]int`) to store the heap's elements.

## Functions

### `Push`

```go
func (h *Heap) Push(value int)
```

Inserts a new `value` into the heap and restores the heap property by performing an upward heapification.

### `heapifyUp`

```go
func (h *Heap) heapifyUp(index int)
```

Restores the heap property by comparing the element at `index` with its parent and swapping them if necessary. The process continues until the heap property is satisfied or the element reaches the root.

### `Top`

```go
func (h *Heap) Top() (int, error)
```

Returns the maximum value (the root) of the heap without removing it. Returns an error if the heap is empty.

### `Pop`

```go
func (h *Heap) Pop() (int, error)
```

Removes and returns the maximum value (the root) of the heap. It restores the heap property by performing a downward heapification. Returns an error if the heap is empty.

### `heapifyDown`

```go
func (h *Heap) heapifyDown(index int)
```

Restores the heap property by comparing the element at `index` with its left and right children, swapping it with the larger child if necessary. The process continues until the heap property is satisfied or the element reaches a leaf.

### `leftChild`

```go
func leftChild(index int) int
```

Returns the index of the left child of the element at `index`.

### `rightChild`

```go
func rightChild(index int) int
```

Returns the index of the right child of the element at `index`.

### `parent`

```go
func parent(index int) int
```

Returns the index of the parent of the element at `index`.

### `Size`

```go
func (h *Heap) Size() int
```

Returns the number of elements in the heap.

### `Empty`

```go
func (h *Heap) Empty() bool
```

Returns `true` if the heap is empty, otherwise `false`.
