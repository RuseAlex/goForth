package forth

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) String() string {
	return fmt.Sprintf("%d", s.items)
}

// Checks if the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Add a new item to the stack
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Return the length of the stack
func (s *Stack) Len() int {
	return len(s.items)
}

// Removes the top element in the stack and returns its valuie
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("can't pop; stack empty")
	}
	popItem := s.items[s.Len()-1]
	s.items = s.items[:(s.Len() - 1)]
	return popItem, nil
}

func (s *Stack) Dup() error {
	dupItem := s.items[s.Len()-1]
	s.items = append(s.items, dupItem)
	return nil
}

func (s *Stack) Drop() error {
	if s.Len() == 0 {
		return nil
	}
	s.items = s.items[:(s.Len() - 1)]
	return nil
}

func (s *Stack) Swap() error {
	if s.Len() < 2 {
		return nil
	}
	firstNum := s.items[s.Len()-1]
	secondNum := s.items[s.Len()-2]
	s.items[s.Len()-1] = secondNum
	s.items[s.Len()-2] = firstNum
	return nil
}

func (s *Stack) Over() error {
	if s.Len() < 2 {
		return nil
	}
	secondNum := s.items[s.Len()-2]
	s.items = append(s.items, secondNum)
	return nil
}

func (s *Stack) Rot() error {
	if s.Len() < 3 {
		return nil
	}
	firstNum := s.items[s.Len()-1]
	secondNum := s.items[s.Len()-2]
	thirdNum := s.items[s.Len()-3]
	s.items[s.Len()-1] = thirdNum
	s.items[s.Len()-2] = firstNum
	s.items[s.Len()-3] = secondNum
	return nil
}
