package heap

import (
	"fmt"
)

type Heap struct {
	data []int
}

func (h *Heap) Push(value int) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

func (h *Heap) heapifyUp(index int) {
	for h.data[index] > h.data[parent(index)] {
		h.data[index], h.data[parent(index)] = h.data[parent(index)], h.data[index]
		index = parent(index)
	}
}

func (h *Heap) Top() (int, error) {
	if h.Empty() {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.data[0], nil
}

func (h *Heap) Pop() (int, error) {
	if h.Empty() {
		return 0, fmt.Errorf("heap is empty")
	}
	value := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return value, nil
}

func (h *Heap) heapifyDown(index int) {
	for {
		left := leftChild(index)
		right := rightChild(index)
		largest := index
		if left < len(h.data) && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < len(h.data) && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) Size() int {
	return len(h.data)
}

func (h *Heap) Empty() bool {
	return len(h.data) == 0
}
