package structures

import "errors"

type Stack struct {
	items []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(elem interface{}) {
	s.items = append(s.items, elem)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	val := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return val, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	return s.items[len(s.items) - 1], nil
}
