package ch3

import "github.com/dishbreak/go-cracking-coding/lib/stack"

type SortStack struct {
	*stack.IntStack
	tmp *stack.IntStack
}

func (s *SortStack) Push(i int) {
	if s.Empty() {
		s.IntStack.Push(i)
		return
	}

	for !s.Empty() && s.Peek() > i {
		s.tmp.Push(s.Pop())
	}
	s.IntStack.Push(i)
	for !s.tmp.Empty() {
		s.IntStack.Push(s.tmp.Pop())
	}
}

func NewSortStack() *SortStack {
	return &SortStack{
		IntStack: &stack.IntStack{},
		tmp:      &stack.IntStack{},
	}
}
