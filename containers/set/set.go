package set

type Set struct {
	data map[interface{}]bool
}

type SetIterator struct {
	set   *Set
	index int
}

func (si *SetIterator) Next() {
	si.index++
}

func (si *SetIterator) HasNext() bool {
	return si.index < len(si.set.data)
}

func (s *Set) Begin() *SetIterator {
	return &SetIterator{
		set:   s,
		index: 0,
	}
}

func (s *Set) End() *SetIterator {
	return &SetIterator{
		set:   s,
		index: len(s.data),
	}
}

func New() *Set {
	return &Set{
		data: make(map[interface{}]bool),
	}
}

func Insert(s *Set, value interface{}) {
	s.data[value] = true
}

func Contains(s *Set, value interface{}) bool {
	return s.data[value]
}

func Erase(s *Set, value interface{}) {
	delete(s.data, value)
}

func Size(s *Set) int {
	return len(s.data)
}

func (s *Set) Empty() bool {
	return len(s.data) == 0
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Insert(value interface{}) {
	s.data[value] = true
}

func (s *Set) Erase(value interface{}) {
	delete(s.data, value)
}

func (s *Set) Clear() {
	s.data = make(map[interface{}]bool)
}
