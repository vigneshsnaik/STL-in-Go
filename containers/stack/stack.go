package stack

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) Clear() {
	s.data = make([]interface{}, 0)
}
