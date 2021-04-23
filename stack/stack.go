package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	list []interface{}
}

func NewStack() *Stack {
	l := make([]interface{}, 0, 100)
	return &Stack{list: l}
}

func (s *Stack) Len() int {
	return len(s.list)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Len() == 0 {
		return nil, errors.New("stack is empty")
	}
	length := s.Len()
	v := s.list[length-1]
	s.list = s.list[:length-1]
	return v, nil
}

func (s *Stack) Push(item interface{}) error {
	s.list = append(s.list, item)
	return nil
}
func (s *Stack) PrintAll() {
	for l := len(s.list) - 1; l >= 0; l-- {
		fmt.Println("this item is :", s.list[l])
	}
}
