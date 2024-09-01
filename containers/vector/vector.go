package vector

import (
	"fmt"
)

type Vector struct {
	data []interface{}
}
type VectorIterator struct {
	vector *Vector
	index  int
}

func (vi *VectorIterator) Next() {
	vi.index++
}

func (vi *VectorIterator) HasNext() bool {
	return vi.index < len(vi.vector.data)
}

func (v *Vector) Begin() *VectorIterator {
	return &VectorIterator{
		vector: v,
		index:  0,
	}
}

func (v *Vector) End() *VectorIterator {
	return &VectorIterator{
		vector: v,
		index:  len(v.data),
	}
}

func (vi *VectorIterator) Value() (interface{}, error) {
	if vi.index < 0 || vi.index >= len(vi.vector.data) {
		return nil, fmt.Errorf("index out of range")
	}
	return vi.vector.data[vi.index], nil
}

func New() *Vector {
	return &Vector{
		data: make([]interface{}, 0),
	}
}
func NewVectorWithSize(size int, value int) *Vector {
	return &Vector{
		data: make([]interface{}, size),
	}
}
func NewVectorWithValues(values ...interface{}) *Vector {
	vec := &Vector{
		data: make([]interface{}, len(values)),
	}
	copy(vec.data, values)
	return vec
}

func (v *Vector) PushBack(value interface{}) {
	v.data = append(v.data, value)
}

func (v *Vector) At(index int) (interface{}, error) {
	if index < 0 || index >= len(v.data) {
		return nil, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

func (v *Vector) PopBack() (interface{}, error) {
	if len(v.data) == 0 {
		return nil, fmt.Errorf("vector is empty")
	}
	value := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return value, nil
}

func (v *Vector) Size() int {
	return len(v.data)
}

func (v *Vector) Empty() bool {
	return len(v.data) == 0
}

func (v *Vector) Reverse() {
	for i, j := 0, len(v.data)-1; i < j; i, j = i+1, j-1 {
		v.data[i], v.data[j] = v.data[j], v.data[i]
	}
}

func (v *Vector) Swap(i, j int) error {
	if i < 0 || i >= len(v.data) || j < 0 || j >= len(v.data) {
		return fmt.Errorf("index out of range")
	}
	v.data[i], v.data[j] = v.data[j], v.data[i]
	return nil
}

func (v *Vector) Back() (interface{}, error) {
	if len(v.data) == 0 {
		return nil, fmt.Errorf("vector is empty")
	}
	return v.data[len(v.data)-1], nil
}

func (v *Vector) Front() (interface{}, error) {
	if len(v.data) == 0 {
		return nil, fmt.Errorf("vector is empty")
	}
	return v.data[0], nil
}

func (v *Vector) Clear() {
	v.data = make([]interface{}, 0)
}

func (v *Vector) Erase(index int) error {
	if index < 0 || index >= len(v.data) {
		return fmt.Errorf("index out of range")
	}
	v.data = append(v.data[:index], v.data[index+1:]...)
	return nil
}

func (v *Vector) Insert(index int, value interface{}) error {
	if index < 0 || index > len(v.data) {
		return fmt.Errorf("index out of range")
	}
	v.data = append(v.data[:index], append([]interface{}{value}, v.data[index:]...)...)
	return nil
}

func (v *Vector) Assign(values ...interface{}) {
	v.data = make([]interface{}, len(values))
	copy(v.data, values)
}

func (v *Vector) AssignWithSize(size int, value interface{}) {
	v.data = make([]interface{}, size)
	for i := 0; i < size; i++ {
		v.data[i] = value
	}
}

func (v *Vector) Capacity() int {
	return cap(v.data)
}
